package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/saahil-mahato/authorization-service/internal/api"
	"github.com/saahil-mahato/authorization-service/internal/config"
	"github.com/saahil-mahato/authorization-service/internal/models"
	"github.com/saahil-mahato/authorization-service/internal/repository"
	"github.com/saahil-mahato/authorization-service/internal/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func StartServer() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.AuthorizationConfig{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	log.Println("Database and table created successfully!")

	authRepo := repository.NewAuthorizatonRepository(db)
	authService := service.NewAuthorizationService(authRepo)

	r := gin.Default()
	api.SetupRoutes(r, authService)

	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
