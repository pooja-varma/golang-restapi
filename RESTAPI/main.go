// main.go
package main

import (
	"log"
	"net/http"

	"restapi.com/myuser/myapi/routes"
)

func main() {
	router := routes.SetupRoutes()

	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
