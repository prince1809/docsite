package main

import (
	"flag"
	"text/template"
)

// command is a subcommand handler and its set.
type command struct {
}

// commander represents a top-level command with subcommands.
type commander []*command

// run runs the command
func (c commander) run(flagSet *flag.FlagSet, cmdName string, usage *template.Template, args []string) {

}
