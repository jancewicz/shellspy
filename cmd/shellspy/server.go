package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/jancewicz/shellspy"
)

var (
	listen = flag.Bool("l", false, "Listen")
	host   = flag.String("h", "localhost", "Host")
	port   = flag.Int("p", 0, "Port")
)

func startServer() {
	addr := fmt.Sprintf("%s:%d", *host, *port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	log.Printf("Server started on %s", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error occured while listening for connection: %v", err.Error())
		} else {
			go proccessClient(conn)
		}
	}
}

func proccessClient(conn net.Conn) {
	defer conn.Close()

	file, err := os.Create("shellspy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(conn)
	writer := io.MultiWriter(conn, file)

	for {
		command := shellspy.ReadUserInput(reader, conn)

		writer.Write([]byte("> " + command))
		writer.Write([]byte("\n"))

		if command == "exit" {
			fmt.Println("connection closed")
			conn.Close()
		}

		cmd, err := shellspy.CommandFromInput(command)
		if err != nil {
			log.Fatal(err)
		}

		cmd.Stdout = io.MultiWriter(os.Stdout, writer)
		cmd.Stderr = os.Stderr

		if err := shellspy.HandleCommand(cmd); err != nil {
			log.Fatal(err)
		}
	}
}

func startClient(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("can't connet to server on addres: %s", addr)
		return
	}

	_, err = io.Copy(conn, os.Stdout)
	if err != nil {
		fmt.Printf("connection error: %s", err)
	}
	fmt.Println("Recording session to 'shellspy.txt'")
}
