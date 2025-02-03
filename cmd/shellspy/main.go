package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/jancewicz/shellspy"
)

// Your CLI goes here!
func main() {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	cmd, err := shellspy.CommandFromInput(text)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cmd)
}
