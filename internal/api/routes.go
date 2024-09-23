package api

import (
	"github.com/gin-gonic/gin"
	"github.com/saahil-mahato/authorization-service/internal/api/handlers"
	"github.com/saahil-mahato/authorization-service/internal/service"
)

const (
	AUTH    = "/auth"
	AUTH_ID = "/auth/:id"
	CHECK   = "/check"
)

func SetupRoutes(r *gin.Engine, authService service.AuthorizationService) {
	h := handlers.NewHandler(authService)

	r.POST(AUTH, h.CreateAuth)
	r.GET(AUTH_ID, h.GetAuth)
	r.GET(AUTH, h.ListAuth)
	r.PUT(AUTH_ID, h.UpdateAuth)
	r.DELETE(AUTH_ID, h.DeleteAuth)
	r.POST(CHECK, h.CheckAuthorization)
}
