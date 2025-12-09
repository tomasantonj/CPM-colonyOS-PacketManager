package cli

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/colonyos/cpm/internal/infra/registry"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search for packages in the registry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		cpmHome, err := GetCPMHome()
		if err != nil {
			fmt.Printf("Error getting CPM home: %v\n", err)
			return
		}

		regService, err := registry.NewMockRegistryService(cpmHome)
		if err != nil {
			fmt.Printf("Error initializing registry: %v\n", err)
			return
		}

		results, err := regService.Search(query)
		if err != nil {
			fmt.Printf("Error searching packages: %v\n", err)
			return
		}

		if len(results) == 0 {
			fmt.Println("No packages found.")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "NAME\tMATCH")
		for _, name := range results {
			// Extract just the name part from filename if possible?
			// Mock Registry just searches filenames 'pkg-ver.cpm'
			display := strings.TrimSuffix(name, ".cpm")
			fmt.Fprintf(w, "%s\t%s\n", display, name)
		}
		w.Flush()
	},
}
