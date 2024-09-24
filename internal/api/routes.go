package api

import (
	"github.com/gin-gonic/gin"
	"github.com/saahil-mahato/authorization-service/internal/api/handlers"
	"github.com/saahil-mahato/authorization-service/internal/service"
)

// API routes for Schema
const (
	SCHEMA                          = "/schema"
	SCHEMA_TENANT_ID                = "/schema/:id"
	SCHEMA_TENANT_ID_SCHEMA_VERSION = "/schema/:id/:version"
)

func SetupRoutes(r *gin.Engine, schemaService service.SchemaService) {
	schemaHandler := handlers.NewHandler(schemaService)

	r.POST(SCHEMA, schemaHandler.WriteSchema)
	r.GET(SCHEMA_TENANT_ID, schemaHandler.ListSchema)
	r.GET(SCHEMA_TENANT_ID_SCHEMA_VERSION, schemaHandler.ReadSchema)
}
