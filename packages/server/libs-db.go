package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	SQLFile  string
	instance *gorm.DB
}

type Unit struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Name string `gorm:"not null"`
	MDNS string `gorm:"unique;not null"`

	Pots []Pot
}

type Pot struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Name string `gorm:"not null"`

	UnitID uint `gorm:"not null"`
}

type Job struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	UnitID uint `gorm:"not null"`
	Unit   Unit

	PotID uint `gorm:"not null"`
	Pot   Pot

	WaterQty  int       `gorm:"not null"`
	StartTime time.Time `gorm:"not null"`
	Status    int       `gorm:"default:1"`
}

type Log struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	UnitID uint
	Unit   Unit

	JobID uint
	Job   Job

	Message string `gorm:"not null"`
}

func (db *DB) setup() {
	var err error

	db.SQLFile = "./sqlite.db"

	// create & setup the DB if it hasn't been initialized
	if _, err := os.Stat(db.SQLFile); os.IsNotExist(err) {
		db._create()
	}

	db.instance, err = gorm.Open(sqlite.Open(db.SQLFile), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Couldn't open %s", db.SQLFile))
	}

	// Migrate the schemas
	db.instance.AutoMigrate(&Unit{})
	db.instance.AutoMigrate(&Pot{})
	db.instance.AutoMigrate(&Job{})
	db.instance.AutoMigrate(&Log{})
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
