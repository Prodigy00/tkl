package subcommands

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/enescakir/emoji"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Task string

type TaskFile struct {
	file  os.File
	tasks []Task
	io.Reader
	io.Writer
}

func NewTaskFile() *TaskFile {
	f, err := os.Create("tklfile.txt")
	if err != nil {
		log.Fatal("Could not create taskfile")
	}

	defer closeFile(f)

	return &TaskFile{
		file:  *f,
		tasks: []Task{},
	}
}

func closeFile(f *os.File) {
	closeErr := f.Close()
	if closeErr != nil {
		log.Fatalf("error occurred closing the file. Err: %v", closeErr)
	}
}

func readFile(f *os.File) (ts []string, err error) {
	var taskslice []string
	reader := bufio.NewScanner(f)

	for reader.Scan() {
		taskslice = append(taskslice, reader.Text())
	}

	if scanErr := reader.Err(); scanErr != nil {
		return nil, scanErr
	}
	return taskslice, nil
}

type Level int

const (
	NormalTaskLevel Level = iota
	ImportantTaskLevel
	CriticalTaskLeveL
)

func (t *TaskFile) AddTask(tk Task) {
	t.tasks = append(t.tasks, tk)

	var filename = t.file.Name()

	if len(filename) < 3 {
		filename = "taskfile.txt"
		fmt.Println("got here")
	}

	//append content to file
	tklFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 666)
	if err != nil {
		fmt.Errorf("error opening file")
	}

	defer closeFile(tklFile)

	_, writeErr := tklFile.WriteString(string(tk))

	if writeErr != nil {
		log.Fatal(writeErr)
	}
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

	tklF := NewTaskFile()
	tklF.tasks = append(tklF.tasks, taskentry)

	tklF.AddTask(taskentry)
}
