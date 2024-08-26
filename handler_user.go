package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thomas21-jquerry/rssagg/internal/database"
)

func (apiCfg *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string `json: "name"`
	}
	decoder :=json.NewDecoder(r.Body);
	params := parameters{};
	err := decoder.Decode(&params);
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("error passing json %v", err));
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	});
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("error adding new user %v", err));
		return
	}
	responseWithJSON(w, 201, databaseUserToUser(user))
}


func (apiCfg *apiConfig)handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User){
	responseWithJSON(w, 200,databaseUserToUser(user))
}