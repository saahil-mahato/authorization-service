package service

import (
	"context"

	v1 "buf.build/gen/go/permifyco/permify/protocolbuffers/go/base/v1"
	permify_grpc "github.com/Permify/permify-go/grpc"
)

type PermissionService interface {
	CheckAccessControl(tenantID string, metaData *v1.PermissionCheckRequestMetadata, entity *v1.Entity, permission string, subject *v1.Subject) (string, error)
}

type permissionService struct {
	client *permify_grpc.Client
}

func NewPermissionService(client *permify_grpc.Client) PermissionService {
	return &permissionService{client: client}
}

func (s *permissionService) CheckAccessControl(tenantID string, metaData *v1.PermissionCheckRequestMetadata, entity *v1.Entity, permission string, subject *v1.Subject) (string, error) {
	cr, err := s.client.Permission.Check(context.Background(), &v1.PermissionCheckRequest{
		TenantId:   tenantID,
		Metadata:   metaData,
		Entity:     entity,
		Permission: permission,
		Subject:    subject,
	})

	if err != nil {
		return "ERROR", err
	}

	return cr.Can.String(), nil
}
