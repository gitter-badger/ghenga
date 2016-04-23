package main

import (
	"ghenga/db"
	"log"
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

func (opts *cmdFakedata) Execute(args []string) (err error) {
	dbm, cleanup, e := OpenDB()
	if e != nil {
		return e
	}
	defer CleanupErr(&err, cleanup)

	log.Printf("create %v people...", opts.People)

	err = db.CreateFakePeople(dbm, opts.People)
	if err != nil {
		panic(err)
	}
	log.Printf("done")

	return nil
}
