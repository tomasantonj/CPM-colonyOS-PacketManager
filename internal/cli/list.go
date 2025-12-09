package cli

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/colonyos/cpm/internal/infra/storage"
	"github.com/colonyos/cpm/internal/usecase"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed packages",
	Run: func(cmd *cobra.Command, args []string) {
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

		uc := usecase.NewListPackagesUseCase(stateService)
		releases, err := uc.Execute()
		if err != nil {
			fmt.Printf("Error listing packages: %v\n", err)
			return
		}

		if len(releases) == 0 {
			fmt.Println("No packages installed.")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "NAME\tVERSION\tINSTALLED\tCOLONY_ID")
		for _, r := range releases {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", r.Name, r.Version, r.InstallTime.Format("2006-01-02 15:04:05"), r.ColonyID)
		}
		w.Flush()
	},
}
