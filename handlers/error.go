package handlers

import (
	"net/http"
	"encoding/json"
	"log"
)

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		response := map[string]interface{}{
			"errors": map[string]interface{}{
				"detail": "Bad Request",
			},
		}
	json.NewEncoder(w).Encode(response)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		response := map[string]interface{}{
			"errors": map[string]interface{}{
				"detail": "Not Found",
			},
		}
	json.NewEncoder(w).Encode(response)
}

func customErrorString(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	response := message
	json.NewEncoder(w).Encode(response)
}

func okResponse(w http.ResponseWriter) {
	jsonResp, err := json.Marshal("OK")
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}