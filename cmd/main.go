package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"tasklist-cli/subcommands"
)

var (
	appVersion  = "0.1.0"
	_version    = flag.Bool("v", false, "installed version of tasklist cli")
	_versionSub = flag.Bool("version", false, "installed version of tasklist cli")
)

func main() {
	flag.Parse()

	if *_version || *_versionSub {
		fmt.Println(appVersion)
		os.Exit(0)
	}

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
	case "help":
		subcommands.GetHelp().PrintHelpList()
	default:
		log.Fatalf("Unrecognized command %q. Command must be one of add, list, update, delete", cmd)
	}
}
