package models

// User 用户
type User struct {
	ID     string `json:"id"`
	UID    string `json:"uid"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
