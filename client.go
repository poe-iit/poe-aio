package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	serverAddress := "127.0.0.1:65432"
	protocol := "tcp"

	// create a socket for connecting to the server
	sock, err := net.Dial(protocol, serverAddress)

	if err != nil {
		log.Output(1, err.Error())
	}

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		// write text to socket, sending it to the server
		fmt.Fprintf(sock, text+"\n")

		// listen for reply
		message, _ := bufio.NewReader(sock).ReadString('\n')
		log.Output(1, "Message from server: "+message)

	}
}
