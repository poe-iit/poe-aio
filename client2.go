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



var fireOutPin *gpio.Pin
var shooterOutPin *gpio.Pin
var envOutPin *gpio.Pin

func main() {
	
	err := initPins()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer gpio.Close()

	//go listenForSmoke()

	serverAddress := "192.168.2.50:65432"
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
			writeToGPIO("Fire")
		} else if message == "shooter" {
			log.Output(1, "SHOOTER FROM HEADLESS CLIENT")
			writeToGPIO("Shooter")
		} else if message == "enviormental" {
			log.Output(1, "ENV FROM GUI")
			writeToGPIO("Enviormental")
		} else if message == "safety" {
			log.Output(1, "default")
			writeToGPIO("Safety")
		}
		

	}
}

func initPins() (err error) {
	err = gpio.Open()
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	log.Output(1, "GPIO connection Opened")

	// inits the pins, sets the pins to either input or output

	fireOutPin = gpio.NewPin(22)
	fireOutPin.SetMode(gpio.Output)
	fireOutPin.Write(gpio.High)
	shooterOutPin = gpio.NewPin(23)
	shooterOutPin.SetMode(gpio.Output)
	shooterOutPin.Write(gpio.High)
	envOutPin = gpio.NewPin(24)
	envOutPin.SetMode(gpio.Output)
	
	log.Output(1, "Pins initialized")
	
	return err

}



func writeToGPIO(emergencyType string) {
	log.Output(1, "Writing to GPIO")
	switch emergencyType {
	case "Fire":
		triggerButton(fireOutPin)
	case "Shooter":
		triggerButton(shooterOutPin)

	case "Enviormental":
		triggerButton(envOutPin)	
	}
}


func listenForSmoke() {
	log.Output(1, "Listening for smoke")
	smokePin := gpio.NewPin(13)

	for {
		if !smokePin.Read() == true {
			log.Output(1, "SMOKE DETECTED")
			writeToGPIO("Fire")
			time.Sleep(5 * time.Second)
		}
	}

}


func triggerButton(pin *gpio.Pin) {

	pin.Write(gpio.Low)
	time.Sleep(250 * time.Millisecond)
	pin.Write(gpio.High)

}
