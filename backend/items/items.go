package items

import (
	"encoding/json"
	"fmt"
	"go-angular/auth"
	"go-angular/db"
	"go-angular/models"
	"net/http"
)


var Items = []models.Item{
    {Task: "Task 1", Done: false},
    {Task: "Task 2", Done: true},
    {Task: "Task 3", Done: false},
}

type AddItemRequest struct {
	Task  string `json:"task"`
	Done  bool `json:"done"`
}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(Items)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func GetUserItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, err := auth.CheckToken(r)
	if err != nil {
		fmt.Println("Token check failed", err)
		return
	}

	items, err := db.GetUserItems(user.ID)

	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func GetUserItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	return 
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, err := auth.CheckToken(r)
	if err != nil {
		fmt.Println("Token check failed", err)
		return
	}

	var newItem AddItemRequest
	json.NewDecoder(r.Body).Decode(&newItem)

	newItemAdd := models.Item{Task: newItem.Task, Done: newItem.Done, OwnerId: user.ID} 

	newItemId, err := db.CreateItem(newItemAdd)
	if err != nil {
		fmt.Println("Unable to add item to database", err)
		return
	}

	newItemAdd.OwnerId = uint64(newItemId)
	fmt.Println(newItemAdd)
	return 
}