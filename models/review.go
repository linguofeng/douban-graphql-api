package models

// Review 评论
type Review struct {
	ID            string  `json:"id"`
	Title         string  `json:"title"`
	Type          string  `json:"type"`
	Content       string  `json:"abstract"`
	Rating        Rating  `json:"rating"`
	User          User    `json:"user"`
	CommentsCount int     `json:"comments_count"`
	LinkersCount  int     `json:"linkers_count"`
	ShareCount    int     `json:"timeline_share_count"`
	HTML          string  `json:"content"`
	CreatedAt     string  `json:"create_time"`
	Subject       Subject `json:"subject"`
}
