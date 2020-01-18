package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func WithError(w http.ResponseWriter, code int, message string) {
	WithJSON(w, code, map[string]string{"error": message})
}

func WithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
