package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-angular/db"
	"go-angular/models"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

// Define the struct that mirrors the expected JSON structure
type SignUpRequest struct {
    Username  string `json:"username"`
    Password1 string `json:"password1"`
    Password2 string `json:"password2"`
}

func HashPassword(password string) (string, error) {
    // Generate a salted hash for the password
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }

    // Return the hashed password as a string
    return string(hash), nil
}

// ComparePasswords compares a plain password with a hashed password.
func ComparePasswords(hashedPassword, password string) bool {
    // Compare the hashed password with the plain password
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}

func SignUp(database *sql.DB, w http.ResponseWriter, r *http.Request) {

	var req SignUpRequest

	// Parse the JSON body into the SignUpRequest struct
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		if err == io.EOF {
			http.Error(w, "Request body is empty", http.StatusBadRequest)
		} else {
			http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		}
		return
	}

	// Close the request body when done (best practice)
	defer r.Body.Close()
	
	fmt.Println("Sign-up")

	if req.Password1 != req.Password2 {
		fmt.Fprint(w, "Passwords do not match")
		return 
	} else if len(req.Password1) < 5 {
		fmt.Fprint(w, "Password must be longer than 5 characters")
		return 
	}

	hash, err := HashPassword(req.Password1)

	user := models.User{
		Username: req.Username,
		Password: hash,
	}

	fmt.Println(user)

	userID, err := db.CreateUser(database, user)

	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	newUser, err := db.GetUserById(database, userID)

	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	fmt.Println("New user : ", newUser.ID, newUser.Username)

	fmt.Fprintf(w, "Received POST request. Username: %s, Password1: %s, Password2: %s, Hash: %s. User Id in the database : %v", req.Username, req.Password1, req.Password2, hash, userID)
}

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

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	 }
	
	 if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	 }

	 // Extract claims (payload) from the token
	 if claims, ok := token.Claims.(jwt.MapClaims); ok {
        return claims, nil
    }
	
	return nil, fmt.Errorf("could not extract claims")
}

func LoginHandler(database *sql.DB, w http.ResponseWriter, r *http.Request) {
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

	claims, err := VerifyToken(tokenString)

	fmt.Println(claims)

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
        claims, err := VerifyToken(tokenString)
        if err != nil {
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprintf(w, `{"error": "Invalid token: %v"}`, err)
            return
        }

		fmt.Println(claims["username"])

        // If the token is valid, proceed to the next handler
        next(w, r)
    }
}