package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	database, err := PostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Init(); err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
    log.Fatal("Error loading .env file")
  }

	port := os.Getenv("PORT")
	
	server := NewAPIServer(port, database)
	server.Run()
}