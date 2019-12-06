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

	
	go startWebApp()
	log.Output(1, "Started Web UI and http server")



}



func startWebApp() {
	router := mux.NewRouter()
	router.HandleFunc("/button", handleRequests)
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	log.Fatal(http.ListenAndServe(":12345", router))
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

		err = sendMessage(emergencyType)

		if err != nil {
			log.Output(1, err.Error())

		}

        
    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}



func sendMessage(emergencyType string) (err error) {

	APIURL := "http://192.168.2.51:12345/lights"

	response, err := http.PostForm(APIURL,
      url.Values{"emergency": {emergencyType}})
	
	

	//okay, moving on...
	if err != nil {
		return err
	}

	
	fmt.Println(response)
	}