package database

import (
	"log"

	"gorm.io/driver/sqlite"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateBb bool
	Env           string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {

	dbInstance := NewDb()
	dbInstance.Env = "Test"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory"
	dbInstance.AutoMigrateBb = true
	dbInstance.Debug = true

	connection, err := dbInstance.Connect()

	if err != nil {
		log.Fatalf("Test DB error: %v", err)
	}
	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	if d.Env != "Test" {
		d.Db, err = gorm.Open(postgres.Open(d.Dsn), &gorm.Config{})
	} else {
		d.Db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	}
	return nil, err
}
