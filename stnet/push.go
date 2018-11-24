package stnet

import (
	"io"
	"log"
	"net/http"
)

func Push() {
	log.Println("Starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if p, ok := w.(http.Pusher); ok {
			if err := p.Push("/app.css", nil); err != nil {
				log.Printf("Push err: %v", err)
			}
		}
		io.WriteString(w,
			`<html>
				<head>
					<link rel="stylesheet" type="text/css" href="app.css">
				</head>
				<body>
				<p>hello</p>
				</body>
			</html>
		`)
	})
	http.HandleFunc("/app.css", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w,
			`p {
				text-align:center;
				color: red;
			}`)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
