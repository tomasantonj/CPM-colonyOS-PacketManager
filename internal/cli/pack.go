package cli

import (
	"fmt"
	"github.com/colonyos/cpm/internal/infra/storage"
	"github.com/colonyos/cpm/internal/usecase"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(packCmd)
}

var packCmd = &cobra.Command{
	Use:   "pack [directory]",
	Short: "Package a directory into a .cpm artifact",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}
		
		pkgService := storage.NewFsPackageService()
		uc := usecase.NewPackPackageUseCase(pkgService)
		
		err := uc.Execute(dir)
		if err != nil {
			fmt.Printf("Error packing package: %v\n", err)
			return
		}
		
		fmt.Println("Success!")
	},
}
