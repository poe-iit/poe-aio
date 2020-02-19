package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
	//"net/http/httputil"
	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"net/url"
	"crypto/tls"
	"crypto/x509"
	"crypto/rand"
	//"math/rand"
	//"time"
)


func main() {

	 // ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	 // if !ok {
		// 	 panic("failed to parse root certificate")
	 // }

	startWebApp()

}


func startWebApp() {
	router := mux.NewRouter()
	//router.HandleFunc("/button", handleRequests)
	router.HandleFunc("/button", handleRequests)
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("../website").HTTPBox())) // starts the web UI
  //router.Schemes("https")
	log.Output(1, "Started Web UI and http server")

	server := &http.Server{
			Addr:         ":8081",
			Handler:      router,
		}

	server.TLSConfig = configTLS()

	//log.Fatal(http.ListenAndServe(":8080", router))

	// localProxyUrl, _ := url.Parse("http://127.0.0.1:8100/")
	// localProxy := httputil.NewSingleHostReverseProxy(localProxyUrl)
	// http.Handle("/", localProxy)

	log.Println("About to listen on 8081. Go to https://localhost:8081")
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
	//log.Fatal(http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", router))

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
	// MIIDPjCCAiYCCQC2wBvwPhDsVjANBgkqhkiG9w0BAQsFADBhMQswCQYDVQQGEwJV
	// UzELMAkGA1UECAwCSUwxEDAOBgNVBAcMB0NoaWNhZ28xDDAKBgNVBAoMA1BPRTEM
	// MAoGA1UECwwDNTEwMRcwFQYDVQQDDA5sb2NhbGhvc3Q6ODA4MTAeFw0yMDAyMTkw
	// MTQ1NThaFw0yMTAyMTgwMTQ1NThaMGExCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJJ
	// TDEQMA4GA1UEBwwHQ2hpY2FnbzEMMAoGA1UECgwDUE9FMQwwCgYDVQQLDAM1MTAx
	// FzAVBgNVBAMMDmxvY2FsaG9zdDo4MDgxMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
	// MIIBCgKCAQEAsm3vDUSFqSXcdzT+2YT1ZwLoQ/HyOMsg1e/OfPp0tvDpoqhG4+Wo
	// KrcAWwh3/XVBg/NwMwy9mqkinPeRsSaFMrlWMf1SRSxNKkx1HywCMzv5DTQJrU8L
	// Hdqh9MaaA4sAEKTBKeeexvtQT37dScuwqMo0We73mHegv+ux9ByNSIW8yA1VbjtC
	// 9Cyf8eWVLibFZA6zIS6Qf+yEpr3qZ0Ukt3fg1FZFT64sJ1XVJvG3CgMq+P9Rb+Zc
	// XWBlc0NQpt0H/Ok2TATQXPbPyXAUN3JJnhtmjNowiSJXsxAo6I+mNB0eyLeDSATu
	// s5weUubGXlO+P7Ey06dl2gXGQIx1JULXUQIDAQABMA0GCSqGSIb3DQEBCwUAA4IB
	// AQCCOE8znV66/LpmNsZm9gPM1PeVN4GkFFn6uZMhJTatwlTuO95FV7SCtP2vJTMc
	// +8Fe+Mp1QlXxCxPVUjVzZZwQkPKeV83Ht09kFloQNFD+afb9c8Q1IxC6YtBgt4VG
	// M+CfGCTl208qwOL7TiP5hAzUCiAh+tMYnzcl16q8BtcBrhDeO/LfSQDBkRgHVr1/
	// zJJ8E1Cn5uE4UTN43mUpxUE7VdC3KrKC9lcx/H5YrYer2E/gRb0Rk8oYM7KShqPm
	// 5f6KJFgcZNPGsXuQw31YSsi1dt/QdnZg7u9dDS4TkHzils/7zlNtCWwvVrWZ9TU/
	// vKmXKUCKTAEIzwlS0LrqMNrs
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
  fmt.Println("Attempting to send message...")
	APIURL := "https://localhost:12345/lights"

	response, err := http.PostForm(APIURL,
	  url.Values{"emergency": {emergencyType}})

	if err != nil {
		return err
	}

	fmt.Println("response is: " , response)

	return err
}
