package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Fasunle/integrating-go-with-databases/data"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // this driver must be imported for it to work
)

var addr = ":3000"

type Config struct{}

func init() {
	// Load the environment variables in from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app := Config{}
	// open database connection
	db := data.Open()

	fmt.Println("Connected to database")
	data.New(db)     // create the models and instantiate the database
	defer db.Close() // always close connection when the server is done

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
