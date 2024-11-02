package main

import (
	"fmt"
	"go-angular/auth"
	"go-angular/db"
	"go-angular/items"
	"go-angular/models"
	"net/http"
)

func main() {
	items.Items = append(items.Items, models.Item{Task: "Task 4", Done: true})
	items.Items = append(items.Items, models.Item{Task: "Task 5", Done: false})

	database := db.InitDb()

	http.HandleFunc("/api/items", corsMiddleware(auth.ProtectRoute(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			items.GetAllItems(database, w, r) 
	   	default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	   }
    })))
	http.HandleFunc("GET /api/items/{id}", corsMiddleware(auth.ProtectRoute(func(w http.ResponseWriter, r *http.Request) {
        items.GetUserItem(database, w, r)
    })))
	http.HandleFunc("POST /api/items/{id}", corsMiddleware(auth.ProtectRoute(func(w http.ResponseWriter, r *http.Request) {
        items.AddUserItem(database, w, r)
    })))

	// Apply CORS middleware to /api/login
	http.HandleFunc("/api/login", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		// Handle login requests
		auth.LoginHandler(database, w, r)
	}))

	// http.HandleFunc("POST /api/login", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
	// 	auth.LoginHandler(database, w, r)
	// }))
	http.HandleFunc("POST /api/register", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		auth.Register(database, w, r)
	}))
	http.HandleFunc("POST /api/users", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		auth.Register(database, w, r)
	}))
	http.HandleFunc("POST /api/users/{id}", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		auth.GetUserHandler(database, w, r)
	}))

	fmt.Println("Server started at http://localhost:3000")
    http.ListenAndServe(":3000", nil)
}


// CORS middleware to handle cross-origin requests
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// fmt.Println("CORS Middleware")
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers to allow any origin (or restrict it to your frontend domain)
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept, Origin")
		w.Header().Set("Access-Control-Allow-Credentials", "true") // Enable credentials (cookies, etc.)

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	}
}