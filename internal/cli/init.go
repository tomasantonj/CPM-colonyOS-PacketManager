package cli

import (
	"fmt"
	"github.com/colonyos/cpm/internal/infra/storage"
	"github.com/colonyos/cpm/internal/usecase"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init [name]",
	Short: "Create a new CPM package",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		
		// Wire up dependencies (Manual DI for now)
		pkgService := storage.NewFsPackageService()
		uc := usecase.NewInitPackageUseCase(pkgService)
		
		err := uc.Execute(name)
		if err != nil {
			fmt.Printf("Error initializing package: %v\n", err)
			return
		}
		
		fmt.Printf("Successfully initialized package '%s'\n", name)
	},
}
