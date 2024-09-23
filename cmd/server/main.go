package server

import (
	"github.com/gin-gonic/gin"
	"github.com/saahil-mahato/authorization-service/internal/api"
	"github.com/saahil-mahato/authorization-service/internal/service"
)

func StartServer() {

	authService := service.NewAuthorizationService()

	r := gin.Default()
	api.SetupRoutes(r, authService)

	r.Run()
}
