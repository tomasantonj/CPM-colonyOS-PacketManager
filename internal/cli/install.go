package cli

import (
	"fmt"
	"strings"

	"github.com/colonyos/cpm/internal/engine"
	"github.com/colonyos/cpm/internal/infra/colony"
	"github.com/colonyos/cpm/internal/infra/registry"
	"github.com/colonyos/cpm/internal/infra/storage"
	"github.com/colonyos/cpm/internal/usecase"
	"github.com/spf13/cobra"
)

var (
	setFlags     []string
	colonyHost   string
	colonyPort   int
	colonyID     string
	colonyPrvKey string
	cpmVersion   string
)

func init() {
	installCmd.Flags().StringArrayVar(&setFlags, "set", []string{}, "Set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	installCmd.Flags().StringVar(&colonyHost, "host", "localhost", "ColonyOS server host")
	installCmd.Flags().IntVar(&colonyPort, "port", 50080, "ColonyOS server port")
	installCmd.Flags().StringVar(&colonyID, "colonyid", "", "Colony ID (required)")
	installCmd.Flags().StringVar(&colonyPrvKey, "prvkey", "", "Private Key (required)")
	installCmd.Flags().StringVar(&cpmVersion, "version", "", "Package version (required if installing from registry)")

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

		regService, err := registry.NewMockRegistryService(cpmHome)
		if err != nil {
			fmt.Printf("Error initializing registry: %v\n", err)
			return
		}

		// Mock client for uninstall notifications
		sdk := colony.NewMockSDK()
		client := colony.NewColonyClient(sdk)

		uc := usecase.NewInstallPackageUseCase(pkgService, renderer, client, stateService, regService)

		overrides := make(map[string]interface{})
		for _, s := range setFlags {
			parts := strings.SplitN(s, "=", 2)
			if len(parts) == 2 {
				overrides[parts[0]] = parts[1]
			}
		}

		// Inject CLI flags into overrides if appropriate
		if colonyID != "" {
			overrides["colonyId"] = colonyID
		}

		err = uc.Execute(path, overrides, cpmVersion)
		if err != nil {
			fmt.Printf("Error installing package: %v\n", err)
			return
		}

		fmt.Println("Installation complete.")
	},
}
