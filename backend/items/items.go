package items

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-angular/models"
)


var Items = []models.Item{
    {Task: "Task 1", Done: false},
    {Task: "Task 2", Done: true},
    {Task: "Task 3", Done: false},
}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println(Items)
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
	fmt.Println(Items)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(Items)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}