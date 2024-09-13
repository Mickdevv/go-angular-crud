package main

import (
	"fmt"
	"go-angular/auth"
	"go-angular/db"
	"go-angular/items"
	"go-angular/models"
	"net/http"
	// SQLite driver
)

func main() {
	items.Items = append(items.Items, models.Item{Task: "Task 4", Done: true})
	items.Items = append(items.Items, models.Item{Task: "Task 5", Done: false})

	database := db.InitDb()

	http.HandleFunc("GET /api/items", corsMiddleware(auth.ProtectRoute(items.GetAllItems)) {
		// TODO : add explicit database reference to be used in handler functions
	})

	http.HandleFunc("POST /api/login", corsMiddleware(auth.LoginHandler))
	http.HandleFunc("POST /api/sign-up", corsMiddleware(auth.SignUp))

	fmt.Println("Server started at :3000")
    http.ListenAndServe(":3000", nil)
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