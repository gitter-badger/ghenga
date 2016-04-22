package main

import (
	"fmt"
	"ghenga/db"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "USAGE: %s db.sqlite3\n", os.Args[0])
		os.Exit(1)
	}

	dbfile := os.Args[1]
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

	fmt.Printf("create 500 people...")
	err = db.CreateFakePeople(dbmap, 500)
	if err != nil {
		panic(err)
	}
	fmt.Printf("done\n")
}
