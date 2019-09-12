package models

// Celebrity 明星
type Celebrity struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Avatar string   `json:"cover_url"`
	Roles  []string `json:"roles"`
}
