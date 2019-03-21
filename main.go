package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// This is a simple HTTP server that provides two endpoints:
// GET /uuid returns a uuidv4 value
// POST /sha returns the sha256 hash of the request body
//
// This is application is designed for testing other systems.

var port int

func main() {
	flag.IntVar(&port, "port", 8000, "tcp port to listen on")
	flag.Parse()

	http.Handle("/uuid", logger(http.HandlerFunc(uuidHandler)))
	http.Handle("/sha", logger(http.HandlerFunc(shaHandler)))

	log.Println("starting server")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func uuidHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", uuid.New())
}

func shaHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	sum := sha256.Sum256(body)
	fmt.Fprintf(w, "%x", sum)
}

func logger(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer log.Printf("%s %s %s", r.Method, r.URL, r.RemoteAddr)
		f.ServeHTTP(w, r)
	})
}
