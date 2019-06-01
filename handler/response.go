package handler

import (
	"encoding/json"
	"net/http"
)

func sendJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
