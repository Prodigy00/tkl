package subcommands

import (
	"fmt"
	"sort"
)

type Help map[string]string

const (
	AddCmd    = "add"
	ListCmd   = "list"
	UpdateCmd = "update"
	DeleteCmd = "delete"
)

const (
	addMsg = `
    subcommand name: add
    subcommand options:
         -t : task name(string),
         -l : level of importance (int)(0,1,2)
	     -h : description for add subcommand
	     -help: description for add subcommand
    default actions:
       adds the date of creation automatically
	`
	listMsg = `
    subcommand name: list
    subcommand options:
         -n : limit number of tasks you want to list, defaults to 0 which means no limits
         -h : description for list subcommand
         -help : description for list subcommand
    default actions:
       lists tasks
	`
	updateMsg = `
    subcommand name: update
    subcommand options:
         None yet
    default actions:
       update a task
	`
	deleteMsg = `
    subcommand name: delete
    subcommand options:
         None yet
    default actions:
       delete task(s)
	`
)

func GetHelp() *Help {
	dM := make(Help)
	dM[AddCmd] = addMsg
	dM[ListCmd] = listMsg
	dM[UpdateCmd] = updateMsg
	dM[DeleteCmd] = deleteMsg

	return &Help{
		AddCmd:    dM[AddCmd],
		ListCmd:   dM[ListCmd],
		UpdateCmd: dM[UpdateCmd],
		DeleteCmd: dM[DeleteCmd],
	}
}

func (h *Help) PrintSubcommandHelp(name string) {
	var cmd = *h
	fmt.Println(cmd[name])
}

func (h *Help) PrintHelpList() {
	cmdMap := *h

	keyNames := make([]string, 0, 4)

	for k, _ := range cmdMap {
		keyNames = append(keyNames, k)
	}
	sort.Strings(keyNames)

	for _, cmd := range keyNames {
		fmt.Printf("subcommand: %v", cmd)
		fmt.Println()
		fmt.Print("usage: ")
		fmt.Println(cmdMap[cmd])
	}
}
