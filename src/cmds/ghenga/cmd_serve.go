package main

import (
	"fmt"
	"ghenga/server"
	"log"
)

type cmdServe struct {
	Port   uint   `short:"p" long:"port"   default:"8080"   description:"set the port for the HTTP server"`
	Addr   string `short:"b" long:"bind"   default:""       description:"bind to this address"`
	Public string `          long:"public" default:"public" description:"the directory to server static files from"`
}

func init() {
	_, err := parser.AddCommand("serve",
		"start server",
		"The server command starts the HTTP server",
		&cmdServe{})
	if err != nil {
		panic(err)
	}
}

func (opts *cmdServe) Execute(args []string) (err error) {
	dbmap, cleanup, e := OpenDB()
	if e != nil {
		return e
	}
	defer CleanupErr(&err, cleanup)

	log.Printf("starting server at %v:%d", opts.Addr, opts.Port)

	env := &server.Env{
		ListenAddr: fmt.Sprintf("%s:%d", opts.Addr, opts.Port),
		DbMap:      dbmap,
	}

	err = server.ListenAndServe(env)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
