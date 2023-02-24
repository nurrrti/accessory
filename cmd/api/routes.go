package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router { // Initialize a new httprouter router instance.
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/accessory", app.createAccessoryHandler)
	router.HandlerFunc(http.MethodGet, "/v1/accessory/:id", app.showAccessoryHandler)
	router.HandlerFunc(http.MethodPut, "/v1/accessory/:id", app.updateAccessoryHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodGet, "/v1/user/:id", app.showUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/wish", app.createWishHandler)
	router.HandlerFunc(http.MethodGet, "/v1/wish/:id", app.showWishHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/accessory/:id", app.deleteAccessoryHandler)

	return nil
}
