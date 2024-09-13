package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func InitDb() {
    db, err := sql.Open("sqlite3", "./todo.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create the necessary tables if they don't exist
    createTables(db)
}

func createTables(db *sql.DB) {
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
func CreateUser(db *sql.DB, username, password string) (int64, error) {
    query := `INSERT INTO users (username, password) VALUES (?, ?, ?)`
    result, err := db.Exec(query, username, password)
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}

func GetUserById()

func GetUserByUsername(username string) {

}