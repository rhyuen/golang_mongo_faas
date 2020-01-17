package response

import (
	"encoding/json"
	"net/http"
)

func withError(w http.ResponseWriter, code int, message string) {
	withJSON(w, code, map[string]string{"error": message})
}

func withJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
