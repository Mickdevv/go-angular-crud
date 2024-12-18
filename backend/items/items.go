package items

import (
	"encoding/json"
	"fmt"
	"go-angular/auth"
	"go-angular/db"
	"go-angular/models"
	"net/http"
	"strconv"
)

type AddItemRequest struct {
	Title  string `json:"title"`
	Description  string `json:"description"`
	Done  bool `json:"done"`
}
type UpdateItemRequest struct {
	Title  string `json:"title"`
	Description  string `json:"description"`
	Done  bool `json:"done"`
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
		http.Error(w, `{"error":"Error encoding JSON"}`, http.StatusInternalServerError)
		return
	}
	return 
}

func GetUserItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url_id := r.PathValue("id")
	item_id, _ := strconv.Atoi(url_id)

	user, err := auth.CheckToken(r)
	if err != nil {
		fmt.Println("Token check failed", err)
		return
	}

	item, err := db.GetUserItem(user.ID, uint64(item_id))
	if err != nil {
		fmt.Println(err)
		http.Error(w, `{"error":"Error getting item"}`, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, `{"error":"Error encoding JSON"}`, http.StatusInternalServerError)
		return
	}
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

	newItemAdd := models.Item{Title: newItem.Title, Description: newItem.Description, Done: newItem.Done, OwnerId: user.ID} 

	newItemId, err := db.CreateItem(newItemAdd)
	if err != nil {
		fmt.Println("Unable to add item to database", err)
		return
	}

	newItemAdd.OwnerId = uint64(newItemId)
	fmt.Println(newItemAdd)
	return 
}

func RemoveItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(1)

	url_id := r.PathValue("id")
	item_id, _ := strconv.Atoi(url_id)

	user, err := auth.CheckToken(r)
	if err != nil {
		fmt.Println("Token check failed", err)
		return
	}

	item, err := db.GetUserItem(user.ID, uint64(item_id))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting item", http.StatusInternalServerError)
		return
	}

	if item.OwnerId != user.ID {
		http.Error(w, `{"error":"User unauthorized or item does not exist"}`, http.StatusUnauthorized)
		return 
	}
	fmt.Println(2)
	err = db.RemoveItem(int64(item_id))
	if err != nil {
		fmt.Println("Delete failed", err)
		return
	}

	err = json.NewEncoder(w).Encode(item.ID)
	if err != nil {
		fmt.Println("Encoding item failed", err)
		return
	}
	return
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(3)

	url_id := r.PathValue("id")
	item_id, _ := strconv.Atoi(url_id)

	user, err := auth.CheckToken(r)
	if err != nil {
		fmt.Println("Token check failed", err)
		return
	}

	item, err := db.GetUserItem(user.ID, uint64(item_id))
	if err != nil {
		fmt.Println(err)
		http.Error(w, `{"error":"Error getting item"}`, http.StatusInternalServerError)
		return
	}

	if item.OwnerId != user.ID {
		http.Error(w, `{"error":"User unauthorized or item does not exist"}`, http.StatusUnauthorized)
		return 
	}
	fmt.Println(4)
	
	var updateRequest AddItemRequest
	json.NewDecoder(r.Body).Decode(&updateRequest)

	updatedItem := models.Item{ID: item.ID, OwnerId: item.OwnerId, Description: updateRequest.Description, Done: updateRequest.Done, Title: updateRequest.Title}
 	db.UpdateItem(updatedItem)
	return
}