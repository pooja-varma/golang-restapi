// routes/routes.go
package routes

import (
	"net/http"

	"restapi.com/myuser/myapi/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define routes for user-related HTTP requests
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	// Define default route for all other HTTP requests
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to my API!"))
	})

	return router
}
