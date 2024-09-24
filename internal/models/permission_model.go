package models

type CheckAccessControlPayload struct {
	SnapToken     string `json:"snap_token"`
	SchemaVersion string `json:"schema_version"`
	Depth         int32  `json:"depth"`
	EntityType    string `json:"entity_type"`
	EntityID      string `json:"entity_id"`
	Permission    string `json:"permission"`
	SubjectType   string `json:"subject_type"`
	SubjectID     string `json:"subject_id"`
}
