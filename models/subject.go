package models

// Subject 主题
type Subject struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Cover Cover  `json:"cover"`
}
