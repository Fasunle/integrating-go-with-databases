package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RealIP, middleware.Logger, middleware.Throttle(5), middleware.Recoverer, middleware.Heartbeat("/"))

	r.Get("/auth/login", app.UserLogin)
	r.Get("/auth/logout", app.UserLogout)
	r.Get("/auth/signup", app.UserSignup)
	r.Get("/auth/reset", app.UserReset)

	return r
}
