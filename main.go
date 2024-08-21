package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load(".env")

	port := os.Getenv("PORT");
	if port == "" {
		log.Fatal("Port not found");
	} else {
		fmt.Println("PORT: ", port);
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
	router.Mount("/v1", v1Router)


	srv := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}
	fmt.Printf("server starting on port : %v", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}