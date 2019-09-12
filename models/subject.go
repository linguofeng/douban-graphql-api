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
	ID            string   `json:"id"`
	Type          string   `json:"type"`
	Title         string   `json:"title"`
	OriginalTitle string   `json:"original_title"`
	Image         Image    `json:"pic"`
	Year          string   `json:"year"`
	Genres        []string `json:"genres"`
	Countries     []string `json:"countries"`
	Durations     []string `json:"durations"`
	Intro         string   `json:"intro"`
	Cover         Cover    `json:"cover"`
	Rating        Rating   `json:"rating"`
	Actors        []Actor  `json:"actors"`
}

type Image struct {
	Small  string `json:"small"`
	Large  string `json:"large"`
	Normal string `json:"normal"`
}
