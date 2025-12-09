package usecase

import (
	"fmt"
	"os"

	"github.com/colonyos/cpm/pkg/domain"
)

type PublishPackageUseCase struct {
	pkgService      domain.PackageService
	registryService domain.RegistryService
}

func NewPublishPackageUseCase(pkgService domain.PackageService, registryService domain.RegistryService) *PublishPackageUseCase {
	return &PublishPackageUseCase{
		pkgService:      pkgService,
		registryService: registryService,
	}
}

func (u *PublishPackageUseCase) Execute(path string) error {
	// 1. Pack the package first to ensure we have a fresh artifact
	// Need to load manifest to get name/version
	manifest, err := u.pkgService.LoadManifest(path)
	if err != nil {
		return fmt.Errorf("failed to load manifest: %w", err)
	}

	artifactPath, err := u.pkgService.Pack(path, manifest.Name, manifest.Version)
	if err != nil {
		return fmt.Errorf("failed to pack package: %w", err)
	}
	defer os.Remove(artifactPath) // Clean up local artifact after publish? Or keep it?
	// For now, let's keep it or maybe log that we published it.
	// Actually typical behavior is to pack to temp or pack and upload.
	// Detailed implementation: Pack returns path to created file.

	// 2. Publish to registry
	if err := u.registryService.Publish(artifactPath); err != nil {
		return fmt.Errorf("failed to publish package: %w", err)
	}

	fmt.Printf("Package %s version %s published successfully.\n", manifest.Name, manifest.Version)
	return nil
}
