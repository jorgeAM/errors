package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/jorgeAM/error/migrations"
	"github.com/jorgeAM/error/routes"
)

func main() {
	var migrate string

	flag.StringVar(&migrate, "migrate", "no", "Make migrations")
	flag.Parse()
	if migrate == "yes" {
		log.Println("comenzó la migración")
		migrations.Migrate()
		log.Println("Finalizó la migración")
	}

	r := routes.InitRoutes()
	s := &http.Server{
		Addr:    ":4000",
		Handler: r,
	}

	log.Println("Servidor corriendo en http://localhost:8000")
	log.Println(s.ListenAndServe())
	log.Println("Finalizo la ejecución del programa")
}
