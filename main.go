package main

import (
	"fmt"
	"go-dessert/config"
	"go-dessert/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.InitDB()
	defer config.CloseDB()

	router := mux.NewRouter()
	routes.SetupAPIRoutes(router)

	port := ":8000"
	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}