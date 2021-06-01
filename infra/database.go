package infra

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseHandler interface {
	Open() (*sql.DB, error)
	Close(db *sql.DB) error
}

type Database struct {
	DSN string
}

// Open open a database connection
func (d *Database) Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", d.DSN)
	if err != nil {
		return nil, err
	}

	return db, err
}

// Close closes a database connection
func (d *Database) Close(db *sql.DB) error {
	return db.Close()
}

func InitDB(db *sql.DB) error {
	sql := "CREATE TABLE position (ID INTEGER PRIMARY KEY, longitude REAL DEFAULT 0.0, latitude REAL DEFAULT 0.0, time_stamp DATETIME DEFAULT CURRENT_DATETIME, uuid char(36) UNIQUE)"
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sql)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
