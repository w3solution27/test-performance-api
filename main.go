package main

import (
	"log"
	"net/http"
	"project/middleware"
	"project/routes"

	"github.com/gorilla/mux"
)

func main() {
	// roteador
	router := mux.NewRouter()

	// middleware
	router.Use(middleware.Logger)

	// rotas
	routes.RegisterRoutes(router)

	// Iniciando server
	log.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
