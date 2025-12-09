package storage

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"path/filepath"

	"github.com/colonyos/cpm/pkg/domain"
	"gopkg.in/yaml.v3"
)

type FsPackageService struct{}

func NewFsPackageService() *FsPackageService {
	return &FsPackageService{}
}

func (s *FsPackageService) Initialize(name string) error {
	// 1. Create directory
	if _, err := os.Stat(name); !os.IsNotExist(err) {
		return fmt.Errorf("directory %s already exists", name)
	}

	if err := os.MkdirAll(filepath.Join(name, "templates"), 0755); err != nil {
		return err
	}

	// 2. Create colony.yaml
	manifest := domain.ColonyManifest{
		APIVersion:  "v1",
		Name:        name,
		Version:     "0.1.0",
		Description: "A ColonyOS package",
		Maintainers: []domain.Maintainer{{Name: "Your Name"}},
	}

	if err := s.writeYAML(filepath.Join(name, "colony.yaml"), manifest); err != nil {
		return err
	}

	// 3. Create values.yaml
	defaultValues := map[string]interface{}{
		"replicas": 1,
		"resources": map[string]interface{}{
			"cpu": "1000m",
			"mem": "512Mi",
		},
	}
	if err := s.writeYAML(filepath.Join(name, "values.yaml"), defaultValues); err != nil {
		return err
	}

	// 4. Create README.md
	readmeContent := fmt.Sprintf("# %s\n\nDescription: %s\n", name, manifest.Description)
	if err := os.WriteFile(filepath.Join(name, "README.md"), []byte(readmeContent), 0644); err != nil {
		return err
	}

	return nil
}

func (s *FsPackageService) LoadManifest(path string) (*domain.ColonyManifest, error) {
	manifestPath := filepath.Join(path, "colony.yaml")
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest at %s: %w", manifestPath, err)
	}

	var manifest domain.ColonyManifest
	if err := yaml.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}

	return &manifest, nil
}

func (s *FsPackageService) Pack(path string, name string, version string) (string, error) {
	// Verify path exists
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		return "", fmt.Errorf("%s is not a directory", path)
	}

	// Define artifact name: {name}-{version}.cpm
	artifactName := fmt.Sprintf("%s-%s.cpm", name, version)

	// Create output file
	outFile, err := os.Create(artifactName)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	// Create gzip writer
	gw := gzip.NewWriter(outFile)
	defer gw.Close()

	// Create tar writer
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// Walk through the directory and add files to tar
	err = filepath.Walk(path, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Configure header
		header, err := tar.FileInfoHeader(fi, fi.Name())
		if err != nil {
			return err
		}

		// Update name to be relative to the package root
		relPath, err := filepath.Rel(path, file)
		if err != nil {
			return err
		}
		header.Name = relPath

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			defer data.Close()
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	return artifactName, nil
}

func (s *FsPackageService) Unpack(artifactPath string, destPath string) error {
	file, err := os.Open(artifactPath)
	if err != nil {
		return err
	}
	defer file.Close()

	gr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gr.Close()

	tr := tar.NewReader(gr)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(destPath, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return err
			}
			f, err := os.Create(target)
			if err != nil {
				return err
			}
			if _, err := io.Copy(f, tr); err != nil {
				f.Close()
				return err
			}
			f.Close()
		}
	}
	return nil
}

// Helper to write YAML files
func (s *FsPackageService) writeYAML(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)
	return encoder.Encode(data)
}
