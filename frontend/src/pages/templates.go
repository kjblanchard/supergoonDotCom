package pages

import (
	"text/template"
)

var templates *template.Template

func LoadTemplates() {
	// templates = template.Must(template.ParseGlob("templates/*.html"))
	templates = template.Must(template.ParseGlob("templates/**/*.html"))
}

func GetTemplates() *template.Template {
	return templates
}
