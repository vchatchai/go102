package stnet

import (
	"fmt"
	"log"
	"net/http"
)

func HandleGroup() {
	log.Println("Starting Server...")
	mainMux := http.NewServeMux()
	mainMux.Handle("/api/", http.StripPrefix("/api", restModule()))
	mainMux.Handle("/ui/", http.StripPrefix("/ui", uiModule()))

	if err := http.ListenAndServe(":8080", mainMux); err != nil {
		panic(err)
	}
}

func restModule() http.Handler {
	//Separate Mux for all REST
	restApi := http.NewServeMux()
	restApi.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"id":1, "name":"John"}]`)
	})
	return restApi
}

func uiModule() http.Handler {
	//Separate Mux for all UI
	ui := http.NewServeMux()
	ui.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `<html><body>Hello from UI</body></html>`)
	})
	return ui
}
