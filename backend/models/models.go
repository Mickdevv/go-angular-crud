package models

type Item struct {
	ID uint64 `json:"id"`
	Task string `json:"task"`
	Done bool `json:"done"`
	OwnerId uint64 `json:"ownerId"`
}

type User struct {
	ID uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}
