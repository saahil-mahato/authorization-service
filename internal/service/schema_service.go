package service

import (
	"context"

	v1 "buf.build/gen/go/permifyco/permify/protocolbuffers/go/base/v1"
	permify_grpc "github.com/Permify/permify-go/grpc"
)

type SchemaService interface {
	WriteSchema(tenantID, schema string) (*v1.SchemaWriteResponse, error)
	ListSchema(tenantID, continuousToken string, pageSize uint32) (*v1.SchemaListResponse, error)
	ReadSchema(tenantID, schemaVersion string) (*v1.SchemaReadResponse, error)
}

type schemaService struct {
	client *permify_grpc.Client
}

func NewSchemaService(client *permify_grpc.Client) SchemaService {
	return &schemaService{client: client}
}

func (s *schemaService) WriteSchema(tenantID, schema string) (*v1.SchemaWriteResponse, error) {
	sr, err := s.client.Schema.Write(context.Background(), &v1.SchemaWriteRequest{
		TenantId: tenantID,
		Schema:   schema,
	})

	return sr, err
}

func (s *schemaService) ListSchema(tenantID, continuousToken string, pageSize uint32) (*v1.SchemaListResponse, error) {
	sr, err := s.client.Schema.List(context.Background(), &v1.SchemaListRequest{
		TenantId:        tenantID,
		ContinuousToken: continuousToken,
		PageSize:        pageSize,
	})

	return sr, err
}

func (s *schemaService) ReadSchema(tenantID, schemaVersion string) (*v1.SchemaReadResponse, error) {
	sr, err := s.client.Schema.Read(context.Background(), &v1.SchemaReadRequest{
		TenantId: tenantID,
		Metadata: &v1.SchemaReadRequestMetadata{
			SchemaVersion: schemaVersion,
		},
	})

	return sr, err
}
