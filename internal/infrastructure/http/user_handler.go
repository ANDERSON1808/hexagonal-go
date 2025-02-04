package httphandler

import (
	"ANDERSON1808/hexagonal-go/internal/application"
	"ANDERSON1808/hexagonal-go/internal/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CreateUser @Summary Crear un nuevo usuario
// @Description Crea un usuario y lo almacena en la base de datos
// @Tags Usuarios
// @Accept  json
// @Produce  json
// @Param user body domain.User true "Datos del usuario"
// @Success 201 {object} domain.User
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Failed to create user"
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "❌ Invalid request", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		http.Error(w, "❌ Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUser @Summary Obtener un usuario por ID
// @Description Devuelve los detalles de un usuario según su ID
// @Tags Usuarios
// @Produce json
// @Param id path int true "ID del usuario"
// @Success 200 {object} domain.User
// @Failure 400 {object} string "Invalid ID"
// @Failure 404 {object} string "User not found"
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "❌ Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(uint(id))
	if err != nil {
		http.Error(w, "❌ User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// GetAllUsers @Summary Obtener todos los usuarios
// @Description Devuelve una lista de todos los usuarios registrados en la base de datos
// @Tags Usuarios
// @Produce json
// @Success 200 {array} domain.User
// @Failure 500 {object} string "Failed to fetch users"
// @Router /users/all [get]
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		http.Error(w, "❌ Failed to fetch users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// DeleteUser @Summary Eliminar un usuario por ID
// @Description Elimina un usuario de la base de datos según su ID
// @Tags Usuarios
// @Param id path int true "ID del usuario a eliminar"
// @Success 204 {string} string "User deleted successfully"
// @Failure 400 {object} string "Invalid ID"
// @Failure 404 {object} string "User not found"
// @Failure 500 {object} string "Failed to delete user"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "❌ Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteUser(uint(id))
	if err != nil {
		http.Error(w, "❌ Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// CreateUsersConcurrently @Summary Crear múltiples usuarios concurrentemente
// @Description Recibe una lista de usuarios y los crea de manera concurrente para mejorar el rendimiento
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param users body []domain.User true "Lista de usuarios a crear"
// @Success 202 {string} string "Processing users creation concurrently..."
// @Failure 400 {object} string "Invalid request payload"
// @Router /users/concurrent [post]
func (h *UserHandler) CreateUsersConcurrently(w http.ResponseWriter, r *http.Request) {
	var users []domain.User
	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	go h.service.CreateUsersConcurrently(users)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Processing users creation concurrently..."))
}
