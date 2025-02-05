package main

import (
	_ "ANDERSON1808/hexagonal-go/docs"
	"ANDERSON1808/hexagonal-go/internal/application/usecases"
	"ANDERSON1808/hexagonal-go/internal/infrastructure/db"
	httphandler "ANDERSON1808/hexagonal-go/internal/infrastructure/http"
	"ANDERSON1808/hexagonal-go/internal/infrastructure/routes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Hexagonal Golang API
// @version 1.0
// @description API para manejo de usuarios con arquitectura Hexagonal en Golang.
// @termsOfService http://swagger.io/terms/
// @contact.name Anderson Dev
// @contact.email anderson@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
func main() {
	fmt.Println("🚀 Servidor corriendo en el puerto 8080...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("main failed env", err)
	}

	postgres := db.NewPostgresDB()
	db.RunMigrations(postgres.DB)

	userRepo := db.NewUserRepository(postgres)
	userService := usecases.NewUserService(userRepo)
	userHandler := httphandler.NewUserHandler(userService)

	router := mux.NewRouter()

	routes.RegisterUserRoutes(router, userHandler)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("✅ Servidor iniciado en http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("❌ Error al iniciar el servidor: %v", err)
	}
}
