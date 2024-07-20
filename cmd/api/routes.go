package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// Handler registration to URLs
	router.HandlerFunc(http.MethodGet, "/v1/", app.homeHandler)
	router.HandlerFunc(http.MethodGet, "/v1/library", app.libraryHandler)

	router.HandlerFunc(http.MethodGet, "/v1/books/:id", app.getBookHandler)
	router.HandlerFunc(http.MethodPost, "/v1/books/create", app.createBookHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/books/:id", app.updateBookHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/books/:id", app.deleteBookHandler)

	router.HandlerFunc(http.MethodGet, "/v1/user/signup", app.userSignupForm)
	router.HandlerFunc(http.MethodPost, "/v1/user/signup", app.userSignupPost)
	router.HandlerFunc(http.MethodGet, "/v1/user/login", app.userLoginForm)
	router.HandlerFunc(http.MethodPost, "/v1/user/login", app.userLoginPost)
	router.HandlerFunc(http.MethodPost, "/v1/user/logout", app.userLogoutPost)

	return router
}
