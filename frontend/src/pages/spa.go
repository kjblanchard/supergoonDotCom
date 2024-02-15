package pages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	debug bool = false
)

type spaRequestBody struct {
	TemplateFile string `json:"TemplateFile"`
}

func SpaHandler(w http.ResponseWriter, r *http.Request) {
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
	if debug {
		debugRequest(&w, r, &requestBody)
	} else {
		request(&w, r, &requestBody)
	}

}

func debugRequest(w *http.ResponseWriter, r *http.Request, requestBody *spaRequestBody) {
	buf := &bytes.Buffer{}
	err := GetTemplates().ExecuteTemplate(buf, fmt.Sprintf("%s.html", requestBody.TemplateFile), nil)
	if err != nil {
		http.Error(*w, "Internal Server Error", http.StatusInternalServerError)
	}
	log.Printf("The return is: %s", buf.String())
	buf.WriteTo(*w)
}

func request(w *http.ResponseWriter, r *http.Request, requestBody *spaRequestBody) {
	err := GetTemplates().ExecuteTemplate(*w, fmt.Sprintf("%s.html", requestBody.TemplateFile), nil)
	if err != nil {
		http.Error(*w, "Internal Server Error", http.StatusInternalServerError)
	}
}
