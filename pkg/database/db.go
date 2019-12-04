package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	// We blank import sqlite here because gorm needs it.
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const _DefaultDBPath = varBasePath + "db.sqlite3"

var db *DB

// DB represents a persistent storage.
type DB struct {
	*gorm.DB
}

// CreateDB prepares and returns new storage.
//
// It should be run at the start of the program.
func CreateDB(dialect string, args ...interface{}) *DB {
	if len(args) > 0 && args[0] == "" {
		args[0] = _DefaultDBPath
	}
	var err error

	dbase, err := gorm.Open(dialect, args...)
	if err != nil {
		logrus.Fatalf("couldn't open sqlite database %v: %v", args, err)
	}

	dbase.AutoMigrate(&dbUserModel{})
	dbase.AutoMigrate(&dbServerModel{})
	dbase.AutoMigrate(&dbRevokedModel{})
	dbase.AutoMigrate(&dbNetworkModel{})

	dbPTR := &DB{DB: dbase}
	db = dbPTR
	return dbPTR
}

// Cease closes the database.
//
// It should be run at the exit of the program.
func (db *DB) Cease() {
	db.DB.Close()
}
