package main

import (
	"bufio"
	"flag"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

// VERSION app version
const VERSION = "0.2.1"

var portFlag int
var tokenFlag string
var certFlag string
var keyFlag string

// init default flag and flag value
func init() {
	const (
		defaultPort  = 8809
		defaultToken = "token.txt"
		defaultCert  = "ssl.csr"
		defaultKey   = "ssl.key"
	)
	flag.IntVar(&portFlag, "port", defaultPort, "The proxy server port.")
	flag.StringVar(&tokenFlag, "token", defaultToken, "The token file for authentication.")
	flag.StringVar(&certFlag, "cert", defaultCert, "The cert file name for tls.")
	flag.StringVar(&keyFlag, "key", defaultKey, "The key file name for tls.")
}

// read token access auth
func readToken() error {
	f, err := os.Open(tokenFlag)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	return setToken(buf)
}

func main() {
	flag.Parse()

	log.Printf("goproc v%s (built %s)\n", VERSION, runtime.Version())
	err := readToken()
	var router http.Handler
	if err != nil {
		log.Println("WARNING: without token")
		router = InitRouter(false)
	} else {
		router = InitRouter(true)
	}

	_, certErr := os.Stat(certFlag)
	_, keyErr := os.Stat(keyFlag)

	if (certErr == nil || os.IsExist(certErr)) && (keyErr == nil || os.IsExist(keyErr)) {
		log.Println("Start TLS Server at :" + strconv.Itoa(portFlag))
		err = http.ListenAndServeTLS(":"+strconv.Itoa(portFlag), certFlag, keyFlag, router)
	} else {
		log.Println("Start Server at :" + strconv.Itoa(portFlag))
		err = http.ListenAndServe(":"+strconv.Itoa(portFlag), router)
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}
