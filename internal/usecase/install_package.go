package usecase

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/colonyos/cpm/internal/engine"
	"github.com/colonyos/cpm/pkg/domain"
)

type InstallPackageUseCase struct {
	pkgService domain.PackageService
	renderer   domain.TemplateEngine
	submitter  domain.Submitter
}

func NewInstallPackageUseCase(pkgService domain.PackageService, renderer domain.TemplateEngine, submitter domain.Submitter) *InstallPackageUseCase {
	return &InstallPackageUseCase{
		pkgService: pkgService,
		renderer:   renderer,
		submitter:  submitter,
	}
}

func (u *InstallPackageUseCase) Execute(path string, setValues map[string]interface{}) error {
	// 0. Prepare workPath (handle archive vs directory)
	workPath := path
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
	for _, spec := range specs {
		// Detect type (Executor vs Workflow) - logic needed
		// For now, just submit as workflow if it looks like one, or print.

		jsonBytes, _ := json.MarshalIndent(spec, "", "  ")

		// Heuristic to guess type?
		// Real implementation would look for "conditions" key for Executors or "process" for Workflows.

		err := u.submitter.SubmitWorkflow(jsonBytes)
		if err != nil {
			return err
		}
	}

	return nil
}
