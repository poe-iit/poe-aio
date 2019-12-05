// This is client 1 which is the headless button client. 
// Button can be pressed depending on the emergency
// Forwards emergency event to the server
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	
	"github.com/warthog618/gpio"
)



func main() {

	serverAddress := "192.168.2.50:65432"
	protocol := "tcp"

	// create a socket for connecting to the server
	sock, err := net.Dial(protocol, serverAddress)

	if err != nil {
		log.Output(1, err.Error())
	}

	log.Output(1, "Opening GPIO connection")

	err = gpio.Open()
	if err != nil {
		log.Fatal(err.Error())
	}
	
	defer gpio.Close()
	log.Output(1, "GPIO connection Opened")
	log.Output(1, "Waiting for Button Press")

	// Map buttons to pins
	firePin := gpio.NewPin(21)
	shooterPin := gpio.NewPin(20)

	for {
		// read emergency from GPIO buttons
		emergencyType, err := listenForButtonPress(firePin, shooterPin)

		if err != nil {
			log.Output(1, err.Error())
		}

		// The above code will normally block until a button is pressed
		//emergencyType := "client1 fire"

		// write emergency to server
		fmt.Fprintf(sock, emergencyType+"\n")
		fmt.Println("Sent message")

	}
}

func listenForButtonPress(firePin *gpio.Pin, shooterPin *gpio.Pin) (event string, err error)  {
	
	for {
		
		if firePin.Read() {
			return "fire", err
		}

		if shooterPin.Read() {
			return "shooter", err
		}
		
	}



}
