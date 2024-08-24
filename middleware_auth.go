package main

import (
	"fmt"
	"net/http"

	"github.com/thomas21-jquerry/rssagg/internal/auth"
	"github.com/thomas21-jquerry/rssagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig)middlewareAuth(handler authHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		apiKey, err := auth.GetAPIKey(r.Header)
		if err!=nil {
			responseWithError(w, 403, fmt.Sprintf("Auth error %v", err));
			return
		}
		user, err := cfg.DB.GetUser(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, 404, fmt.Sprintf("error no user %v", err));
			return
		}
		handler(w,r,user)
	}
}