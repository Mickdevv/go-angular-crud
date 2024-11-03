package models

type Item struct {
	Task string `json:"task"`
	Done bool `json:"done"`
	OwnerId uint64 `json:"ownerId"`
}

type User struct {
	ID uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

// type ItemList struct {
// 	OwnerId uint32 `json:"ownerId"`
// 	Items []Item `json:"items"`
// }