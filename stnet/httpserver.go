package stnet

import (
	"fmt"
	"net/http"
)

type SimpleHTTP struct {
}

func (s SimpleHTTP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World")
}

func HttpServer() {
	fmt.Println("Starting HTTP server on port 8080")
	s := &http.Server{
		Addr:    ":8080",
		Handler: SimpleHTTP{},
	}

	s.ListenAndServe()
}
