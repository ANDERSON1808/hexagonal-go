package routes

import (
	httphandler "ANDERSON1808/hexagonal-go/internal/infrastructure/http"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterUserRoutes(router *mux.Router, userHandler *httphandler.UserHandler) {
	userRoutes := router.PathPrefix("/users").Subrouter()

	userRoutes.HandleFunc("", userHandler.CreateUser).Methods(http.MethodPost)
	userRoutes.HandleFunc("/all", userHandler.GetAllUsers).Methods(http.MethodGet)
	userRoutes.HandleFunc("/{id}", userHandler.GetUser).Methods(http.MethodGet)
	userRoutes.HandleFunc("/concurrent", userHandler.CreateUsersConcurrently).Methods(http.MethodPost)
	userRoutes.HandleFunc("/{id}", userHandler.DeleteUser).Methods(http.MethodDelete)

	// Ruta para Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
