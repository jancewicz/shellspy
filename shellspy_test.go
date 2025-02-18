package shellspy_test

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/jancewicz/shellspy"
)

// Your tests go here!
func TestCommandInput(t *testing.T) {
	input := strings.NewReader("pwd")
	expected := "/home/jancewicz/Repos/shellspy"

	scanner := bufio.NewScanner(input)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
	command := scanner.Text()

	cmd, err := shellspy.CommandFromInput(command)
	if err != nil {
		t.Fatal(err)
	}

	r, w, _ := os.Pipe()

	cmd.Stdout = w
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)
	result := strings.TrimSpace(buf.String())

	if result != expected {
		t.Errorf("wrong output, got: %s, want: %s", strings.TrimSpace(buf.String()), expected)
	}
}

func MockReadingInput(commands []string) func() string {
	idx := 0

	return func() string {
		if idx < len(commands) {
			command := commands[idx]
			idx++

			return command
		}
		// fallback
		return "exit"
	}
}
