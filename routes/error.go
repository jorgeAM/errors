package routes

import (
	"github.com/gorilla/mux"
	"github.com/jorgeAM/error/controllers"
)

// SetUpErrorRouter -> set up error's routes
func SetUpErrorRouter(router *mux.Router) {
	r := router.PathPrefix("/api/errors").Subrouter()
	r.HandleFunc("/{code}", controllers.GetError).Methods("GET")
	r.HandleFunc("/", controllers.NewError).Methods("POST")
	r.HandleFunc("/save-errors", controllers.SaveErrorsFromFile).Methods("POST")
}
