package cli

import (
	"os"
	"path/filepath"
)

// GetCPMHome returns the path to the CPM state directory.
// It checks CPM_HOME env var first, falling back to ~/.cpm.
func GetCPMHome() (string, error) {
	if env := os.Getenv("CPM_HOME"); env != "" {
		return env, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".cpm"), nil
}
