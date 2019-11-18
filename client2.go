package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
	
	"github.com/warthog618/gpio"
)


var smokePin *gpio.Pin
var fireOutPin *gpio.Pin
var shooterOutPin *gpio.Pin
var envOutPin *gpio.Pin

func main() {
	
	err := initPins()
	if err != nil {
		log.Fatal(err.Error())
	}

	go listenForSmoke()

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

func initPins() (err error) {
	err = gpio.Open()
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	defer gpio.Close()
	log.Output(1, "GPIO connection Opened")

	// inits the pins, sets the pins to either input or output
	smokePin = gpio.NewPin(13)
	fireOutPin = gpio.NewPin(14)
	fireOutPin.SetMode(gpio.Output)
	shooterOutPin = gpio.NewPin(16)
	shooterOutPin.SetMode(gpio.Output)
	envOutPin = gpio.NewPin(17)
	envOutPin.SetMode(gpio.Output)
	
	log.Output(1, "Pins initialized")
	
	return err

}



func writeToGPIO(emergencyType string) {
	log.Output(1, "Writing to GPIO")
	switch emergencyType {
	case "Fire":
		fireOutPin.Write(gpio.High)
	case "Shooter":
		shooterOutPin.Write(gpio.High)

	case "Enviormental":
		envOutPin.Write(gpio.High)	
	}
}


func listenForSmoke() {
	log.Output(1, "Listening for smoke")

	for {
		if smokePin.Read() == true {
			log.Output(1, "SMOKE DETECTED")
			writeToGPIO("Fire")
			time.Sleep(5 * time.Second)
		}
	}

}