package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/jancewicz/shellspy"
)

// Your CLI goes here!
func main() {
	reader := bufio.NewReader(os.Stdin)
	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command)

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
