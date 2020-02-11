package main

import (
	"os"
	"fmt"
	"net/http"
)

var version string

func init() {
	version = os.Getenv("VERSION")
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		hostname, err := os.Hostname()

		if err != nil {
			fmt.Fprintf(w, "Error getting hostname\n")
			return
		}

		fmt.Fprintf(os.Stdout, "Host: %s, Version: %s\n", hostname, version)
	
		fmt.Fprintf(w, "Host: %s, Version: %s\n", hostname, version)

		// w.WriteHeader(http.StatusInternalServerError)
		// w.Write([]byte("500 - Something bad happened!"))
	})

	http.HandleFunc("/live", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("I'm healthy!"))
	})

	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("I'm ready to serve HTTP requests!"))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Listening on port 8080...")
}