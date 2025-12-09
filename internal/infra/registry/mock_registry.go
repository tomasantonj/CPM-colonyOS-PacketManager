package registry

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type MockRegistryService struct {
	basePath string
}

func NewMockRegistryService(homeDir string) (*MockRegistryService, error) {
	// Use a subdirectory in CPM home for the "remote" registry
	path := filepath.Join(homeDir, "registry_mock")
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, fmt.Errorf("failed to create mock registry dir: %w", err)
	}
	return &MockRegistryService{basePath: path}, nil
}

func (r *MockRegistryService) Publish(artifactPath string) error {
	fileName := filepath.Base(artifactPath)
	destPath := filepath.Join(r.basePath, fileName)

	// Copy file to "remote"
	srcFile, err := os.Open(artifactPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	fmt.Printf("[MockRegistry] Published %s to %s\n", fileName, r.basePath)
	return nil
}

func (r *MockRegistryService) Fetch(packageName, version string) (string, error) {
	// Guess filename
	fileName := fmt.Sprintf("%s-%s.cpm", packageName, version)
	remotePath := filepath.Join(r.basePath, fileName)

	// Check if exists
	if _, err := os.Stat(remotePath); os.IsNotExist(err) {
		return "", fmt.Errorf("package %s version %s not found in registry", packageName, version)
	}

	// For fetch, we return the path to the file.
	// In a real scenario, we would download it to a temp dir.
	// Here we simulate download by copying to a temp location?
	// Or just returning the path since it's on disk?
	// Let's copy to allow simulation of "download".

	tempDir, err := os.MkdirTemp("", "cpm-fetch-*")
	if err != nil {
		return "", err
	}

	destPath := filepath.Join(tempDir, fileName)

	srcFile, err := os.Open(remotePath)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		return "", err
	}

	fmt.Printf("[MockRegistry] Fetched %s from %s\n", fileName, r.basePath)
	return destPath, nil
}

func (r *MockRegistryService) Search(query string) ([]string, error) {
	files, err := os.ReadDir(r.basePath)
	if err != nil {
		return nil, err
	}

	var results []string
	for _, f := range files {
		if !f.IsDir() && strings.Contains(f.Name(), ".cpm") {
			if strings.Contains(f.Name(), query) {
				results = append(results, f.Name())
			}
		}
	}
	return results, nil
}
