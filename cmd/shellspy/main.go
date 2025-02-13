package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Your CLI goes here!
func main() {
	file, err := os.Create("shellspy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("Recording session to 'shellspy.txt'")

	flag.Parse()
	if *listen {
		startServer()
		return
	}

	if len(flag.Args()) < 2 {
		fmt.Println("host and port number required")
		return
	}

	serverHost := flag.Arg(0)
	serverPort := flag.Arg(1)

	startClient(fmt.Sprintf("%s:%s", serverHost, serverPort), file)
}
