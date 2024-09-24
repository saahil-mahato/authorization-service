package service

import (
	"context"
	"strings"

	v1 "buf.build/gen/go/permifyco/permify/protocolbuffers/go/base/v1"
	permify_grpc "github.com/Permify/permify-go/grpc"
	"google.golang.org/protobuf/types/known/anypb"
)

type DataService interface {
	WriteAuthorizationData(tenantID, relation, attribute, value string, metaData *v1.DataWriteRequestMetadata, entity *v1.Entity, subject *v1.Subject) (*v1.DataWriteResponse, error)
}

type dataService struct {
	client *permify_grpc.Client
}

func NewDataService(client *permify_grpc.Client) DataService {
	return &dataService{client: client}
}

func (s *dataService) WriteAuthorizationData(tenantID, relation, attribute, value string, metaData *v1.DataWriteRequestMetadata, entity *v1.Entity, subject *v1.Subject) (*v1.DataWriteResponse, error) {
	anyValue, err := anypb.New(&v1.StringValue{
		Data: value,
	})
	if err != nil {
		return nil, err
	}

	request := &v1.DataWriteRequest{
		TenantId: tenantID,
		Metadata: metaData,
		Tuples: []*v1.Tuple{
			{
				Entity:   entity,
				Relation: relation,
				Subject:  subject,
			},
		},
	}

	if len(strings.TrimSpace(attribute)) > 0 {
		request.Attributes = []*v1.Attribute{
			{
				Entity:    entity,
				Attribute: attribute,
				Value:     anyValue,
			},
		}
	}

	cr, err := s.client.Data.Write(context.Background(), request)

	return cr, err
}
