package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type item struct {
	Task string `json:"task"`
	Done bool `json:"done"`
}

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Items []item `json:"items"`
}

var items = []item{
    {Task: "Task 1", Done: false},
    {Task: "Task 2", Done: true},
    {Task: "Task 3", Done: false},
}


func main() {
	items = append(items, item{Task: "Task 4", Done: true})
	items = append(items, item{Task: "Task 5", Done: false})

	http.HandleFunc("GET /api/items", corsMiddleware(getAllItems))

	fmt.Println("Server started at :3000")
    http.ListenAndServe(":3000", nil)
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println(items)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(items)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}


// CORS middleware to handle cross-origin requests
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins. You can restrict it to your Angular app's URL
		w.Header().Set("Access-Control-Allow-Methods", "*") // Allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "*")

		// If it's an OPTIONS request, just return
		if r.Method == "OPTIONS" {
			return
		}

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	}
}