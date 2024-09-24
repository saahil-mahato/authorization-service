package models

type WriteSchemaPayload struct {
	TenantID string `json:"tenant_id"`
	Schema   string `json:"schema"`
}
