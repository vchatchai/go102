package stnet

import (
	"fmt"
	"net/http"
)

type PostServer string

func (s PostServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Prior ParseForm: %v\n", req.Form)
	req.ParseForm()
	fmt.Printf("Post ParseForm: %v\n", req.Form)
	fmt.Println("Param1 is : " + req.Form.Get("param1"))
	rw.Write([]byte(string(s)))

}

func ParseForm() {
	srv := http.Server{
		Addr:    ":8080",
		Handler: PostServer("PostServer"),
	}

	fmt.Println("Server is starting...")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
