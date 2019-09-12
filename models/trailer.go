package models

// Trailer 预告
type Trailer struct {
	CoverURL string `json:"cover_url"`
	VideoURL string `json:"video_url"`
	Runtime  string `json:"runtime"`
}
