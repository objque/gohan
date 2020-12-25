package httputils

import (
	"encoding/json"
	"net/http"
)

func WriteHeader(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	return enc.Encode(v)
}
