package pages

import (
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

	// Now, requestBody contains the decoded JSON data
	fmt.Printf("TemplateFile: %s, Location: %s\n", requestBody.TemplateFile, requestBody.Location)

	// Handle the rest of your logic here...

	// Respond with success or appropriate response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request processed successfully"))
}
