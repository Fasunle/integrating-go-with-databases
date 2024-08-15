package main

import (
	"fmt"
	"log"
	"net/http"
)

var addr = ":3000"

type Config struct{}

func main() {
	app := Config{}
	// start the server on a port and connect the router
	server := &http.Server{
		Addr:    addr,
		Handler: app.routes(),
	}

	fmt.Println("Starting listening to server on ", addr)
	err := server.ListenAndServe()

	if err != nil {
		log.Panicln("Error starting server")
	}
}
