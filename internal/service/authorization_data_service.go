package service

import (
	"context"

	permify_payload "buf.build/gen/go/permifyco/permify/protocolbuffers/go/base/v1"
	permify_grpc "github.com/Permify/permify-go/grpc"
	"github.com/saahil-mahato/authorization-service/internal/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
)

type AuthorizationService interface {
	CreateConfig(config *models.AuthorizationData) error
	GetConfig(id uint) (*models.AuthorizationData, error)
	ListConfigs() ([]*models.AuthorizationData, error)
	UpdateConfig(config *models.AuthorizationData) error
	DeleteConfig(id uint) error
	CheckAuthorization(url, resource, role, action string) (bool, string, error)
}

type authorizationService struct {
	client *permify_grpc.Client
}

func NewAuthorizationService() AuthorizationService {
	client, err := permify_grpc.NewClient(
		permify_grpc.Config{
			Endpoint: `localhost:3478`,
		},
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil
	}
	return &authorizationService{client: client}
}

func (s *authorizationService) CreateConfig(config *models.AuthorizationData) error {
	value, err := anypb.New(&permify_payload.StringValue{
		Data: config.Value,
	})
	if err != nil {
		return err
	}

	_, err = s.client.Data.Write(context.Background(), &permify_payload.DataWriteRequest{
		TenantId: "t1",
		Metadata: &permify_payload.DataWriteRequestMetadata{
			SchemaVersion: "",
		},
		Tuples: []*permify_payload.Tuple{
			{
				Entity: &permify_payload.Entity{
					Type: config.EntityType,
					Id:   config.EntityID,
				},
				Relation: "viewer",
				Subject: &permify_payload.Subject{
					Type: config.SubjectType,
					Id:   config.SubjectID,
				},
			},
		},
		Attributes: []*permify_payload.Attribute{
			{
				Entity: &permify_payload.Entity{
					Type: config.EntityType,
					Id:   config.EntityID,
				},
				Attribute: config.Attribute,
				Value:     value,
			},
		},
	})

	return err
}

func (s *authorizationService) GetConfig(id uint) (*models.AuthorizationData, error) {
	return nil, nil
}

func (s *authorizationService) ListConfigs() ([]*models.AuthorizationData, error) {
	return nil, nil
}

func (s *authorizationService) UpdateConfig(config *models.AuthorizationData) error {
	return nil
}

func (s *authorizationService) DeleteConfig(id uint) error {
	return nil
}

func (s *authorizationService) CheckAuthorization(url, resource, role, action string) (bool, string, error) {
	return false, "", nil
}
