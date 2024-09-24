package api

import (
	permify_grpc "github.com/Permify/permify-go/grpc"
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

// API routes for permission
const (
	PERMISSION_CHECK = "/permission/check/:id"
)

// API routes for authorization data
const (
	DATA_WRITE = "/data/write/:id"
)

func SetupRoutes(r *gin.Engine, permifyClient *permify_grpc.Client) {
	schemaService := service.NewSchemaService(permifyClient)
	schemaHandler := handlers.NewSchemaHandler(schemaService)

	permissionService := service.NewPermissionService(permifyClient)
	permissionHandler := handlers.NewPermissionHandler(permissionService)

	dataService := service.NewDataService(permifyClient)
	dataHandler := handlers.NewDataHandler(dataService)

	r.POST(SCHEMA, schemaHandler.WriteSchema)
	r.GET(SCHEMA_TENANT_ID, schemaHandler.ListSchema)
	r.GET(SCHEMA_TENANT_ID_SCHEMA_VERSION, schemaHandler.ReadSchema)

	r.POST(PERMISSION_CHECK, permissionHandler.CheckAccessControl)

	r.POST(DATA_WRITE, dataHandler.WriteAuthorizationData)
}
