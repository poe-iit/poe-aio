package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
  //"net/http/httputil"
	//"net/url"
//  "crypto/tls"
 // "crypto/x509"
 // "io/ioutil"
	//"time"

	//"github.com/warthog618/gpio"
	//"os"
	//"github.com/faiface/beep"
	//"github.com/faiface/beep/mp3"
	//"github.com/faiface/beep/speaker"

	"github.com/gorilla/mux"
	"crypto/x509"
	"crypto/tls"
	"crypto/rand"

)



/*var fireOutPin *gpio.Pin
var shooterOutPin *gpio.Pin
var envOutPin *gpio.Pin
var smokePin *gpio.Pin*/

func main() {

	/*err := initPins()
	if err != nil {
		log.Fatal(err.Error())
	}*/
	//defer gpio.Close()

	//go listenForSmoke()
	log.Output(1, "Listening for smoke")

	// start the http server to listen to requests from the server
	router := mux.NewRouter()
	router.HandleFunc("/lights", handleRequests)
  //log.Fatal(http.ListenAndServe(":12345", router))

  // localProxyUrl, _ := url.Parse("http://127.0.0.1:8100/")
	// localProxy := httputil.NewSingleHostReverseProxy(localProxyUrl)
	// http.Handle("/", localProxy)

	server := &http.Server{
			Addr:         ":12345",
			Handler:      router,
		}

	server.TLSConfig = configTLS()

	log.Println("About to listen on 12345. Go to https://localhost:12345")
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
	//log.Fatal(http.ListenAndServeTLS(":12345", "cert.pem", "key.pem", router))


}

func configTLS() *tls.Config {
	TLSConfig := &tls.Config{}
	// TLSConfig.CipherSuites = []uint16{
	// 	tls.TLS_FALLBACK_SCSV,
	// 	tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	// 	tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
	// 	tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
	// 	tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	// 	tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
	// 	tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
	// 	tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
	// 	tls.TLS_RSA_WITH_AES_256_CBC_SHA,
	// 	tls.TLS_RSA_WITH_AES_128_CBC_SHA,
	// 	tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
	// }

	TLSConfig.Rand = rand.Reader
	TLSConfig.MinVersion = tls.VersionTLS10
	TLSConfig.SessionTicketsDisabled = false
	TLSConfig.InsecureSkipVerify = false
	TLSConfig.ClientAuth = tls.VerifyClientCertIfGiven
	TLSConfig.PreferServerCipherSuites = true
	TLSConfig.ClientSessionCache = tls.NewLRUClientSessionCache(1000)
	TLSConfig.RootCAs = loadCertificates()

	return TLSConfig
}


func loadCertificates() *x509.CertPool {

	// rootPEM := `-----BEGIN CERTIFICATE-----
	// MIIDQDCCAigCCQDKGKeSZN7AbTANBgkqhkiG9w0BAQsFADBiMQswCQYDVQQGEwJV
	// UzELMAkGA1UECAwCSUwxEDAOBgNVBAcMB0NoaWNhZ28xDDAKBgNVBAoMA1BPRTEM
	// MAoGA1UECwwDNTEwMRgwFgYDVQQDDA9sb2NhbGhvc3Q6MTIzNDUwHhcNMjAwMjE5
	// MDE0NjI2WhcNMjEwMjE4MDE0NjI2WjBiMQswCQYDVQQGEwJVUzELMAkGA1UECAwC
	// SUwxEDAOBgNVBAcMB0NoaWNhZ28xDDAKBgNVBAoMA1BPRTEMMAoGA1UECwwDNTEw
	// MRgwFgYDVQQDDA9sb2NhbGhvc3Q6MTIzNDUwggEiMA0GCSqGSIb3DQEBAQUAA4IB
	// DwAwggEKAoIBAQCXiUJkn8dr0KPPUStDSn3ym0atEB9yqvqAxHYewkVAwvNBD9F7
	// a3uYfWR6oAUIOJkfLwdW5N8v/+MKH15jMmbLNyCM3whcwg7upXY/aBMSw4PWcMAY
	// RjhAn788P9MEjcHJY02TSx/Bnaz2PXwtPcR/d11fpeS0OyPK+fO2AMxr7wFIjipY
	// mufV49QdU6RK1Yie9wcd5//aCW3GFJVRWW1HfTR65WjdAmIH1MtmmRyH8CVD6wTb
	// eZjpeNElaObT8afAGjAS0AUdk7qqd+wrN+6hxhwxo8/OjsMVu1woGh1/hmkY057H
	// +m5doAWSl641Ed+8hYn6ZY07ep77CdOXPA7bAgMBAAEwDQYJKoZIhvcNAQELBQAD
	// ggEBAFUDBzEZ6LYtDd/LKXDvn8dvmamnX/ZpXGRWXRLdXqGaLHelbrpFxIq64b8s
	// YtuYxP8aGLZ7AB7efnDFNq5UjP0ce+R/8xyl2UjQKnpKcNNMUctOrzKv6wWafrun
	// qAeHAutKZFuI1t1r3Hf/vkHE4hjDKlMvGWVkQXU6KVgtAAx1OUAZDiR278wJHQlU
	// j+fPfYyazl+kWCEzrR1e4NSNUdDkQPT6SQy0L1MxjnySGu0yuCcHqwsPEt4lBCX0
	// 9j0Fe2aq2QhaqDiQd2eDo/gMCWE++D/LGvJaYRL0N8L8CtcvXQTfK04cKsr201Nk
	// luVD4BNNQ5dj0rGtXDBkaW0WERI=
	// -----END CERTIFICATE-----`

	pem, err := ioutil.ReadFile("cert.pem")

	if err != nil {
		panic(err)
	}

	// rootCAs := x509.NewCertPool()
	// rootCAs.AppendCertsFromPEM([]byte(rootPEM))s

	rootCertPool := x509.NewCertPool()
	if !rootCertPool.AppendCertsFromPEM(pem) {
		panic("Failed appending certs")
	}

	return rootCertPool
}

/*
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
*/

/*
func writeToGPIO(emergencyType string) {
	log.Output(1, "Writing to GPIO")

	switch emergencyType {
	case "Fire":
		triggerButton(fireOutPin)
		audio("./audio/fire.mp3")

		//audio("../audio/fire.mp3")
	case "Shooter":
		triggerButton(shooterOutPin)
		audio("./audio/shooter.mp3")

	case "Enviormental":
		triggerButton(envOutPin)
		audio("./audio/env.mp3")
	}
}
*/
/*
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
*/
/*
func triggerButton(pin *gpio.Pin) {

	pin.Write(gpio.Low)
	time.Sleep(250 * time.Millisecond)
	pin.Write(gpio.High)


}*/




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

/*
		// trigger the respective pins
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
*/


    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}


/*
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

} */
