// All-rounder for demo server
// TODO
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/atotto/clipboard"
	"github.com/b4b4r07/go-simple-server/port"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func logging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	n := port.Get(8000, 8080)
	if !clipboard.Unsupported {
		err := clipboard.WriteAll(fmt.Sprintf("curl localhost:%d", n))
		if err != nil {
			panic(err)
		}
	}
	log.Printf("Serving %d...", n)
	http.ListenAndServe(fmt.Sprintf(":%d", n), logging(http.DefaultServeMux))
}
