package stnet

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type RequestStringServer string

func (s RequestStringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Printf("Received form data: %v\n", req.Form)
	fmt.Printf("Received header: %v\n", req.Header)
	rw.Write([]byte(string(s)))
}

func RequestCreateServer(add string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: RequestStringServer(" Hello world "),
	}
}

func Request() {
	s := RequestCreateServer(addr)
	go s.ListenAndServe()

	time.Sleep(2 * time.Second)

	form := url.Values{}
	form.Set("id", "5")
	form.Set("name", "Wolfgang")

	req, err := http.NewRequest(http.MethodPost,
		"http://"+addr,
		strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	res.Body.Close()

	fmt.Printf("Response from server %v\n", string(data))
}
