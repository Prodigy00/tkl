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
	//fmt.Printf(" cmd: %s || args: %s ", cmd, args)

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

//replace task file struct with a file that will be created if it doesn't already exist

//func root(args []string) error {
//	if len(args) < 1 {
//		return errors.New("You must pass a sub-command")
//	}
//
//	cmds := []tasklist.Runner{
//		tasklist.NewAddTaskCommand(),
//	}
//
//	subcommand := os.Args[1]
//
//	for _, cmd := range cmds {
//		if cmd.Name() == subcommand {
//			err := cmd.Init(os.Args[2:])
//			if err != nil {
//				log.Fatal(err)
//			}
//			return cmd.Run()
//		}
//	}
//
//	return fmt.Errorf("Unknown subcommand: %s", subcommand)
//}
//
//func main() {
//	if err := root(os.Args[1:]); err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//}
