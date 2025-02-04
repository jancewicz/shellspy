package main

import (
	"log"
	"os"

	"github.com/jancewicz/shellspy"
)

// Your CLI goes here!
func main() {

	for {
		command := shellspy.ReadUserInput()
		if command == "exit" {
			break
		}

		cmd, err := shellspy.CommandFromInput(command)
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create("logger.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
