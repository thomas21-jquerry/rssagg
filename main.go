package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("hello world");
	godotenv.Load(".env")

	port := os.Getenv("PORT");
	if port == "" {
		log.Fatal("Port not found");
	} else {
		fmt.Println("PORT: ", port);
	}
}