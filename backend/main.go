package main

import (
	"fmt"
	"go-angular/auth"
	"go-angular/db"
	"go-angular/items"
	"log"
	"net/http"
)

func main() {

	db.InitDb()

	http.HandleFunc("GET /api/items/", corsMiddleware(auth.ProtectRoute(func(w http.ResponseWriter, r *http.Request) {
		items.GetUserItems(w, r) 
    })))
	http.HandleFunc("POST /api/items/", corsMiddleware(auth.ProtectRoute(func(w http.ResponseWriter, r *http.Request) {
		items.AddItem(w, r)
    })))

	http.HandleFunc("GET /api/items/{id}/", corsMiddleware(auth.ProtectRoute(func(w http.ResponseWriter, r *http.Request) {
		items.GetUserItem(w, r)
	})))	
	http.HandleFunc("PUT /api/items/{id}/", corsMiddleware(auth.ProtectRoute(func(w http.ResponseWriter, r *http.Request) {
		items.UpdateItem(w, r)
	})))
	http.HandleFunc("DELETE /api/items/{id}/", corsMiddleware(auth.ProtectRoute(func(w http.ResponseWriter, r *http.Request) {
		items.RemoveItem(w, r)
	})))
	http.HandleFunc("OPTIONS /api/items/{id}/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	http.HandleFunc("OPTIONS /api/items/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	http.HandleFunc("OPTIONS /api/login/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	http.HandleFunc("OPTIONS /api/register/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Apply CORS middleware to /api/login
	http.HandleFunc("POST /api/login/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		// Handle login requests
		auth.LoginHandler(w, r)
	}))
	http.HandleFunc("POST /api/register/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		auth.Register(w, r)
	}))
	http.HandleFunc("POST /api/users/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		auth.Register(w, r)
	}))
	http.HandleFunc("POST /api/users/{id}/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		auth.GetUserHandler(w, r)
	}))

	fmt.Println("Server started at http://localhost:3000/")
    http.ListenAndServe(":3000", nil)
}


// CORS middleware to handle cross-origin requests
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// fmt.Println("CORS Middleware")
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers to allow any origin (or restrict it to your frontend domain)
		origin := r.Header.Get("Origin")

		log.Println("Request Origin:", origin)
		log.Println("Request Method:", r.Method)
		log.Println("Request Headers:", r.Header)


		// Allow requests from specific origins (e.g., frontend at http://localhost:4200)
		// if origin == "http://localhost:4200" || origin == "http://localhost:3000" || origin == "http://127.0.0.1:3000" || origin == "http://127.0.0.1:4200" {
		// }
		w.Header().Set("Access-Control-Allow-Origin", origin)
		
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow methods and headers for cross-origin requests
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH, HEAD")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept, Origin, X-Csrf-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true") // Allow credentials (cookies, etc.)

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	}
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}