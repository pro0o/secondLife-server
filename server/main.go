package main

import (
	"log"
	"secondLife/handlers"
	"secondLife/storage"
)

func main() {
	store, err := storage.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database.")
	server := handlers.NewAPIServer(":8081", store)
	server.Run()
}
