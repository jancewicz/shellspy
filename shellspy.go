package shellspy

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
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

func RunShell(file *os.File) {
	for {
		fmt.Print("> ")
		command := ReadUserInput()
		io.WriteString(file, ("> " + command + "\n"))

		if command == "exit" {
			break
		}

		cmd, err := CommandFromInput(command)
		if err != nil {
			log.Fatal(err)
		}

		cmd.Stdout = io.MultiWriter(file, os.Stdout)
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		io.WriteString(file, "\n")
	}
}
