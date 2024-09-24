package server

import (
	"log"

	permify_grpc "github.com/Permify/permify-go/grpc"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/saahil-mahato/authorization-service/internal/api"
	"github.com/saahil-mahato/authorization-service/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartServer() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found")
	}

	client, err := permify_grpc.NewClient(
		permify_grpc.Config{
			Endpoint: `localhost:3478`,
		},
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Error while setting up the permify client. Error %s", err)
	}

	schemaService := service.NewSchemaService(client)

	r := gin.Default()
	api.SetupRoutes(r, schemaService)

	r.Run()
}
