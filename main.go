package main

import (
	"log"
)

func main() {
	database, err := PostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Init(); err != nil {
		log.Fatal(err)
	}
	
	// colocar a porta de serviço no env
	server := NewAPIServer(":8000", database)
	server.Run()
}