package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondJson( w http.ResponseWriter, status int, data any) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}

}

func RespondError( w http.ResponseWriter, status int, mensaje string) {
	RespondJson(w, status, map[string]string{"error": mensaje}) //es un mapa devolver como un json

}