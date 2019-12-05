// This is client 1 which is the headless button client. 
// Button can be pressed depending on the emergency
// Forwards emergency event to the server
package main

import (
	"fmt"
	"log"
	"net"
	"time"
	
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

	// Map buttons to pins
	firePin := gpio.NewPin(21)
	shooterPin := gpio.NewPin(20)

	for {
		// read emergency from GPIO buttons
		emergencyType, err := listenForButtonPress(firePin, shooterPin)

		if err != nil {
			log.Output(1, err.Error())
		}

		// write emergency to server
		fmt.Fprintf(sock, emergencyType+"\n")
		fmt.Println("Sent message")
		time.Sleep(5* time.Second)

	}
}

func listenForButtonPress(firePin *gpio.Pin, shooterPin *gpio.Pin) (event string, err error) {
	fmt.Println("Listening for button Press")
	
	for {
		
		if firePin.Read() {
			return "fire", err
		}

		if shooterPin.Read() {
			return "shooter", err
		}
		
	}



}
