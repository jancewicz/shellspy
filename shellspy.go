package shellspy

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
	"strings"
)

// Your implementation goes here!
func ReadUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	command, _ := reader.ReadString('\n')
	return strings.TrimSpace(command)
}

func CommandFromInput(input string) (*exec.Cmd, error) {
	cmd := exec.Command(input)
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}

	return cmd, nil
}
