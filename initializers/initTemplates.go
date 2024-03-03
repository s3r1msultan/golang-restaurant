package initializers

import (
	"html/template"

	log "github.com/sirupsen/logrus"
)

func InitTemplates() *template.Template {
	tmpl := template.New("")

	patterns := []string{
		"./views/blocks/*.html",
		"./views/pages/*.html",
		// "../../views/blocks/*.html",
		// "../../views/pages/*.html",
		//uncoment this to run the test
	}

	for _, pattern := range patterns {
		globedTemplates, err := tmpl.ParseGlob(pattern)
		if err != nil {
			log.Fatalf("Failed to parse glob pattern '%s': %v", pattern, err)
		}
		tmpl = globedTemplates
	}

	return tmpl
}
