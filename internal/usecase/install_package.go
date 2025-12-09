package usecase

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/colonyos/cpm/internal/engine"
	"github.com/colonyos/cpm/pkg/domain"
)

type InstallPackageUseCase struct {
	pkgService      domain.PackageService
	renderer        domain.TemplateEngine
	submitter       domain.Submitter
	stateService    domain.StateService
	registryService domain.RegistryService
}

func NewInstallPackageUseCase(pkgService domain.PackageService, renderer domain.TemplateEngine, submitter domain.Submitter, stateService domain.StateService, registryService domain.RegistryService) *InstallPackageUseCase {
	return &InstallPackageUseCase{
		pkgService:      pkgService,
		renderer:        renderer,
		submitter:       submitter,
		stateService:    stateService,
		registryService: registryService,
	}
}

func (u *InstallPackageUseCase) Execute(path string, setValues map[string]interface{}, version string) error {
	// 0. Prepare workPath (handle archive vs directory vs registry fetch)
	workPath := path
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		// Not found locally? Try fetching from registry
		// Assumption: path is the package name
		fmt.Printf("Package %s not found locally, attempting fetch from registry...\n", path)
		if version == "" {
			return fmt.Errorf("version is required when installing from registry")
		}

		artifactPath, err := u.registryService.Fetch(path, version)
		if err != nil {
			return fmt.Errorf("failed to fetch from registry: %w", err)
		}
		// artifactPath is a temp file probably, we need to handle it same as local archive
		path = artifactPath
		// Fallthrough to archive handling
	}

	info, err := os.Stat(path)

	if err != nil {
		return fmt.Errorf("failed to access path: %w", err)
	}

	if !info.IsDir() {
		// It's a file, assume it's an archive
		tempDir, err := os.MkdirTemp("", "cpm-install-*")
		if err != nil {
			return fmt.Errorf("failed to create temp dir: %w", err)
		}
		defer os.RemoveAll(tempDir) // Clean up

		if err := u.pkgService.Unpack(path, tempDir); err != nil {
			return fmt.Errorf("failed to unpack archive: %w", err)
		}
		workPath = tempDir
	}

	// 1. Load Defaults
	values, err := engine.LoadValues(workPath)
	if err != nil {
		return fmt.Errorf("failed to load values: %w", err)
	}

	// 2. Override with --set flags (TODO: Merge logic)
	for k, v := range setValues {
		values[k] = v // Simple top-level override for MVP
	}

	// 3. Render Templates
	renderedBytes, err := u.renderer.Render(workPath, values)
	if err != nil {
		return fmt.Errorf("render failed: %w", err)
	}

	// 4. Parse rendered output (Assuming JSON Array of objects)
	// We mocked Render to return "[{...},{...}]"
	var specs []map[string]interface{}
	if err := json.Unmarshal(renderedBytes, &specs); err != nil {
		// If unmarshal fails as array, try single object? Or maybe it's invalid JSON
		return fmt.Errorf("failed to parse rendered templates as JSON: %w\nOutput: %s", err, string(renderedBytes))
	}

	// 5. Submit each spec
	var lastColonyID string
	var lastName string

	for _, spec := range specs {
		// Detect type (Executor vs Workflow) - logic needed
		// For now, just submit as workflow if it looks like one, or print.

		jsonBytes, _ := json.MarshalIndent(spec, "", "  ")

		// Capture basic info for state
		if n, ok := spec["name"].(string); ok {
			lastName = n
		}
		if c, ok := spec["colonyId"].(string); ok {
			lastColonyID = c
		}

		// Heuristic to guess type?
		// Real implementation would look for "conditions" key for Executors or "process" for Workflows.

		err := u.submitter.SubmitWorkflow(jsonBytes)
		if err != nil {
			return err
		}
	}

	// 6. Save State
	// We need to determine the release Name.
	// Priority: --set name > values.yaml name > manifest name (not loaded here efficiently yet) > directory name
	// For now, let's use the 'name' from the last submitted spec or a default.
	releaseName := "unknown"
	if nameOverride, ok := setValues["name"].(string); ok {
		releaseName = nameOverride
	} else if lastName != "" {
		releaseName = lastName
	}

	// Version? We didn't parse manifest here explicitly in step 1-3.
	// Improvement: Load Manifest in Step 1.

	err = u.stateService.Save(&domain.Release{
		Name:        releaseName, // This logic needs to be more robust
		Version:     "0.1.0",     // Placeholder, need manifest load
		ColonyID:    lastColonyID,
		InstallTime: time.Now(),
	})
	if err != nil {
		fmt.Printf("Warning: failed to save state: %v\n", err)
	}

	return nil
}
