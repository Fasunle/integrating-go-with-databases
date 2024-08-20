package auth

import (
	"errors"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type TokenMaker struct{}
type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (app *TokenMaker) CreateToken(email string, duration time.Duration) (string, error) {
	secret := os.Getenv("SECRET_KEY")
	payload := jwt.MapClaims{
		"exp":   time.Now().Add(duration).Unix(),
		"email": email,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(secret))
}

func (app *TokenMaker) VerifyToken(t string) (bool, error) {
	secret := os.Getenv("SECRET_KEY")
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return false, errors.New("error occurred while parsing token")
	}

	switch {
	case token.Valid:
		return true, nil
	case errors.Is(err, jwt.ErrTokenMalformed) || errors.Is(err, jwt.ErrTokenNotValidYet) || errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrSignatureInvalid):
		return false, nil
	default:
		return false, errors.New("couldn't handle this token")
	}

}

func CreateTokens(email string) (Tokens, error) {
	tm := TokenMaker{}
	tokens := Tokens{
		AccessToken:  "",
		RefreshToken: "",
	}
	// create access token with email and 15 minutes expiration
	access_token, err := tm.CreateToken(email, 15*time.Minute)

	if err != nil {
		return tokens, err
	}
	// create refresh token with email and 1 hour expiration
	refresh_token, err := tm.CreateToken(email, 60*time.Minute)

	if err != nil {
		return tokens, err
	}

	tokens.AccessToken = access_token
	tokens.RefreshToken = refresh_token

	return tokens, nil

}
