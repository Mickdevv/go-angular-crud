package models

type Item struct {
	ID uint64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Done bool `json:"done"`
	OwnerId uint64 `json:"ownerId"`
}

type User struct {
	ID uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}
