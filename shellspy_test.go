package shellspy_test

import (
	"testing"

	"github.com/jancewicz/shellspy"
)

// Your tests go here!
func TestCommandInput(t *testing.T) {
	input := "ls"

	cmd, err := shellspy.CommandFromInput(input)
	if err != nil {
		t.Fatal(err)
	}

	if cmd.Args[0] != input {
		t.Errorf("provided args are not the same as user input, got: %s, want: %s", cmd.Args[0], input)
	}
}
