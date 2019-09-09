package models

// Rating 评分
type Rating struct {
	Count int     `json:"count"`
	Max   int     `json:"max"`
	Value float32 `json:"value"`
}
