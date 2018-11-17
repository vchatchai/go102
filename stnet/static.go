package stnet

import (
	"net/http"
)

func Static() {
	fileSrv := http.FileServer(http.Dir("./"))
	fileSrv = http.StripPrefix("/html", fileSrv)
	http.HandleFunc("/welcome", serveWelcome)
	http.Handle("/html/", fileSrv)
	http.ListenAndServe(":8080", nil)
}

func serveWelcome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "welcome.txt")
}
