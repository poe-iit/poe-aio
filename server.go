package main

import (
	"bufio"
	"log"
	"net"
	"strings"
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
	// This function takes the connection and reads the data within in to determine what to do
	log.Output(1, "Handling Connection")

	for {
		rawdata, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			log.Output(1, err.Error())
		}

		data := strings.TrimSpace(string(rawdata)) //clean up the data

		if data == "test" {
			log.Output(1, "successfully read data, found test")
			conn.Write([]byte("server read your message correctly"))
		} else {
			log.Output(1, "client message did not contain test")
			conn.Write([]byte("server did not find test in your message"))
		}

	}
}
