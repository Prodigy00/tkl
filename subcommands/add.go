package subcommands

import (
	"flag"
	"fmt"
	"github.com/enescakir/emoji"
	"log"
	"os"
	"strings"
	"tasklist-cli/util"
	"time"
)

type Task string

type TaskFile struct {
	file  os.File
	tasks []Task
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

type Level int

const (
	NormalTaskLevel Level = iota
	ImportantTaskLevel
	CriticalTaskLeveL
)

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
	fmt.Println("Task added successfully!")
}

func Add(args []string) {
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

	//fmt.Println("id: ", id, " :", priority, " task: ", *taskname, " date: ", taskdate)
	entry := strings.Builder{}
	entry.WriteString(fmt.Sprintf("%v: ", id))
	entry.WriteString(fmt.Sprintf("%v ", priority))
	entry.WriteString(fmt.Sprintf("%v ", *taskname))
	entry.WriteString(fmt.Sprintf("%v", taskdate))

	taskentry := Task(entry.String())

	tkF := NewTaskFile()
	tkF.AddTask(taskentry)
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
