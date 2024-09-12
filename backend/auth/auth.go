package auth

import (
	"encoding/json"
	"fmt"
	"go-angular/models"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")


func CreateToken(username string) (string, error) {
	
	fmt.Println("Creating token")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"username": username,
		"exp": time.Now().Add(time.Hour*24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		fmt.Printf("Error creating token : %v", err)
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	 }
	
	 if !token.Valid {
		return fmt.Errorf("invalid token")
	 }
	
	 return nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u models.User
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("\nThe user request value %v", u)

	if u.Username == "Chek" && u.Password == "123456" {
		tokenString, err := CreateToken(u.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Errorf("\nNo username found")
			return  
		}
		fmt.Printf("\nToken : %v", tokenString)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
		return 
	}
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}

	tokenString = tokenString[len("Bearer "):]

	err := VerifyToken(tokenString)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return 
	}

	fmt.Fprint(w, "Welcome to the protected route")
}

// ProtectRoute is the middleware that validates the token.
func ProtectRoute(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        tokenString := r.Header.Get("Authorization")

        // Check if the Authorization header is present
        if tokenString == "" {
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprint(w, `{"error": "Missing authorization header"}`)
            return
        }

        // Extract token from "Bearer <token>"
        if strings.HasPrefix(tokenString, "Bearer ") {
            tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        } else {
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprint(w, `{"error": "Invalid authorization header"}`)
            return
        }

        // Verify the token
        err := VerifyToken(tokenString)
        if err != nil {
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprintf(w, `{"error": "Invalid token: %v"}`, err)
            return
        }

        // If the token is valid, proceed to the next handler
        next(w, r)
    }
}