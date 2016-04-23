package main

import (
	"fmt"
	"ghenga/db"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/jmoiron/modl"
)

type globalOptions struct {
	Environment string `short:"e" long:"environment" default:"production" description:"Environment to use"`
}

func (opts *globalOptions) DatabaseFilename() string {
	switch opts.Environment {
	case "test":
		return "db/test.db"
	case "production":
		return "db/production.db"
	}

	fmt.Fprintf(os.Stderr, "invalid environment %q, using production", opts.Environment)
	opts.Environment = "production"
	return opts.DatabaseFilename()
}

// OpenDB opens the database, which is created if necessary. Before exit,
// cleanup() should be called to properly close the database connection.
func OpenDB() (dbm *modl.DbMap, cleanup func() error, err error) {
	dbm, err = db.Init(globalOpts.DatabaseFilename())
	if err != nil {
		return nil, nil, err
	}

	cleanup = func() error { return dbm.Db.Close() }
	return dbm, cleanup, nil
}

var globalOpts = globalOptions{}
var parser = flags.NewParser(&globalOpts, flags.HelpFlag|flags.PassDoubleDash)

