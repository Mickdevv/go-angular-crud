package models

type Item struct {
	Task string `json:"task"`
	Done bool `json:"done"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ItemList struct {
	Owner User `json:"user"`
	Items []Item `json:"items"`
}