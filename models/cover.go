package models

// Cover 封面
type Cover struct {
	URL    string `json:"url"`
	Shape  string `json:"shape"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
