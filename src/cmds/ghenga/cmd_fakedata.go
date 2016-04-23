package main

import (
	"errors"
	"fmt"
	"ghenga/db"
	"log"
	"os"
)

type cmdFakedata struct {
	People int `short:"p" long:"people" default:"500" description:"Number of fake people profiles to create"`
}

func init() {
	_, err := parser.AddCommand("fakedata",
		"fill db with fake dat",
		"This command fills the data with fake information that looks real",
		&cmdFakedata{})
	if err != nil {
		panic(err)
	}
}

func (opts *cmdFakedata) Execute(args []string) error {
	if len(args) != 1 {
		return errors.New("no database file specified")
	}

	dbfile := os.Args[0]
	dbmap, err := db.Init(dbfile)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := dbmap.Db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Printf("create %v people...", opts.People)

	err = db.CreateFakePeople(dbmap, opts.People)
	if err != nil {
		panic(err)
	}
	fmt.Printf("done\n")

	return nil
}
