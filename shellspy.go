package shellspy

import (
	"errors"
	"os/exec"
)

// Your implementation goes here!
func CommandFromInput(input string) (*exec.Cmd, error) {
	cmd := exec.Command(input)
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}

	return cmd, nil
}
