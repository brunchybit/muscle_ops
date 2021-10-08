package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func enableHeaders(w *http.ResponseWriter) {
	//(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Accep", "application/json")
}

func EncodeAndWrite(w http.ResponseWriter, r *http.Request, status int, value interface{}) {
	enableHeaders(&w)
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	err := enc.Encode(value)
	if err != nil {
		log.Fatalf("error encoding result: %s", err)
	}
}
