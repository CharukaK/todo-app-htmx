package main

import (
	"database/sql"
	"os"

	"github.com/charukak/todo-app-htmx/backend/internal/resource/ping"
	"github.com/charukak/todo-app-htmx/backend/internal/resource/todo"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_NAME = "todo.db"
	DB_PATH = "./" + DB_NAME
)

func main() {
	createDBIfNotExist()
	db := openDB()
	createTableIfNotExist(db)

	r := gin.Default()
	ping.RegisterPingHandler(r)
	todo.RegisterTodoHandler(r, db)
	r.Run() // listen and serve on
}

func createDBIfNotExist() {
	if _, err := os.Stat(DB_PATH); os.IsNotExist(err) {
		file, err := os.Create(DB_PATH)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}
}

func openDB() *sql.DB {
	db, err := sql.Open("sqlite3", DB_PATH)

	if err != nil {
		panic(err)
	}

	return db
}

func createTableIfNotExist(db *sql.DB) {
	sqlStmt := `
	create table if not exists todos (id integer not null primary key, title text, description text, status bool);
	`
	_, err := db.Exec(sqlStmt)

	if err != nil {
		panic(err)
	}
}
