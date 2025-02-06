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

func TestRunShell(t *testing.T) {
	file, err := os.CreateTemp("", "test.txt")
	if err != nil {
		t.Fatalf("error occured during file creation: %v", err)
	}

	defer os.Remove(file.Name())
	defer file.Close()

	commands := []string{"ls", "pwd", "echo hello"}
	readInput := MockReadingInput(commands)

	err = shellspy.RunShell(file, readInput)
	if err != nil {
		t.Fatalf("error occured during running the shell: %v", err)
	}

	content, err := os.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("error occured during reading the file: %v", err)
	}

	for _, cmd := range commands {
		if !strings.Contains(string(content), "> "+cmd) {
			t.Errorf("expected command %s in output", cmd)
		}
	}
}
