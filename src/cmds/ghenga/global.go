package main

import "github.com/jessevdk/go-flags"

type globalOptions struct{
	Environment string `short:"e" long:"environment" default:"production" description:"Environment to use"`
}

var globalOpts = globalOptions{}
var parser = flags.NewParser(&globalOpts, flags.HelpFlag|flags.PassDoubleDash)
