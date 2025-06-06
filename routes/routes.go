// File: routes/routes.go
package routes

import (
	"go-dessert/handlers"
	"go-dessert/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupAPIRoutes(router *mux.Router) {
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.CORSMiddleware)

	api.HandleFunc("/desserts", handlers.GetAllDessertsHandler).Methods("GET")
	api.HandleFunc("/desserts", handlers.CreateDessertHandler).Methods("POST")
	api.HandleFunc("/desserts/{id}", handlers.GetDessertByIDHandler).Methods("GET")
    api.HandleFunc("/desserts/{id}", handlers.UpdateDessertHandler).Methods("PUT")
    api.HandleFunc("/desserts/{id}", handlers.DeleteDessertHandler).Methods("DELETE")

	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API is healthy!"))
	}).Methods("GET")
}