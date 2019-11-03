package main

import (
	"log"
	"net"
)

func main() {
	address := "127.0.0.1:65432"
	protocol := "tcp4"

	listen, err := net.Listen(protocol, address)

	if err != nil {
		log.Fatal(1, err.Error())
	}

	defer listen.Close()

	log.Output(1, "Server listening on "+address)

	// keeps running and checks for a connection
	// once a connection is established, handle it in another thread
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Output(1, err.Error())
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	log.Output(1, "Handling Connection")
}
