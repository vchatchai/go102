package stnet

import (
	"fmt"
	"net/http"
)

func TLS() {

	fmt.Println("Starting HTTP server on port 8080")

	s := &http.Server{Addr: ":8080", Handler: SimpleHTTP{}}

	err := s.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		panic(err)
	}

}
