package storage

import (
	"database/sql"
	"log"
	"milton"
	"os"
)

type DB struct {
	instance *sql.DB
	SQLFile  string
	log      *milton.Logger
}

func NewDB(SQLFile string) *DB {
	return &DB{
		SQLFile: SQLFile,
	}
}

func (db *DB) Connect() (*sql.DB, error) {
	var err error
	dsn := ""

	// create & setup the DB if it hasn't been initialized
	if _, err := os.Stat(db.SQLFile); os.IsNotExist(err) {
		db._create()
	}

	// Open handle to database like normal
	db.instance, err = sql.Open("sqlite3", dsn)

	if err != nil {
		return nil, err
	}

	return db.instance, nil
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
