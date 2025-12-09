package usecase

import "github.com/colonyos/cpm/pkg/domain"

type ListPackagesUseCase struct {
	stateService domain.StateService
}

func NewListPackagesUseCase(stateService domain.StateService) *ListPackagesUseCase {
	return &ListPackagesUseCase{
		stateService: stateService,
	}
}

func (u *ListPackagesUseCase) Execute() ([]*domain.Release, error) {
	return u.stateService.List()
}
