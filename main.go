package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/thomas21-jquerry/rssagg/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main(){
	godotenv.Load(".env")

	port := os.Getenv("PORT");
	if port == "" {
		log.Fatal("Port not found");
	} 

	dbUrl := os.Getenv("DB_URL");
	if dbUrl == "" {
		log.Fatal("DB Url not found")
	}
	connection, err := sql.Open("postgres",dbUrl);
	if err != nil {
		log.Fatal("cant connect to DB")
	}

	apiCfg := apiConfig{
		DB: database.New(connection),
	}

	router := chi.NewRouter();

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*", },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter();
	v1Router.Get("/err", handleError)
	v1Router.Get("/healthz", handleReadiness)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	router.Mount("/v1", v1Router)


	srv := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}
	fmt.Printf("server starting on port : %v", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}