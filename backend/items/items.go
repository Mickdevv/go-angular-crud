package items

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-angular/auth"
	"go-angular/db"
	"go-angular/models"
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
	fmt.Println(Items)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(Items)
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

	tokenCookie, err := r.Cookie("jwt_token")
	if err != nil {
		fmt.Println("Error occured while reading cookie")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"error": "No valid token cookie found"}`)
		return
	}
	// Verify the token
	claims, err := auth.VerifyToken(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"error": "Invalid token: %v"}`, err)
		return
	}

	var newItem AddItemRequest
	json.NewDecoder(r.Body).Decode(&newItem)

	var OwnerId uint64
	if idFloat, ok := claims["id"].(float64); ok {
		OwnerId = uint64(idFloat)
	
	// If claims["id"] is neither a string nor a float64, handle the error
	} else {
		fmt.Println("id is not in a recognized format (string or float64)")
		return
	}

	newItemAdd := models.Item{Task: newItem.Task, Done: newItem.Done, OwnerId: OwnerId} 

	newItemId, err := db.CreateItem(newItemAdd)
	if err != nil {
		fmt.Println("Unable to add item to database", err)
		return
	}

	newItemAdd.OwnerId = uint64(newItemId)
	fmt.Println(newItemAdd)
	return 
}