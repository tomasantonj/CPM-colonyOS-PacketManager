package cli

import (
	"fmt"
	"strings"

	"github.com/colonyos/cpm/internal/engine"
	"github.com/colonyos/cpm/internal/infra/colony"
	"github.com/colonyos/cpm/internal/infra/storage"
	"github.com/colonyos/cpm/internal/usecase"
	"github.com/spf13/cobra"
)

var setFlags []string

func init() {
	installCmd.Flags().StringArrayVar(&setFlags, "set", []string{}, "Set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	rootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install [path]",
	Short: "Install a CPM package",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		// Wire up dependencies
		pkgService := storage.NewFsPackageService()
		renderer := engine.NewGoTemplateEngine()
		client := colony.NewMockColonyClient()

		uc := usecase.NewInstallPackageUseCase(pkgService, renderer, client)

		overrides := make(map[string]interface{})
		for _, s := range setFlags {
			parts := strings.SplitN(s, "=", 2)
			if len(parts) == 2 {
				overrides[parts[0]] = parts[1]
			}
		}

		err := uc.Execute(path, overrides)
		if err != nil {
			fmt.Printf("Error installing package: %v\n", err)
			return
		}

		fmt.Println("Installation complete.")
	},
}
