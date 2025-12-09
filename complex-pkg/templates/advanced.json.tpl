{
    "name": "{{ required "name is required" .Values.name }}",
    "env": "{{ .Values.environment | upper }}",
    "config_dump": {{ .Values.config | toJson }},
    "feature_count": {{ len .Values.config.features }}
}