package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/laiane.muckler/api-rest-project/app"
)

func main() {
	database, err := app.PostgresConnection()
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

	server := app.NewAPIServer(port, database)
	server.Run()
}
