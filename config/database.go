/*
package config

import (

	"database/sql"
	"log"
	"net/http"

)

	func InitDB() *sql.DB {
		db, err := sql.Open("sqlite3", "students.db")
		if err != nil {
			log.Fatal(err)
		}
		//Students Table
		db.Exec(`
		CREATE TABLE IF NOT EXISTS students (
			id TEXT PRIMARY KEY,
			name TEXT,
			major TEXT,
			gpa REAL
		)
		`)

		log.Println("Server started at :8080")
		http.ListenAndServe(":8080", nil)

		return db
	}
*/
package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./students.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS students (
		id TEXT PRIMARY KEY,
		name TEXT,
		major TEXT,
		gpa REAL
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
