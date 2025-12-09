package main

import (
	"os"
)

func main() {
	content := `{
    "name": "{{ required "name is required" .Values.name }}",
    "env": "{{ .Values.environment | upper }}",
    "config_dump": {{ .Values.config | toJson }},
    "feature_count": {{ len .Values.config.features }}
}`
	err := os.WriteFile("complex-pkg/templates/advanced.json", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}
