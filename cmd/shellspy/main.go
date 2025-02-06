package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jancewicz/shellspy"
)

// Your CLI goes here!
func main() {
	file, err := os.Create("shellspy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("Recording session to 'shellspy.txt'")

	for {
		fmt.Print("> ")
		command := shellspy.ReadUserInput()
		io.WriteString(file, ("> " + command + "\n"))

		if command == "exit" {
			break
		}

		cmd, err := shellspy.CommandFromInput(command)
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
