package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/atotto/clipboard"
	"github.com/b4b4r07/go-simple-server/port"
)

var (
	json = flag.Bool("json", false, "Output the response as JSON")
)

func handler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, World"
	if *json {
		msg = fmt.Sprintf(`{"message": "%s"}`, msg)
	}
	fmt.Fprintf(w, msg)
}

func logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)

	n := port.Get(8000, 8080)
	host := "localhost" + port.WithColon(n)

	// save host to the system clipboard
	if !clipboard.Unsupported {
		err := clipboard.WriteAll(host)
		if err != nil {
			panic(err)
		}
	}

	log.Printf("Start to serve %s", host)
	http.ListenAndServe(port.WithColon(n), logger(http.DefaultServeMux))
}
