package main

import (
	"log"
	"os"

	"github.com/jancewicz/shellspy"
)

// Your CLI goes here!
func main() {
	command := shellspy.ReadUserInput()

	cmd, err := shellspy.CommandFromInput(command)
	if err != nil {
		log.Fatal(err)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
