package usecase

import (
	"fmt"
	"github.com/colonyos/cpm/pkg/domain"
)

type InitPackageUseCase struct {
	pkgService domain.PackageService
}

func NewInitPackageUseCase(pkgService domain.PackageService) *InitPackageUseCase {
	return &InitPackageUseCase{
		pkgService: pkgService,
	}
}

func (u *InitPackageUseCase) Execute(name string) error {
	if name == "" {
		return fmt.Errorf("package name cannot be empty")
	}
	
	err := u.pkgService.Initialize(name)
	if err != nil {
		return fmt.Errorf("failed to initialize package: %w", err)
	}
	
	return nil
}
