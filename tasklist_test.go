package tasklist_test

import (
	"bytes"
	tasklist "tasklist-cli"
	"testing"
)

/**
*Arguments are all strings that follow a CLI command.
*Options are arguments with dashes (single or double) that are followed by user input and
*modify the operation of the command.
*Flags are boolean options that do not take user input.
 */

func TestAddTask(t *testing.T) {
	t.Parallel()

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
