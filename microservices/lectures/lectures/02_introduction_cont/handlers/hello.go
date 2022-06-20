package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// Implement the `ServeHTTP` interface: https://pkg.go.dev/net/http?utm_source=gopls#Handler
//	type Handler interface {
//		ServeHTTP(ResponseWriter, *Request)
//	}
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Client access")

	// Read the request body
	if d, err := ioutil.ReadAll(r.Body); err != nil {
		http.Error(w, "Oops!", http.StatusBadRequest)
		return
	} else {
		log.Printf("Get data: %s\n", d)

		// Response
		fmt.Fprintf(w, "Hello %s", d)
	}
}
