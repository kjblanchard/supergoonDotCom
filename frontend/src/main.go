package main

import (
	"log"
	"net/http"

	pages "github.com/kjblanchard/supergoonDotCom/pages"
)



func main() {
	pages.LoadTemplates()
	http.HandleFunc("/", pages.HomepageHandler)
	http.HandleFunc("/spa", pages.SpaHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
