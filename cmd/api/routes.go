package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RealIP, middleware.Logger, middleware.Throttle(5), middleware.Recoverer, middleware.Heartbeat("/"))

	r.Post("/auth/login", app.UserLogin)
	r.Post("/auth/logout", app.UserLogout)
	r.Post("/auth/signup", app.UserSignup)
	r.Post("/auth/reset", app.UserReset)
	r.Post("/auth/confirm-password", app.UserConfirmPassword)

	return r
}
