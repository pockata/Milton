package storage

import (
	"database/sql"
	"milton"
	"os"
)

type DB struct {
	instance *sql.DB
	SQLFile  string
	log      milton.Logger
}

func NewDB(SQLFile string, log milton.Logger) *DB {
	return &DB{
		SQLFile: SQLFile,
		log:     log,
	}
}

func (db *DB) Connect() (*sql.DB, error) {
	dsn := ""

	// create & setup the DB if it hasn't been initialized
	if _, err := os.Stat(db.SQLFile); os.IsNotExist(err) {
		db._create()
	}

	// Open handle to database like normal
	var err error
	db.instance, err = sql.Open("sqlite3", dsn)

	if err != nil {
		return nil, err
	}

	return db.instance, nil
}

func (db *DB) _create() {
	db.log.Info("creating sqlite file", "file", db.SQLFile)

	file, err := os.Create(db.SQLFile) // Create SQLite file
	if err != nil {
		db.log.Error("error creating sqlite file", "err", err.Error())
		os.Exit(1)
	}

	defer file.Close()
}
