package stnet

import (
	"html/template"
	"net/http"
)

func Template() {
	tp, err := template.ParseFiles("template.tpl")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		err := tp.Execute(rw, "John Doe")
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
