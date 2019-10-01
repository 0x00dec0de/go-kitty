package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, hostname)
	})
	fmt.Println("Listen on 0.0.0.0:8888")
	log.Fatal(http.ListenAndServe("0.0.0.0:8888", nil))
}
