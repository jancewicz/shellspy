package shellspy_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/jancewicz/shellspy"
)

// Your tests go here!
func TestCommandInput(t *testing.T) {
	reader := bufio.NewReader(os.Stdin)

	command, _ := reader.ReadString('\n')

	cmd, err := shellspy.CommandFromInput(command)
	if err != nil {
		t.Fatal(err)
	}

	if cmd.Args[0] != command {
		t.Errorf("provided args are not the same as user input, got: %s, want: %s", cmd.Args[0], command)
	}
}
