package db

import (
	"database/sql"

	// import the sqlite driver
	_ "github.com/mattn/go-sqlite3"

	"github.com/jmoiron/modl"
)

// Init initialises and opens the database.
func Init(dbfile string) (*modl.DbMap, error) {
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
