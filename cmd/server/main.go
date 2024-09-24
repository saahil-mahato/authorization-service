package server

import (
	"log"
	"os"

	permify_grpc "github.com/Permify/permify-go/grpc"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/saahil-mahato/authorization-service/internal/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartServer() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found")
	}

	permifyClient, err := permify_grpc.NewClient(
		permify_grpc.Config{
			Endpoint: os.Getenv("PERMIFY_HOST"),
		},
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Error while setting up the permify client. Error %s", err.Error())
	}

	r := gin.Default()
	api.SetupRoutes(r, permifyClient)

	r.Run()
}
