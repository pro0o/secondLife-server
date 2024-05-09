package main

import (
	"log"
	"secondLife/handlers"
)

func main() {
	if err := handlers.LoadModel(); err != nil {
		log.Fatal(err)
		return
	}
	server := handlers.NewAPIServer()
	server.Run()
}
