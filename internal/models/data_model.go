package models

type AuthorizationDataPayload struct {
	SchemaVersion   string `json:"schema_version"`
	EntityType      string `json:"entity_type"`
	EntityID        string `json:"entity_id"`
	Relation        string `json:"relation"`
	SubjectType     string `json:"subject_type"`
	SubjectID       string `json:"subject_id"`
	SubjectRelation string `json:"subject_relation"`
	Attribute       string `json:"attribute"`
	Value           string `json:"value"`
}
