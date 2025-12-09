package engine

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

type GoTemplateEngine struct{}

func NewGoTemplateEngine() *GoTemplateEngine {
	return &GoTemplateEngine{}
}

// Render reads files from the package directory, and returns a map of filename -> rendered content
// For MVP, if we return specific JSONs, the interface might need adjustment, but for now let's return a map or a joined byte slice.
// The interface said: Render(packagePath string, values map[string]interface{}) ([]byte, error)
// Let's assume it returns a JSON list or a multi-document stream?
// Or maybe it renders to a specific output structure.
// Let's implement it to find all .json files in templates/ and render them.

func (e *GoTemplateEngine) Render(packagePath string, values map[string]interface{}) ([]byte, error) {
	templatesDir := filepath.Join(packagePath, "templates")

	// Check if directory exists
	if _, err := os.Stat(templatesDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("templates directory not found in %s", packagePath)
	}

	var parsedTemplates []string

	// Walk through templates directory
	err := filepath.Walk(templatesDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		// Only process .json or .yaml files?
		if !strings.HasSuffix(path, ".json") && !strings.HasSuffix(path, ".yaml") {
			return nil
		}

		// Parse the file content as a template
		tmplName := filepath.Base(path)
		tmpl, err := template.New(tmplName).Option("missingkey=error").ParseFiles(path)
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", path, err)
		}

		// Execute the template with values
		var buf bytes.Buffer
		data := map[string]interface{}{
			"Values": values,
		}

		if err := tmpl.Execute(&buf, data); err != nil {
			return fmt.Errorf("failed to render template %s: %w", path, err)
		}

		parsedTemplates = append(parsedTemplates, buf.String())
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Join all rendered templates.
	// If they are JSONs, we might want to return them as a JSON array or just concatenated?
	// ColonyOS expects JSON. If we have multiple JSONs (workflow + executors),
	// typically we might want to return a list of objects or handle them separately.
	// For MVP, let's return them as a JSON array wrapped in string.
	// Actually, usually a package manager submits one main thing or a set of things.
	// Let's join them with newlines for now, assuming the submitter will parse them or they are separate.
	// Wait, if we concat JSONs like `{...}{...}` it's valid JSON stream but not valid JSON file.
	// Let's make it a JSON Array `[{...}, {...}]` if multiple.

	if len(parsedTemplates) == 0 {
		return nil, fmt.Errorf("no templates found")
	}

	result := "[" + strings.Join(parsedTemplates, ",") + "]"
	return []byte(result), nil
}

// Helper to load values.yaml
func LoadValues(packagePath string) (map[string]interface{}, error) {
	valuesPath := filepath.Join(packagePath, "values.yaml")
	file, err := os.Open(valuesPath)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]interface{}), nil
		}
		return nil, err
	}
	defer file.Close()

	var values map[string]interface{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&values); err != nil {
		return nil, err
	}

	return values, nil
}
