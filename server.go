package main

import (
	
	"log"

	"fmt"
	"net/http"
	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"net/url"
)




func main() {
	startWebApp()
}



func startWebApp() {
	router := mux.NewRouter()
	router.HandleFunc("/button", handleRequests)
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("./website").HTTPBox())) // starts the web UI
	log.Output(1, "Started Web UI and http server")
	log.Fatal(http.ListenAndServe(":5050", router))
}


func handleRequests(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "handling request from webpage")

 
    switch r.Method {
	case "POST":
		
		log.Output(1, "Post Request Recieved")
		err := r.ParseForm()
		
		if err != nil {
			log.Output(1, err.Error())
		}

		
		emergencyType := r.Form["emergency"][0] // get the emergency out of the POST request
		err = sendMessage(emergencyType)  // send the message to the ceiling client

		if err != nil {
			log.Output(1, err.Error())

		}

        
    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}



func sendMessage(emergencyType string) (err error) {

	APIURL := "http://127.0.0.1:12345/lights"

	response, err := http.PostForm(APIURL,
	  url.Values{"emergency": {emergencyType}})

	if err != nil {
		return err
	}

	fmt.Println(response)

	return err
}
