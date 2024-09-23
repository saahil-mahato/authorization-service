package models

type AuthorizationSchema struct {
	TenantID string `json:"tenant_id"`
	Schema   string `json:"schema"`
}

type AuthorizationData struct {
	TenantID    string `json:"tenant_id"`
	TenantName  string `json:"tenant_name"`
	EntityType  string `json:"entity_type"`
	EntityID    string `json:"entity_id"`
	Relation    string `json:"relation"`
	SubjectType string `json:"subject_type"`
	SubjectID   string `json:"subject_id"`
	Attribute   string `json:"attribute"`
	Value       string `json:"value"`
}
