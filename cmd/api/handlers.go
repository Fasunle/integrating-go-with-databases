package main

import (
	"fmt"
	"net/http"
)

func (app *Config) UserLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the user login endpoint")
}
func (app *Config) UserLogout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the user logout endpoint")
}
func (app *Config) UserSignup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the user signup endpoint")
}
func (app *Config) UserReset(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the user reset endpoint")
}
