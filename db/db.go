package db

import (
	"database/sql"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> 8e986f10914499cf35aed1d00ed73fe98c9c2b56

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("could not establish connection")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

<<<<<<< HEAD
=======
	createUsersTable := `
	   	CREATE TABLE IF NOT EXISTS users (
	   		id INTEGER PRIMARY KEY AUTOINCREMENT,
	   		email TEXT NOT NULL UNIQUE,
	   		password TEXT NOT NULL
	   		)
	   	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("could not create users table")
	}

>>>>>>> 8e986f10914499cf35aed1d00ed73fe98c9c2b56
	createEventsTable := `
	   	CREATE TABLE IF NOT EXISTS events (
	   		id INTEGER PRIMARY KEY AUTOINCREMENT,
	   		name TEXT NOT NULL,
	   		description TEXT NOT NULL,
	   		location TEXT NOT NULL,
	   		dateTime DATETIME NOT NULL,
<<<<<<< HEAD
	   		user_id INTEGER
	   		)
	   	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		fmt.Println(err)
=======
	   		user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id) 
	   		)
	   	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("could not create events table")
>>>>>>> 8e986f10914499cf35aed1d00ed73fe98c9c2b56
	}

}
