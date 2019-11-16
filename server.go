package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	//"fmt"
)

type Client struct { 
	name string
	connection net.Conn
	ip string
	port string

}

func main() {
	client := Client{}
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
			conn.Close()

		} else {
			
			// get ip and port of connection, determines if its client 1 or client 2
			connAddress := strings.Split(conn.RemoteAddr().String(), ":")
			connIP, connPort := connAddress[0], connAddress[1]
			
			if connIP == "192.168.2.51" {
				client = Client{
					name: "client1",
					connection: conn,
					ip: connIP, 
					port: connPort}
			} else if connIP == "192.169.2.52" {
				client = Client{
					name: "client2",
					connection: conn,
					ip: connIP, 
					port: connPort}
			} else {
				client = Client{
					name: "client1",
					connection: conn,
					ip: connIP, 
					port: connPort}
			}
			
			go handleConnection(client)	
		}
	}
}


func handleConnection(client Client) {
	// This function takes the connection and reads the data within in to determine what to do
	log.Output(1, "Handling Connection for " +client.name)
	var message string

	for {

		if client.name == "client1" {
			message = getDataFromClient(client.connection)
			//fmt.Println(message)
			switch {
			case strings.Contains(message, "fire"):
				log.Output(1, "Fire Detected from client 1, forwarding to ceiling client2")

			case strings.Contains(message, "Shooter"):
				log.Output(1, "Shooter Detected from client 1, forwarding to ceiling client2")
			case message == "connectionBroke":
				// handles cases where the clients close unexpectadly 
				break
			
			} 
			break
		}

	}
}



// This function takes in the connection and reads the raw data out of it
// Only data with a newline appended to the end will be read
func getDataFromClient(connection net.Conn) (data string) {
	rawdata, err := bufio.NewReader(connection).ReadString('\n')

	if err != nil {
		log.Output(1, err.Error())
		connection.Close()
		return "connectionBroke"
	}

	data = strings.TrimSpace(string(rawdata)) //clean up the data
	
	return strings.ToLower(data)

}
