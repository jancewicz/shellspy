package shellspy

import (
	"errors"
	"log"
	"os/exec"
)

// Your implementation goes here!
func CommandFromInput(input string) (*exec.Cmd, error) {
	cmd := exec.Command(input)
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	return cmd, nil
}
