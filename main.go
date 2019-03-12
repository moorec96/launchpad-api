package main

import (
	"launchpad-api/api"
	"launchpad-api/services"
	"net/http"
)

func main() {

	api.HandleRoutes()
	services.InitiateDatabase()
	http.ListenAndServe(":8080", nil)
}
