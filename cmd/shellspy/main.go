package main

import (
	"flag"
	"fmt"
)

func main() {
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

	startClient(fmt.Sprintf("%s:%s", serverHost, serverPort))
}
