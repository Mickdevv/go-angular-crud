package main

import (
	"fmt"
	"go-angular/auth"
	"go-angular/items"
	"go-angular/models"
	"net/http"
)



func main() {
	items.Items = append(items.Items, models.Item{Task: "Task 4", Done: true})
	items.Items = append(items.Items, models.Item{Task: "Task 5", Done: false})

	http.HandleFunc("GET /api/items", corsMiddleware(auth.ProtectRoute(items.GetAllItems)))
	http.HandleFunc("GET /api/protected", corsMiddleware(auth.ProtectedHandler))
	http.HandleFunc("POST /api/login", corsMiddleware(auth.LoginHandler))

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