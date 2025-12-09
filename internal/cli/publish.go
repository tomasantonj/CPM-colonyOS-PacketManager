package cli

import (
	"fmt"

	"github.com/colonyos/cpm/internal/infra/registry"
	"github.com/colonyos/cpm/internal/infra/storage"
	"github.com/colonyos/cpm/internal/usecase"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(publishCmd)
}

var publishCmd = &cobra.Command{
	Use:   "publish [path]",
	Short: "Publish a package to the registry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		cpmHome, err := GetCPMHome()
		if err != nil {
			fmt.Printf("Error getting CPM home: %v\n", err)
			return
		}

		// Services
		pkgService := storage.NewFsPackageService()

		// Use MockRegistry
		regService, err := registry.NewMockRegistryService(cpmHome)
		if err != nil {
			fmt.Printf("Error initializing registry: %v\n", err)
			return
		}

		uc := usecase.NewPublishPackageUseCase(pkgService, regService)
		err = uc.Execute(path)
		if err != nil {
			fmt.Printf("Error publishing package: %v\n", err)
			return
		}

		// Success message handled by usecase output usually or here
		// Usecase prints, but let's be sure
	},
}
