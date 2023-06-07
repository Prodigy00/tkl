package subcommands

import (
	"flag"
	"fmt"
	"github.com/enescakir/emoji"
	"log"
	"os"
	"strconv"
	"strings"
	"tasklist-cli/util"
	"time"
)

type Task string

type TaskFile struct {
	file  os.File
	tasks []Task
}

func Add(args []string) {
	flag := flag.NewFlagSet("add", flag.ExitOnError)

	var (
		taskname = flag.String("t", "", "name of the task to be added. (Required)")
		level    = flag.Int("l", 0, "urgency level of task. 0 for the least urgent tasks, 1 for mildly urgent and 2 for top priority tasks")
		help     = flag.Bool("h", false, "help flag for add subcomand")
		helpSub  = flag.Bool("help", false, "help flag for add subcomand")
	)

	err := flag.Parse(args)
	if err != nil {
		log.Fatal(err)
	}

	if *help || *helpSub {
		GetHelp().PrintSubcommandHelp("add")
		os.Exit(0)
	}

	var id string
	var priority emoji.Emoji
	var taskdate string
	var taskidnum int

	//check if task exists first
	tkF := NewTaskFile()

	if len(tkF.tasks) > 0 {
		taskidnum = len(tkF.tasks)
	}

	if *taskname != "" {
		taskdate = time.Now().Local().String()
		id = "task–" + strconv.Itoa(taskidnum+1)
	}

	switch *level {
	case 2:
		priority = emoji.RedSquare
	case 1:
		priority = emoji.YellowSquare
	default:
		priority = emoji.GreenSquare
	}

	entry := strings.Builder{}
	entry.WriteString(fmt.Sprintf("%v: ", id))
	entry.WriteString(fmt.Sprintf("%v ", priority))
	entry.WriteString(fmt.Sprintf("%v ", *taskname))
	entry.WriteString(fmt.Sprintf("%v", taskdate))

	taskentry := Task(entry.String())

	tkF.AddTask(taskentry)
}

func NewTaskFile() *TaskFile {
	//check if file already exists first before creating a new file
	var filename = "tklfile.txt"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		f, err := os.Create(filename)
		if err != nil {
			log.Fatal("Could not create taskfile")
		}

		defer util.CloseFile(f)

		return &TaskFile{
			file:  *f,
			tasks: []Task{},
		}
	}

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("could not open taskfile")
	}
	defer util.CloseFile(f)

	tasks, err := util.ReadTasksFromFile(f)
	if err != nil {
		fmt.Println("error getting tasks", err)
	}

	var ft []Task
	for _, t := range tasks {
		ft = append(ft, Task(t))
	}

	return &TaskFile{
		file:  *f,
		tasks: ft,
	}
}

func (t *TaskFile) AddTask(tk Task) {
	var filename = t.file.Name()

	if len(filename) < 3 {
		filename = "taskfile.txt"
		fmt.Println("got here")
	}

	err := AppendTaskToFile(filename, tk)
	if err != nil {
		fmt.Println("err appending to file", err)
	}
	log.Printf("%s Task added successfully!", emoji.WritingHand)
}

func AppendTaskToFile(filename string, task Task) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer util.CloseFile(file)

	_, err = file.WriteString(fmt.Sprintf("%s\n", task))
	if err != nil {
		return fmt.Errorf("failed to append to file: %v", err)
	}

	return nil
}
