package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type DB struct {
	SQLFile  string
	instance *sql.DB
}

func (db *DB) init() {
	var err error

	db.SQLFile = "./sqlite.db"

	// create & setup the DB if it hasn't been initialized
	if _, err := os.Stat(db.SQLFile); os.IsNotExist(err) {
		db._create()
		db._setup()
	}

	db.instance, err = sql.Open("sqlite3", db.SQLFile) // Open the created SQLite File
	if err != nil {
		panic(fmt.Sprintf("Couldn't open %s", db.SQLFile))
	}
}

func (db *DB) _create() {
	log.Printf("Creating sqlite file %s\n", db.SQLFile)

	file, err := os.Create(db.SQLFile) // Create SQLite file

	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	log.Printf("%s created\n", db.SQLFile)
}

func (db *DB) close() {
	db.instance.Close()
}

func (db *DB) _setup() {
	createStudentTableSQL := `CREATE TABLE log (
        "ID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        "unit" TEXT,
        "msg" TEXT
      );`

	db._runQuery(createStudentTableSQL)
}

func (db *DB) _runQuery(query string) {
	log.Println("Create log table...")
	statement, err := db.instance.Prepare(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec() // Execute SQL Statements
	log.Println("log table created")
}

func (db DB) log(unit string, msg string) {
	query := `INSERT INTO log(unit, msg) VALUES (?, ?)`
	statement, err := db.instance.Prepare(query)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(unit, msg)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
