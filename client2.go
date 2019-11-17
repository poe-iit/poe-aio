package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	
	//"github.com/warthog618/gpio"
)



func main() {

	serverAddress := "127.0.0.1:65432"
	protocol := "tcp"

	// create a socket for connecting to the server
	sock, err := net.Dial(protocol, serverAddress)

	if err != nil {
		log.Output(1, err.Error())
	}

	// start a connection with the server so it knows we exist
	message := "client2"
	fmt.Fprintf(sock, message+"\n")

	for {
		// listen for reply from the server
		rawmessage, _ := bufio.NewReader(sock).ReadString('\n')
		message = strings.TrimSpace(string(rawmessage)) //clean up the data

		if message == "serverhandshake" {
			log.Output(1, "Message from server: "+message)
		} else if message == "fire" {
			log.Output(1,"FIRE FROM HEADLESS CLIENT")
		} else if message == "shooter" {
			log.Output(1, "SHOOTER FROM HEADLESS CLIENT")
		}
		

	}
}