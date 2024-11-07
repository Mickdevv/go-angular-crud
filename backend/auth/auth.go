package auth

import (
	"encoding/json"
	"fmt"
	"go-angular/db"
	"go-angular/models"
	"io"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

// Define the struct that mirrors the expected JSON structure
type RegisterRequest struct {
    Username  string `json:"username"`
    Password1 string `json:"password1"`
    Password2 string `json:"password2"`
}

type LoginRequest struct {
    Username  string `json:"username"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Access  string `json:"access"`
    Refresh string `json:"refresh"`
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

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) (error) {
	id := r.PathValue("id")
	
	var user models.User

	// Convert string to int64
	userID, err := strconv.ParseInt(id, 10, 64)  // Base 10, 64-bit size
	if err != nil {
		fmt.Println("Error converting string to int64:", err)
		return err
	}

	user, err = db.GetUserById(userID)

	if err != nil {
		return err
	}

	fmt.Println(user)

	err = json.NewEncoder(w).Encode(user)

	return nil
}

func Register(w http.ResponseWriter, r *http.Request) {

	var req RegisterRequest

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
	// fmt.Println(req.Password1, hash, "--")

	user := models.User{
		Username: req.Username,
		Password: hash,
	}

	userID, err := db.CreateUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser, err := db.GetUserById(userID)

	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	fmt.Println("New user : ", newUser.ID, newUser.Username)
	fmt.Fprintf(w, "Received POST request. Username: %s, Password1: %s, Password2: %s, Hash: %s. User Id in the database : %v", req.Username, req.Password1, req.Password2, hash, userID)
}

func CreateToken(username string, id uint64) (string, time.Time, error) {

	expirationTime := time.Now().Add(time.Hour*24)
	
	fmt.Println("Creating token")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"username": username,
		"id": id,
		"exp": expirationTime.Unix(),
	})
	

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		fmt.Printf("Error creating token : %v", err)
		return "", expirationTime, err
	}
	return tokenString, expirationTime, nil
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var requestUser LoginRequest
	var loginResponse LoginResponse
	json.NewDecoder(r.Body).Decode(&requestUser)
	fmt.Printf("\nThe user request value %v\n", requestUser)

	databaseUser, err := db.GetUserByUsername(requestUser.Username)
	if err != nil {
		fmt.Fprint(w, "No user found with that username")
		return
	}

	hashedPassword, err := HashPassword(requestUser.Password)
	fmt.Println(requestUser.Password, hashedPassword, "--")

	if err != nil {
		fmt.Fprint(w, "Password error")
		return
	}

	if ComparePasswords(databaseUser.Password, requestUser.Password) {
		tokenString, tokenExpiration, err := CreateToken(requestUser.Username, databaseUser.ID)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "\nNo username found")
			return
		}

		fmt.Printf("\nToken : %v\n", tokenString)
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt_token",
			Value:    tokenString,
			Expires:  tokenExpiration,
			HttpOnly: false,         // Ensures cookie is inaccessible to JavaScript
			SameSite: http.SameSiteNoneMode, 
			Path:     "/",
			Secure: true,
		})


		loginResponse.Access = tokenString
		loginResponse.Refresh = tokenString
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(loginResponse)
		// fmt.Fprint(w, tokenString)
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
        // tokenString := r.Header.Get("Authorization")

		tokenCookie, err := r.Cookie("jwt_token")
		if err != nil {
			fmt.Println("Error occured while reading cookie")
			w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprint(w, `{"error": "No valid token cookie found"}`)
            return
		}
		// fmt.Println("\nPrinting cookie with name as token")
		// fmt.Println(tokenCookie.Value)

        // Check if the Authorization header is present
        // if tokenString == "" {
        //     w.WriteHeader(http.StatusUnauthorized)
        //     fmt.Fprint(w, `{"error": "Missing authorization header"}`)
        //     return
        // }

        // Extract token from "Bearer <token>"
        // if strings.HasPrefix(tokenString, "Bearer ") {
        //     tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        // } else {
        //     w.WriteHeader(http.StatusUnauthorized)
        //     fmt.Fprint(w, `{"error": "Invalid authorization header"}`)
        //     return
        // }

        // Verify the token
        claims, err := VerifyToken(tokenCookie.Value)
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

func CheckToken(r *http.Request) (models.User, error) {
	var user models.User
	tokenCookie, err := r.Cookie("jwt_token")
	if err != nil {
		return models.User{}, err
	}
	// Verify the token
	claims, err := VerifyToken(tokenCookie.Value)
	if err != nil {
		return models.User{}, err
	}

	if id, ok := claims["id"].(float64); ok {
		user.ID = uint64(id)
	} else {
		fmt.Println("id is not in a recognized format (string or float64)")
		return models.User{}, err
	}

	if username, ok := claims["username"].(string); ok {
		user.Username = string(username)
	} else {
		fmt.Println("id is not in a recognized format (string or float64)")
		return models.User{}, err
	}

	return user, nil
}