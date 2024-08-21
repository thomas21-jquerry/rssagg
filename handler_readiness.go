package main

import (
	"net/http"
)

func handleReadiness(w http.ResponseWriter, r *http.Request){
	responseWithJSON(w, 200, struct{}{})
}