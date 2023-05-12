package tasklist

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func NewAddTaskCommand() *AddCommand {
	atcmd := &AddCommand{
		fs: flag.NewFlagSet("add", flag.ContinueOnError),
	}
	atcmd.fs.StringVar(&atcmd.taskname, "task", "new task", "name of new task to be added")
	return atcmd
}

type AddCommand struct {
	fs       *flag.FlagSet
	taskname string
}

func (ac *AddCommand) Name() string {
	return ac.fs.Name()
}

func (ac *AddCommand) Init(args []string) error {
	return ac.fs.Parse(args)
}

func (ac *AddCommand) Run() error {
	fmt.Println("new task added: ", ac.taskname)
	return nil
}

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

// should exist as a subcommand e.g tkl add
// should be able to add task with -t argument as title
// i.e tkl add -m "taskdescription" -t "tag_which could be project name"
// could also be tkl add -tm "taskdescription":"tag"
func AddTask(args []string) error {
	if len(args) < 1 {
		return errors.New("You must pass a sub-command")
	}

	cmds := []Runner{
		NewAddTaskCommand(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			err := cmd.Init(os.Args[2:])
			if err != nil {
				log.Fatal(err)
			}
			return cmd.Run()
		}
	}

	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}

func Welcome(w io.Writer) {
	welcomeMsg := "Welcome to tasklist!"
	_, err := w.Write([]byte(welcomeMsg))
	if err != nil {
		log.Fatal(err)
	}
}
