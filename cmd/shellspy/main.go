package main

import (
	"fmt"
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
	readInput := shellspy.ReadUserInput

	if err := shellspy.RunShell(readInput, file); err != nil {
		log.Fatal(err)
	}
}
