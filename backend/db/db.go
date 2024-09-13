package db

import (
	"database/sql"
	"go-angular/models"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var db *sql.DB

func InitDb() {
    var err error
    db, err = sql.Open("sqlite3", "./todo.db")
    if err != nil {
        log.Fatal(err)
    }
    // defer DB.Close()

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

    createTodoTable := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        completed BOOLEAN NOT NULL DEFAULT 0,
        user_id INTEGER NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id)
    );`

    if _, err := db.Exec(createUserTable); err != nil {
        log.Fatal(err)
    }

    if _, err := db.Exec(createTodoTable); err != nil {
        log.Fatal(err)
    }
}

// Create User
func CreateUser(user models.User) (int64, error) {
    query := `INSERT INTO users (username, password) VALUES (?, ?)`
    result, err := db.Exec(query, user.Username, user.Password)
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}

func GetUserById() {
    // TODO: Implement this function later
}

func GetUserByUsername(username string) {
    // TODO: Implement this function later
}