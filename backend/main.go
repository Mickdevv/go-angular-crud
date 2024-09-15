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

	http.HandleFunc("/api/items", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
        items.GetAllItems(database, w, r) 
    }))
	http.HandleFunc("/api/items/{id}", corsMiddleware(auth.ProtectRoute(func(w http.ResponseWriter, r *http.Request) {
        items.GetUserItem(database, w, r)
    })))

	http.HandleFunc("POST /api/login", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		auth.LoginHandler(database, w, r)
	}))
	http.HandleFunc("POST /api/sign-up", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		auth.SignUp(database, w, r)
	}))
	http.HandleFunc("POST /api/users", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		auth.SignUp(database, w, r)
	}))
	http.HandleFunc("POST /api/users/{id}", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {

		auth.GetUserHandler(database, w, r)
	}))

	fmt.Println("Server started at http://localhost:3000")
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