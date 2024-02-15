package pages

import (
	// "bytes"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type spaRequestBody struct {
	TemplateFile string `json:"TemplateFile"`
	Location     string `json:"Location"`
}

func SpaHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Just started spa handler")

	body := r.Body
	defer body.Close()

	bodyRead, err := io.ReadAll(body)
	if err != nil {
		log.Printf("Error reading request body: %s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	var requestBody spaRequestBody
	err = json.Unmarshal(bodyRead, &requestBody)
	if err != nil {
		log.Printf("Error decoding JSON: %s", err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Printf("TemplateFile: %s, Location: %s\n", requestBody.TemplateFile, requestBody.Location)
	buf := &bytes.Buffer{}


	// fmt.Fprintf(w, "Failed to load template content properly.")
	err = GetTemplates().ExecuteTemplate(buf, fmt.Sprintf("%s.html", requestBody.TemplateFile), nil)
	log.Printf("The return is: %s", buf.String())
	// err = GetTemplates().ExecuteTemplate(w, fmt.Sprintf("%s.html", requestBody.TemplateFile), nil)
	buf.WriteTo(w)

	if err != nil {
		log.Printf("Error when loading template from SPA:\n %s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
