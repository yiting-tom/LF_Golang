package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (h *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Client leave")

	w.Write([]byte("Goodbye ~"))
}
