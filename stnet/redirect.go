package stnet

import (
	"fmt"
	"log"
	"net/http"
)

func Redirect() {
	log.Println("Server is starting...")
	http.Handle("/secured/handle", http.RedirectHandler("/login", http.StatusTemporaryRedirect))
	http.HandleFunc("/secured/handlefunc", func(rw http.ResponseWriter, r *http.Request) {
		http.Redirect(rw, r, "/login", http.StatusTemporaryRedirect)
	})
	http.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Welcome user! Please login!\n")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
