package stnet

import (
	"fmt"
	"net/http"
)

/*
curl http://localhost:8080/user
curl http://localhost:8080/items/clothes
curl http://localhost:8080/admin/ports
*/
func Handle() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintln(w, "User GET")
		}
		if r.Method == http.MethodPost {
			fmt.Fprintln(w, "User Post")
		}
	})

	itemMux := http.NewServeMux()
	itemMux.HandleFunc("/items/clothes", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Cloths")
	})
	mux.Handle("/items/", itemMux)

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/ports", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "ports")
	})

	mux.Handle("/admin/", http.StripPrefix("/admin", adminMux))

	http.ListenAndServe(":8080", mux)
}
