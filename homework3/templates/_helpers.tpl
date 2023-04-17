{{- define "app.service" -}}
{{- printf "%s-service" .Values.app.name }}
{{- end }}