// This is client 1 which is the headless button client. 
// Button can be pressed depending on the emergency
// Forwards emergency event to the server
package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"net/url"	
	
//	"github.com/warthog618/gpio"
)



func main() {


/*	log.Output(1, "Opening GPIO connection")

	err := gpio.Open()
	if err != nil {
		log.Output(1, err.Error())
	}	
	defer gpio.Close()
	log.Output(1, "GPIO connection Opened")

	// Map buttons to pins
	firePin := gpio.NewPin(21)
	shooterPin := gpio.NewPin(20)
*/
	for {
		// read emergency from GPIO buttons
	  //emergencyType := listenForButtonPress(firePin, shooterPin)
		emergencyType := "fire"
		fmt.Println(emergencyType)

		APIURL := "http://127.0.0.1:12345/button"

		response, err := http.PostForm(APIURL,
		url.Values{"emergency": {emergencyType}})
	
		if err != nil {
			log.Output(1, err.Error())
		}

		
		fmt.Println(response)
		

		fmt.Println("Sent message")
		time.Sleep(1* time.Second)

	}
}


/*
func listenForButtonPress(firePin *gpio.Pin, shooterPin *gpio.Pin) (event string, err error) {
	fmt.Println("Listening for button Press")
	
	for {
		
		if firePin.Read() {
			return "fire", err
		}

		if shooterPin.Read() {
			return "Shooter", err
		}
		
	}



}*/
