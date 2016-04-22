package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/jmoiron/modl"
)

var DBM *modl.DbMap

func main() {
	var err error
	DBM, err = initDB("db/test.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := DBM.Db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	_, err = parser.Parse()
	if e, ok := err.(*flags.Error); ok && e.Type == flags.ErrHelp {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	if err != nil {
		os.Exit(1)
	}
}
