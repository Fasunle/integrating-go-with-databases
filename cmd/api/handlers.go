package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

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
			"user":   u,
		},
	}

	app.WriteJSON(w, http.StatusAccepted, payload)
}
func (app *Config) UserLogout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the user logout endpoint")
}

func (app *Config) UserDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the user delete user account endpoint")
}

func (app *Config) UserSignup(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
	}

	err := app.ReadJSON(w, r, &requestPayload)

	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	user := data.User{
		Email:     requestPayload.Email,
		FirstName: requestPayload.FirstName,
		LastName:  requestPayload.LastName,
		Password:  requestPayload.Password,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	_, err = user.Insert(user)
	if err != nil {
		log.Println("Error creating user", err)
		app.ErrorJSON(w, errors.New("could not create a new user"), http.StatusBadRequest)
		return
	}

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Signed up user %s", requestPayload.Email),
		Data:    user,
	}

	app.WriteJSON(w, http.StatusAccepted, payload)

}

func (app *Config) UserReset(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email string `json:"email"`
	}

	app.ReadJSON(w, r, &requestPayload)

	passwords := data.Password{}

	err := passwords.Insert(requestPayload.Email)

	if err != nil {
		log.Println("Error occurred while resetting password", err)
		app.ErrorJSON(w, errors.New("could not reset the password"), http.StatusBadRequest)
		return
	}

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Reset password for user %s", requestPayload.Email),
		Data:    "visit the link sent to your email to reset your password",
	}

	// TODO: send email to user with the link to reset password

	app.WriteJSON(w, http.StatusAccepted, payload)
}

// change the user password given that the code provided to the user's email is valid
func (app *Config) UserConfirmPassword(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Code     string `json:"code"`
		Password string `json:"password"`
	}

	app.ReadJSON(w, r, &requestPayload)

	passwords := data.Password{}
	valid, err := passwords.ValidateCode(requestPayload.Email, requestPayload.Code)

	if err != nil {
		log.Println("Error occurred while validating password", err)
		app.ErrorJSON(w, errors.New("could not validate the password"), http.StatusBadRequest)
		return
	}

	if !valid {
		log.Println("Invalid code", err)
		app.ErrorJSON(w, errors.New("invalid code"), http.StatusBadRequest)
		return
	}

	user := data.User{}
	u, _ := user.GetByEmail(requestPayload.Email)
	u.ResetPassword(requestPayload.Password)

	u, _ = user.GetByEmail(requestPayload.Email)
	err = passwords.Update(requestPayload.Code, u.Password)

	if err != nil {
		log.Println("Error occurred while updating password", err)
		app.ErrorJSON(w, errors.New("could not update the password"), http.StatusBadRequest)
		return
	}

	// TODO: send email to user that the password has been changed

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Password changed for user %s", requestPayload.Email),
		Data:    "password changed successfully",
	}

	app.WriteJSON(w, http.StatusAccepted, payload)
}
