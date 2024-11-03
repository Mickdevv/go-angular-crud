package db

import (
	"database/sql"
	"fmt"
	"go-angular/models"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var Database *sql.DB

func InitDb() {
    var err error
    Database, err = sql.Open("sqlite3", "./todo.db")
    if err != nil {
        log.Fatal(err)
    }

    // Create the necessary tables if they don't exist
    createTables()
}

func createTables() {
    createUserTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );`

    createItemsTable := `
    CREATE TABLE IF NOT EXISTS items (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        task TEXT NOT NULL,
        done BOOLEAN NOT NULL DEFAULT 0,
        user_id INTEGER NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id)
    );`

    if _, err := Database.Exec(createUserTable); err != nil {
        log.Fatal(err)
    }

    if _, err := Database.Exec(createItemsTable); err != nil {
        log.Fatal(err)
    }
}

// Create User
func CreateUser(user models.User) (int64, error) {
    query := `INSERT INTO users (username, password) VALUES (?, ?)`
    fmt.Println(user)
    result, err := Database.Exec(query, user.Username, user.Password)
    if err != nil {
        fmt.Println(err)
        return 0, err
    }

    return result.LastInsertId()
}

func GetUserById(id int64) (models.User, error) {
    var user models.User

    user.Password = ""

    query := `SELECT id, username FROM users WHERE id = ?`

    err := Database.QueryRow(query, id).Scan(&user.ID, &user.Username)
    
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

func GetUserByUsername(username string) (models.User, error) {
    var user models.User

    user.Password = ""

    query := `SELECT id, username, password FROM users WHERE username = ?`

    err := Database.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
    
    if err != nil {
        if err == sql.ErrNoRows {
            // Handle the case where no user is found
            return user, fmt.Errorf("no user found with username %v", username)
        }
        // Handle any other error that occurred during the query
        return user, err
    }
    fmt.Println("Get user by username", user)

    return user, nil
}

func CreateItem(item models.Item) (int64, error) {
    _, err := GetUserById(int64(item.OwnerId))
    if err != nil {
        fmt.Println(err)
        return 0, err
    }

    query := `INSERT INTO items (task, done, user_id) VALUES (?, ?, ?)`
    fmt.Println(item)
    result, err := Database.Exec(query, item.Task, item.Done, item.OwnerId)
    if err != nil {
        fmt.Println(err)
        return 0, err
    }

    return result.LastInsertId()
}

