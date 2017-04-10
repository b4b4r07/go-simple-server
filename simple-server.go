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
	fmt.Fprintln(w, "Hello, World")
}

var (
	// open browser

	// TODO
	portN = flag.String("port", "", "port e.g. 8000,8080")

	// TODO json?

	// TODO sleep?
)

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	var n int
	n = port.Get()
	if *portN != "" {
		n = port.Get(*portN)
	}
	if !clipboard.Unsupported {
		err := clipboard.WriteAll(fmt.Sprintf("curl localhost:%d", n))
		if err != nil {
			panic(err)
		}
	}
	log.Printf("Serving %d...", n)
	// TODO: access log to stdout
	// http.ListenAndServe(fmt.Sprintf(":%d", n), nil)
	http.Serve(l, hlog.Wrap(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "no-store")
		http.ServeFile(w, r, "."+r.URL.Path)
	}))
}
