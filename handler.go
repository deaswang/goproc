package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"strings"
)

var tokenList map[string]bool

func SetToken(buf *bufio.Reader) error {
	tokenList = make(map[string]bool)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		tokenList[line] = true
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

// add auth for api
func withAuth(needToken bool, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			log.Println("Not allowed methed: ", r.Method, " from ", r.RemoteAddr)
			return
		}
		if needToken {
			token := r.Header.Get("token")
			if value, ok := tokenList[token]; !(ok && value) {
				w.WriteHeader(http.StatusUnauthorized)
				log.Println("Unauthorized request from ", r.RemoteAddr)
				return
			}
		}
		w.Header().Set("Content-Type", "application/json")
		log.Println("request "+r.RequestURI+" from ", r.RemoteAddr)
		h(w, r)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("index")
}
