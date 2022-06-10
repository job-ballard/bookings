package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// Adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	scrfHandler := nosurf.New(next)

	scrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return scrfHandler
}

// Loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
