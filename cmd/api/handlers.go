package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Fasunle/integrating-go-with-databases/auth"
	"github.com/Fasunle/integrating-go-with-databases/data"
)

func (app *Config) UserLogin(w http.ResponseWriter, r *http.Request) {
	// read data from the request object
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.ReadJSON(w, r, &requestPayload)

	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	user := data.User{}
	u, err := user.GetByEmail(requestPayload.Email)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	matched, err := u.PasswordMatches(requestPayload.Password)

	if !matched || err != nil {
		app.ErrorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// create access and refresh tokens
	tokens, err := auth.CreateTokens(u.Email)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", u.Email),
		Data: map[string]any{
			"tokens": tokens,
			// "refresh_token": tokens.RefreshToken,
			"user": u,
		},
	}

	app.WriteJSON(w, http.StatusAccepted, payload)
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
