package models

// Review 评论
type Review struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	Content       string `json:"abstract"`
	Rating        Rating `json:"rating"`
	User          User   `json:"user"`
	CommentsCount int    `json:"comments_count"`
}
