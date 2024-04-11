package db

import (
	"database/sql"
	"fmt"
	"milton/core/ports"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	SQLFile string
	log     ports.Logger
}

func NewDB(SQLFile string, log ports.Logger) *DB {
	return &DB{
		SQLFile: SQLFile,
		log:     log,
	}
}

func (db *DB) Connect() (*sql.DB, error) {
	// create & setup the DB if it hasn't been initialized
	if _, err := os.Stat(db.SQLFile); os.IsNotExist(err) {
		db._create()
	}

	dsn := fmt.Sprintf("file:%s", db.SQLFile)

	// Open handle to database like normal
	return sql.Open("sqlite3", dsn)
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
