package usecase

import (
	"fmt"

	"github.com/colonyos/cpm/pkg/domain"
)

type UninstallPackageUseCase struct {
	stateService domain.StateService
	submitter    domain.Submitter
}

func NewUninstallPackageUseCase(stateService domain.StateService, submitter domain.Submitter) *UninstallPackageUseCase {
	return &UninstallPackageUseCase{
		stateService: stateService,
		submitter:    submitter,
	}
}

func (u *UninstallPackageUseCase) Execute(name string) error {
	// 1. Check if installed
	release, err := u.stateService.Get(name)
	if err != nil {
		return err
	}

	// 2. Submit deletion to ColonyOS (if supported)
	// For MVP, we'll just print or invoke a "delete" workflow if we had one.
	// Since ColonyOS doesn't explicitly have "uninstall" for a random workflow without tracking IDs,
	// we will simulate it or maybe submit a termination signal.
	// Let's print for now via the submitter/client or just log.

	// Real world: We might have saved the process ID in state and now we kill it.
	// We only saved "ColonyID".

	// Let's assume we just remove it from local state effectively "forgetting" it,
	// or maybe we should tell the user "We removed it from CPM, but please check ColonyOS"
	// until we have robust process tracking.

	fmt.Printf("Uninstalling package %s (ColonyID: %s)...\n", release.Name, release.ColonyID)

	// 3. Remove from state
	return u.stateService.Delete(name)
}
