package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Println("Responding with 5XX error: ", msg);
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	responseWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}){
	dat , err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("failed to marshal JSON response: %v", payload);
		w.WriteHeader(500)
		return 
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code);
	w.Write(dat)
}