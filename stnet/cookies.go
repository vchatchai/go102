package stnet

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const cookieName = "X-Cookie"

func Cookies() {
	log.Println("Server is starting...")
	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		c := &http.Cookie{
			Name:    cookieName,
			Value:   "Go is awesome.",
			Expires: time.Now().Add(time.Hour),
			Domain:  "localhost",
		}
		http.SetCookie(w, c)
		fmt.Fprintln(w, "Coook is set!")
	})
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		val, err := r.Cookie(cookieName)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "Cookie is: %s \n", val.Value)
		fmt.Fprintf(w, "Other cookies")
		for _, v := range r.Cookies() {
			fmt.Fprintf(w, "%s => %s\n", v.Name, v.Value)
		}
	})

	http.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
		val, err := r.Cookie(cookieName)
		if err != nil {
			panic(err)
		}
		val.MaxAge = -1
		http.SetCookie(w, val)
		fmt.Fprintln(w, "Cookie is removed!")

	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
