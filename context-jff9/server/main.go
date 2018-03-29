package main

import (
	"flag"
	"fmt"
	//"log"
	"net/http"
	"time"

	"github.com/campoy/justforfunc/context/log"
)

func main() {
	flag.Parse()
	http.HandleFunc("/", log.Decorate(handler))
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Printf("handler started")
	defer log.Printf("handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintln(w, "hello")
}
