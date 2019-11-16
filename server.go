package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	"fmt"
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
		
		// get ip and port of connection, determines if its client 1 or client 2
		connAddress := strings.Split(conn.RemoteAddr().String(), ":")
		connIP, connPort := connAddress[0], connAddress[1]
		fmt.Println("IP: " +connIP)
		fmt.Println("Port: " +connPort)

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
			conn.Close()
			return
		}

		data := strings.TrimSpace(string(rawdata)) //clean up the data
		fmt.Println(data)

		if strings.Contains(string(data), "client1") {
			fmt.Println("client 1 detected")
			conn.Write([]byte("hello client 1"+"\n"))

		} else {
			conn.Write([]byte("are you sure you are client1"+"\n"))
		}

	}
}
