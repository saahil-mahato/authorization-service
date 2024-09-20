package models

type AuthorizationConfig struct {
	ID        int    `json:"id"`
	URL       string `json:"url"`
	Resource  string `json:"resource"`
	Role      string `json:"role"`
	CanCreate bool   `json:"can_create"`
	CanRead   bool   `json:"can_read"`
	CanUpdate bool   `json:"can_update"`
	CanDelete bool   `json:"can_delete"`
}
