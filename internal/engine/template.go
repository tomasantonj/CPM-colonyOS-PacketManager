package engine

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"gopkg.in/yaml.v3"
)

type GoTemplateEngine struct{}

func NewGoTemplateEngine() *GoTemplateEngine {
	return &GoTemplateEngine{}
}

// Render reads files from the package directory, and returns a map of filename -> rendered content
func (e *GoTemplateEngine) Render(packagePath string, values map[string]interface{}) ([]byte, error) {
	templatesDir := filepath.Join(packagePath, "templates")

	// Check if directory exists
	if _, err := os.Stat(templatesDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("templates directory not found in %s", packagePath)
	}

	// Define helper functions
	funcMap := sprig.TxtFuncMap()

	funcMap["toYaml"] = func(v interface{}) (string, error) {
		data, err := yaml.Marshal(v)
		if err != nil {
			return "", err
		}
		return strings.TrimSuffix(string(data), "\n"), nil
	}

	funcMap["required"] = func(warn string, val interface{}) (interface{}, error) {
		if val == nil {
			return nil, fmt.Errorf("%s", warn)
		}
		if s, ok := val.(string); ok && s == "" {
			return nil, fmt.Errorf("%s", warn)
		}
		return val, nil
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

		// Only process .json, .yaml, or .tpl files
		if !strings.HasSuffix(path, ".json") && !strings.HasSuffix(path, ".yaml") && !strings.HasSuffix(path, ".tpl") {
			return nil
		}

		// Parse the file content as a template
		tmplName := filepath.Base(path)

		// Read file content manually to handle BOM
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read template %s: %w", path, err)
		}

		// Strip BOM if present
		const bom = "\xef\xbb\xbf"
		sContent := string(content)
		sContent = strings.TrimPrefix(sContent, bom)

		tmpl, err := template.New(tmplName).Funcs(funcMap).Option("missingkey=error").Parse(sContent)
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
