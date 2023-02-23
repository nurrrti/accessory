package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router { // Initialize a new httprouter router instance.
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	// Register the relevant methods, URL patterns and handler functions for our // endpoints using the HandlerFunc() method. Note that http.MethodGet and
	// http.MethodPost are constants which equate to the strings "GET" and "POST" // respectively.
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/accessory", app.createAccessoryHandler)
	router.HandlerFunc(http.MethodGet, "/v1/accessory/:id", app.showAccessoryHandler)
	router.HandlerFunc(http.MethodPost, "/v1/user", app.createUserHandler)
	router.HandlerFunc(http.MethodGet, "/v1/user/:id", app.showUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/wish", app.createWishHandler)
	router.HandlerFunc(http.MethodGet, "/v1/wish/:id", app.showWishHandler)
	// Return the httprouter instance.
	return router
}
