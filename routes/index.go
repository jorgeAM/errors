package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRoutes -> init routes
func InitRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hi!</h1>")
	})
	SetUpErrorRouter(r)
	return r
}
