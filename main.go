package main

import (
	//"fmt"
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

	//fmt.Println("%+v\n", database)

	server := NewAPIServer(":8000", database)
	server.Run()
}