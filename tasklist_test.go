package tasklist_test

import (
	"bytes"
	"log"
	tasklist "tasklist-cli"
	"testing"
)

/**
*Arguments are all strings that follow a CLI command.
*Options are arguments with dashes (single or double) that are followed by user input and
*modify the operation of the command.
*Flags are boolean options that do not take user input.
 */

// save tasks to a file
func TestAddTask(t *testing.T) {
	t.Parallel()
	stubFile := &bytes.Buffer{}
	cmd := "add"
	task := "new test task"
	err := tasklist.AddTask([]string{cmd, task})
	if err != nil {
		log.Fatal(err)
	}
}

func TestPrintTklWelcomeMsgToWriter(t *testing.T) {
	t.Parallel()
	stubTerminal := &bytes.Buffer{}
	tasklist.Welcome(stubTerminal)
	want := "Welcome to tasklist!"
	got := stubTerminal.String()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
