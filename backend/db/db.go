package db

import (
	"database/sql"
	"fmt"
	"go-angular/models"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)


func InitDb() *sql.DB {
    database, err := sql.Open("sqlite3", "./todo.db")
    if err != nil {
        log.Fatal(err)
    }


    // Create the necessary tables if they don't exist
    createTables(database)

    return database
}

func createTables(database *sql.DB) {
    createUserTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );`

    createTodoTable := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        completed BOOLEAN NOT NULL DEFAULT 0,
        user_id INTEGER NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id)
    );`

    if _, err := database.Exec(createUserTable); err != nil {
        log.Fatal(err)
    }

    if _, err := database.Exec(createTodoTable); err != nil {
        log.Fatal(err)
    }
}

// Create User
func CreateUser(database *sql.DB, user models.User) (int64, error) {
    query := `INSERT INTO users (username, password) VALUES (?, ?)`
    fmt.Println(user)
    result, err := database.Exec(query, user.Username, user.Password)
    if err != nil {
        fmt.Println(err)
        return 0, err
    }

    return result.LastInsertId()
}

func GetUserById(database *sql.DB, id int64) (models.User, error) {
    var user models.User

    user.Password = ""

    query := `SELECT id, username FROM users WHERE id = ?`

    err := database.QueryRow(query, id).Scan(&user.ID, &user.Username)
    
    if err != nil {
        if err == sql.ErrNoRows {
            // Handle the case where no user is found
            return user, fmt.Errorf("no user found with id %d", id)
        }
        // Handle any other error that occurred during the query
        return user, err
    }

    return user, nil
}

func GetUserByUsername(database *sql.DB, username string) (models.User, error) {
    var user models.User

    user.Password = ""

    query := `SELECT id, username, password FROM users WHERE username = ?`

    err := database.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
    
    if err != nil {
        if err == sql.ErrNoRows {
            // Handle the case where no user is found
            return user, fmt.Errorf("no user found with username %v", username)
        }
        // Handle any other error that occurred during the query
        return user, err
    }
    fmt.Println(user)

    return user, nil
}