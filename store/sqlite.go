package store

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDBConn() {
	db, err := sql.Open("sqlite3", "./resources/mindStream.db")
	if err != nil {
		fmt.Println("connect failed")
	} else {
		DB = db
		fmt.Println("connect to sqlite success")
	}
}
