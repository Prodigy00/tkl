package subcommands

import "fmt"

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
    default actions:
       adds the date of creation automatically
	`
	listMsg = `
    subcommand name: list
    subcommand options:
         None yet
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
	for cmd, help := range *h {
		fmt.Printf("subcommand: %v", cmd)
		fmt.Println()
		fmt.Println("usage: ")
		fmt.Println(help)
	}
}
