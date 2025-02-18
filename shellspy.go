package shellspy

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func ReadUserInput(reader *bufio.Reader, conn net.Conn) string {
	command, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Sprintf("error: %s", err)
	}
	return strings.TrimSpace(command)
}

func CommandFromInput(input string) (*exec.Cmd, error) {
	commands := strings.Fields(input)
	initCmd := commands[0]
	optCmds := commands[1:]

	cmd := exec.Command(initCmd, optCmds...)
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}
	return cmd, nil
}

func HandleCommand(cmd *exec.Cmd) error {
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
