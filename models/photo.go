package models

type PhotoResp struct {
	Start int      `json:"start"`
	Count int      `json:"count"`
	Total int      `json:"total"`
	Photo []*Photo `json:"photos"`
}

// Photo 图片
type Photo struct {
	ID    string     `json:"id"`
	Type  string     `json:"type"`
	Image PhotoImage `json:"image"`
}

type PhotoImage struct {
	Small  PhotoImageDetail `json:"small"`
	Large  PhotoImageDetail `json:"large"`
	Normal PhotoImageDetail `json:"normal"`
}

type PhotoImageDetail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
