package main

import (

	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/warthog618/gpio"
	"os"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

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


	// start the http server to listen to requests from the server
	router := mux.NewRouter()
	router.HandleFunc("/lights", handleRequests)
	log.Println("About to listen on 12345. Go to https://localhost:12345")
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

	log.Output(1, "Pins initialized")

	return err

}



func writeToGPIO(emergencyType string) {
	log.Output(1, "Writing to GPIO")

	switch emergencyType {
	case "fire":
		triggerButton(fireOutPin)
		audio("./audio/fire.mp3")

		//audio("../audio/fire.mp3")
	case "shooter":
		triggerButton(shooterOutPin)
		audio("./audio/shooter.mp3")

	case "environmental":
		triggerButton(envOutPin)
		audio("./audio/env.mp3")
	}

}


func listenForSmoke() {

	for{

		if !smokePin.Read() == true {
			log.Output(1, "SMOKE DETECTED")
			writeToGPIO("fire")
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


		// parses emergency type from post request
		emergencyType := r.Form["emergency"][0]
		fmt.Println(emergencyType)

		// trigger the respective pins
		switch emergencyType {
		case "fire":
			writeToGPIO("fire")
		case "shooter":
			writeToGPIO("shooter")
		case "environmental":
			writeToGPIO("environmental")
		case "safety":
			writeToGPIO("safety")

		}



    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}



func audio(pathToFile string) error {
	log.Output(1, "Playing Audio")
	f, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done


	return err

}
