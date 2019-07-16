package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mux sync.Mutex
var count int

func main() {
	http.HandleFunc("/", hander)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func hander(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	count++
	mux.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	fmt.Fprintf(w, "count %d\n", count)
	mux.Unlock()
}
