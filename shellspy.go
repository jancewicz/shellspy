package shellspy

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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
	commands := strings.Fields(input)
	initCmd := commands[0]
	optCmds := commands[1:]

	cmd := exec.Command(initCmd, optCmds...)
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}
	return cmd, nil
}

func HandleCommand(cmd *exec.Cmd, file *os.File) error {
	RedirectOutput(cmd, file)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func RedirectOutput(cmd *exec.Cmd, file *os.File) {
	cmd.Stdout = io.MultiWriter(file, os.Stdout)
	cmd.Stderr = os.Stderr
}

func RunShell(file *os.File) error {
	for {
		fmt.Print("> ")
		command := ReadUserInput()
		io.WriteString(file, ("> " + command + "\n"))

		if command == "exit" {
			break
		}

		cmd, err := CommandFromInput(command)
		if err != nil {
			return err
		}

		if err := HandleCommand(cmd, file); err != nil {
			return err
		}

		io.WriteString(file, "\n")
	}

	return nil
}
