package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Official Documentation: https://golang.org/pkg/net/http/
/*
test the request: curl -d '<text>' localhost:9090/test
It shoud show "Client access ... get data: <text>" at server and show "...Hello <text>" at client
*/

func main() {
	// Define a Controller.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Client access")
	})

	// Define a Controller.
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		// Read request body
		if d, err := ioutil.ReadAll(r.Body); err != nil {
			log.Println(err)
			// Return the `BadRequest` Error
			http.Error(w, "Oops!", http.StatusBadRequest)

			// Or in manually
			// w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("Ooops!"))
			return

		} else {
			log.Printf("Get data: %s\n", d)

			// Response
			fmt.Fprintf(w, "Hello %s", d)
		}
	})

	http.ListenAndServe(":9090", nil)
}
