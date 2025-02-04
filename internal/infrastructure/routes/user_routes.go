package routes

import (
	httphandler "ANDERSON1808/hexagonal-go/internal/infrastructure/http"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router, userHandler *httphandler.UserHandler) {
	userRoutes := router.PathPrefix("/users").Subrouter()

	userRoutes.HandleFunc("", userHandler.CreateUser).Methods(http.MethodPost)
	userRoutes.HandleFunc("/{id}", userHandler.GetUser).Methods(http.MethodGet)
	userRoutes.HandleFunc("/all", userHandler.GetAllUsers).Methods(http.MethodGet)
	userRoutes.HandleFunc("/concurrent", userHandler.CreateUsersConcurrently).Methods(http.MethodPost)
	userRoutes.HandleFunc("/{id}", userHandler.DeleteUser).Methods(http.MethodDelete)
}
