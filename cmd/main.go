package main

import (
	"flag"
	"fmt"
	"github.com/enescakir/emoji"
	"log"
	"strings"
	"time"
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
		add(args)
	case "list":
		list(args)
	case "update":
		update(args)
	case "delete":
		deleteTask(args)
	default:
		log.Fatalf("Unrecognized command %q. Command must be one of add, list, update, delete", cmd)
	}
}

type Task map[string]string

type TaskFile struct {
	Tasks []Task
}

func (t *TaskFile) GetTasks() []Task {
	return t.Tasks
}

func (t *TaskFile) Add(tk Task) {
	t.Tasks = append(t.Tasks, tk)
}

func add(args []string) {
	flag := flag.NewFlagSet("add", flag.ExitOnError)

	var (
		taskname = flag.String("t", "", "name of the task to be added. (Required)")
		level    = flag.Int("l", 0, "urgency level of task. 0 for the least urgent tasks, 1 for mildly urgent and 2 for top priority tasks")
	)

	err := flag.Parse(args)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf(" taskname: %s, level: %d ", *taskname, *level)
	var id string
	var priority emoji.Emoji
	var taskdate string

	if *taskname != "" {
		ta := time.Now().Local().String()
		taskdate = ta
		splitTa := strings.Split(ta, " ")[:2]

		joinedTa := strings.Join(splitTa, "")
		id = "task–" + joinedTa
	}

	switch *level {
	case 2:
		priority = emoji.RedSquare
	case 1:
		priority = emoji.YellowSquare
	default:
		priority = emoji.GreenSquare
	}

	newMap := make(Task)

	newMap[id] = " " + priority.String() + " " + *taskname + " " + taskdate

	var tkf TaskFile

	tkf.Add(newMap)

	for k, v := range tkf.GetTasks() {
		fmt.Println(k, " : ", v)
	}
}
func list(args []string)       {}
func update(args []string)     {}
func deleteTask(args []string) {}

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
