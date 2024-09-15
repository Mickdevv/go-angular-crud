package items

import (
	"database/sql"
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

func GetAllItems(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println(db)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(Items)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func GetUserItems(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println(Items)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(Items)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func GetUserItem(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println(Items)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(Items)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}