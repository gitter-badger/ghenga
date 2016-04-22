package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jmoiron/modl"
)

// initDB initialises and opens the database.
func initDB(dbfile string) (*modl.DbMap, error) {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return nil, err
	}

	dbmap := modl.NewDbMap(db, modl.SqliteDialect{})
	dbmap.AddTable(Person{}, "people").SetKeys(true, "ID")

	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
	}

	return dbmap, nil
}
