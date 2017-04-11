// All-rounder for demo server
// TODO
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

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
		var p []interface{}
		if strings.Contains(*portN, ",") {
			for _, i := range strings.Split(*portN, ",") {
				p = append(p, interface{}(i))
			}
		} else {
			p = append(p, interface{}(*portN))
		}
		n = port.Get(p...)
	}
	if !clipboard.Unsupported {
		err := clipboard.WriteAll(fmt.Sprintf("curl localhost:%d", n))
		if err != nil {
			panic(err)
		}
	}
	log.Printf("Serving %d...", n)
	// TODO: access log to stdout
	http.ListenAndServe(fmt.Sprintf(":%d", n), nil)
}
