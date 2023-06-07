package main

import (
	"flag"
	"log"
	"tasklist-cli/subcommands"
)

var (
	_version = flag.String("v", "0.1.0", "installed version of tasklist cli")
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("kindly specify a subcommand")
	}

	cmd, args := args[0], args[1:]

	switch cmd {
	case "add":
		subcommands.Add(args)
	case "list":
		subcommands.List(args)
	case "update":
		subcommands.Update(args)
	case "delete":
		subcommands.DeleteTask(args)
	default:
		log.Fatalf("Unrecognized command %q. Command must be one of add, list, update, delete", cmd)
	}
}
