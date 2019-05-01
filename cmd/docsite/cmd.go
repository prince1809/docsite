package main

import (
	"flag"
	"log"
	"text/template"
)

// command is a subcommand handler and its set.
type command struct {
}

// commander represents a top-level command with subcommands.
type commander []*command

// run runs the command
func (c commander) run(flagSet *flag.FlagSet, cmdName string, usage *template.Template, args []string) {
	// Parse flags
	flagSet.Usage = func() {
		data := struct {
			FlagUsage func() string
			Commands  []*command
		}{
			FlagUsage: func() string { commandLine.PrintDefaults(); return "" },
			Commands:  c,
		}
		if err := usage.Execute(commandLine.Output(), data); err != nil {
			log.Fatal(err)
		}
	}
}
