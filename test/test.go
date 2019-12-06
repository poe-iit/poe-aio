package main

import (
	"net/http"
	"fmt"
	"net/url"
)

func main() {
	MakeRequest()
}

func MakeRequest() {

	emergencyType := "fire"
	APIURL := "http://192.168.2.50:12345/button"

	response, err := http.PostForm(APIURL,
      url.Values{"emergency": {emergencyType}})
	
	

	//okay, moving on...
	if err != nil {
		fmt.Println(err)
	}

	
	fmt.Println(response)
	}
