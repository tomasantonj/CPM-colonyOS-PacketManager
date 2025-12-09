package cli

import (
	"fmt"

	"github.com/colonyos/cpm/internal/infra/colony"
	"github.com/colonyos/cpm/internal/infra/storage"
	"github.com/colonyos/cpm/internal/usecase"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(uninstallCmd)
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall [name]",
	Short: "Uninstall a package",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		cpmHome, err := GetCPMHome()
		if err != nil {
			fmt.Printf("Error getting CPM home directory: %v\n", err)
			return
		}

		stateService, err := storage.NewJSONStateService(cpmHome)
		if err != nil {
			fmt.Printf("Error initializing state service: %v\n", err)
			return
		}

		// Mock client for uninstall notifications
		sdk := colony.NewMockSDK()

		uc := usecase.NewUninstallPackageUseCase(stateService, sdk)
		err = uc.Execute(name)
		if err != nil {
			fmt.Printf("Error uninstalling package: %v\n", err)
			return
		}

		fmt.Printf("Package %s uninstalled.\n", name)
	},
}
