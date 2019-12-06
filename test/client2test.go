package main

import (

	"fmt"
	"log"
	"net/http"
	"time"
	
	"github.com/warthog618/gpio"
	"io"
	"os"
	"github.com/hajimehoshi/oto"

	"github.com/hajimehoshi/go-mp3"
	"github.com/gorilla/mux"

)



var fireOutPin *gpio.Pin
var shooterOutPin *gpio.Pin
var envOutPin *gpio.Pin
var smokePin *gpio.Pin



func main() {
	
	err := initPins()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer gpio.Close()

	go listenForSmoke()
	log.Output(1, "Listening for smoke")




	router := mux.NewRouter()
	router.HandleFunc("/lights", handleRequests)
	log.Fatal(http.ListenAndServe(":12345", router))
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
	smokePin = gpio.NewPin(13)
	//smokePin.SetMode(gpio.Input)
	//smokePin.Write(gpio.High)
	
	log.Output(1, "Pins initialized")
	
	return err

}



func writeToGPIO(emergencyType string) {
	log.Output(1, "Writing to GPIO")
	switch emergencyType {
	case "Fire":
		triggerButton(fireOutPin)
		audio("../audio/fire.mp3")
	case "Shooter":
		triggerButton(shooterOutPin)
		audio("../audio/shooter.mp3")

	case "Enviormental":
		triggerButton(envOutPin)
		audio("../audio/env.mp3")	
	}
}


func listenForSmoke() {

	for{

		if !smokePin.Read() == true {
			log.Output(1, "SMOKE DETECTED")
			writeToGPIO("Fire")
			time.Sleep(5 * time.Second)
		}

		time.Sleep(1 *time.Millisecond)
	}

	

}


func triggerButton(pin *gpio.Pin) {

	pin.Write(gpio.Low)
	time.Sleep(250 * time.Millisecond)
	pin.Write(gpio.High)
	

}




func handleRequests(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "handling request from webpage")

 
    switch r.Method {
	case "POST":
		log.Output(1, "Post Request Recieved")

        // Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		err := r.ParseForm()
		
		if err != nil {
			log.Output(1, err.Error())
		}

	
		
		emergencyType := r.Form["emergency"][0]

		switch emergencyType {
		case "fire":
			writeToGPIO("Fire")
		case "Shooter":
			writeToGPIO("Shooter")
		case "enviormental":
			writeToGPIO("Enviormental")
		case "safety":
			writeToGPIO("Safety")

		}
		fmt.Println(emergencyType)

        
    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}



func audio(pathToFile string) error {
	log.Output(1, "Playing Audio")
	f, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	p, err := oto.NewPlayer(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}

	fmt.Printf("Length: %d[bytes]\n", d.Length())

	if _, err := io.Copy(p, d); err != nil {
		return err
	}

	p.Close()
	return nil
}
