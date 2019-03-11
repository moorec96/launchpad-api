package main

import (
	"Routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	Routes.HandleRoutes()
	http.ListenAndServe(":8000", router)
}
