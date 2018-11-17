package stnet

import (
	"io"
	"net/http"
)

/*
curl -X GET -I http://localhost:8080/api/users
curl -X GET -H "X-Auth: authenticated" -I http://localhost:8080/api/users
*/
func MiddleWare() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", Secure(func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, `[{"id":"1","login":"ffghi"},{"id":"2","login":"ffghj"}]`)
	}))

	http.ListenAndServe(":8080", mux)

}

func Secure(h http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		sec := r.Header.Get("X-AUTH")
		if sec != "authenticated" {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		h(rw, r)
	}
}
