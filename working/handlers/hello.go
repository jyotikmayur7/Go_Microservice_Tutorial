package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// Defining a function on a struct
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, error := io.ReadAll(r.Body)

	if error != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
	}

	fmt.Fprintf(rw, "Hello %s", d)
}
