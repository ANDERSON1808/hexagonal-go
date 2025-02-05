package routes_test

import (
	"ANDERSON1808/hexagonal-go/internal/application/usecases"
	"ANDERSON1808/hexagonal-go/internal/application/usecases/mocks"
	"ANDERSON1808/hexagonal-go/internal/domain"
	"ANDERSON1808/hexagonal-go/internal/infrastructure/http"
	"ANDERSON1808/hexagonal-go/internal/infrastructure/routes"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUserRoutes(t *testing.T) {
	router := mux.NewRouter()

	mockRepo := new(mocks.UserRepository)
	mockRepo.On("Save", mock.AnythingOfType("*domain.User")).Return(nil)
	mockRepo.On("FindByID", mock.AnythingOfType("uint")).Return(&domain.User{ID: 1, Name: "Test User"}, nil)
	mockRepo.On("FindAll").Return([]domain.User{{ID: 1, Name: "Test User"}}, nil)
	mockRepo.On("Delete", mock.AnythingOfType("uint")).Return(nil)

	userService := usecases.NewUserService(mockRepo)
	mockHandler := httphandler.NewUserHandler(userService)

	routes.RegisterUserRoutes(router, mockHandler)

	testCases := []struct {
		description  string
		method       string
		path         string
		body         interface{}
		expectedCode int
	}{
		{"Create User", http.MethodPost, "/users", domain.User{Name: "John Doe"}, http.StatusCreated},
		{"Get All Users", http.MethodGet, "/users/all", nil, http.StatusOK},
		{"Get User by ID", http.MethodGet, "/users/1", nil, http.StatusOK},
		{"Create Users Concurrently", http.MethodPost, "/users/concurrent", []domain.User{{Name: "Alice"}, {Name: "Bob"}}, http.StatusAccepted},
		{"Delete User", http.MethodDelete, "/users/1", nil, http.StatusNoContent},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			var req *http.Request
			if tc.body != nil {
				bodyBytes, _ := json.Marshal(tc.body)
				req, _ = http.NewRequest(tc.method, tc.path, bytes.NewBuffer(bodyBytes))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req, _ = http.NewRequest(tc.method, tc.path, nil)
			}

			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code, "❌ Falló en %s", tc.description)
		})
	}
}
