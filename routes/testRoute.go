package routes

import (
	"database/sql"

	"go-docker/controllers"

	"github.com/gorilla/mux"
)

// import GetUsers from controllers/testController.go

func RegisterRoutes(db *sql.DB, router *mux.Router) {
	router.HandleFunc("/users", controllers.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUser(db)).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser(db)).Methods("DELETE")
}
