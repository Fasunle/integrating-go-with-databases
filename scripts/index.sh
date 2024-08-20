#!/bin/bash

# Install Go
if ! command -v go &> /dev/null
then
    echo "Go is not installed. Installing..."
    sudo apt-get update
    sudo apt-get install -y golang-go
fi

# create a module
go mod init github.com/Fasunle/integrating-go-with-databases


# Install required Go packages
go get -u github.com/go-chi/chi/v5
go get -u github.com/go-chi/cors
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/joho/godotenv
go get -u github.com/lib/pq
go get -u github.com/golang-jwt/jwt/v5