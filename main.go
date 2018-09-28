package main

import (
	"github.com/jlee2920/golang-api.git/routes"
	"log"
	"net/http"
)

func main() {
	server := router.Server{}
	server.InitializeRoutes()

	log.Fatal(http.ListenAndServe(":8080", server.Router))
}
