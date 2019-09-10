package models

// Subject 主题
type Subject struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Type   string `json:"type"`
	Cover  Cover  `json:"cover"`
	Rating Rating `json:"rating"`
}

type SubjectDetail struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	OriginalTitle string `json:"original_itle"`
	Type          string `json:"type"`
	Intro         string `json:"intro"`
	Cover         Cover  `json:"cover"`
	Rating        Rating `json:"rating"`
}
