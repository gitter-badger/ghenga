package main

import "github.com/jessevdk/go-flags"

type globalOptions struct{}

var globalOpts = globalOptions{}
var parser = flags.NewParser(&globalOpts, flags.HelpFlag|flags.PassDoubleDash)
