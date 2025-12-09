package usecase

import (
	"fmt"

	"github.com/colonyos/cpm/pkg/domain"
)

type PackPackageUseCase struct {
	pkgService domain.PackageService
}

func NewPackPackageUseCase(pkgService domain.PackageService) *PackPackageUseCase {
	return &PackPackageUseCase{
		pkgService: pkgService,
	}
}

func (u *PackPackageUseCase) Execute(path string) error {
	// 1. Load Manifest to get version and name
	manifest, err := u.pkgService.LoadManifest(path)
	if err != nil {
		return fmt.Errorf("failed to load manifest: %w", err)
	}

	// 2. Validate essential fields (optional but good practice)
	if manifest.Name == "" || manifest.Version == "" {
		return fmt.Errorf("manifest must have name and version")
	}

	// 3. Pack using manifest details
	artifact, err := u.pkgService.Pack(path, manifest.Name, manifest.Version)
	if err != nil {
		return fmt.Errorf("failed to pack package: %w", err)
	}

	fmt.Printf("Package created: %s\n", artifact)
	return nil
}
