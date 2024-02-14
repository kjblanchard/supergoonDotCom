package pages

import (
	"fmt"
	"log"
	"net/http"
)

type homepageTemplateData struct {
	ScriptFiles []string
	CssFiles    []string
}

func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	data := homepageTemplateData{
		ScriptFiles: []string{
			// "login",
			"homepage",
			"nav",
			// "addScore",
		},
		CssFiles: []string{},
	}
	err := GetTemplates().ExecuteTemplate(w, "homepage.html", data)
	if err != nil {
		log.Printf("Error when loading template:\n %s", err.Error())
		fmt.Fprintf(w, "Failed to load template content properly.")
	}
}
