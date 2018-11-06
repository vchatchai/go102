package stnet

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type PostStringServer string

func (s PostStringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Printf("Recieved form data: %v\n", req.Form)
	rw.Write([]byte(string(s)))

}

func createServerPost(add string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: PostStringServer("Hello world"),
	}
}

var wg sync.WaitGroup

func Post() {
	s := createServerPost(addr)
	// wg.Add(1)

	go s.ListenAndServe()
	time.Sleep(10 * time.Second)
	// wg.Done()
	// wg.Wait()

	simplePost()
	useRequest()

}

func simplePost() {
	var host string = "http://" + addr

	fmt.Println("Host:", host)

	resp, err := http.Post(host,
		"application/x-www-form-urlencoded",
		strings.NewReader("name=Radek&surname=Sohlich"))
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
	fmt.Println("Response from server ", string(data))
}

func useRequest() {
	hc := http.Client{}
	form := url.Values{}
	form.Add("name", "RaDek")
	form.Add("surname", "Sohlich")

	req, err := http.NewRequest("POST", "http://"+addr, strings.NewReader(form.Encode()))

	req.Header.Add("Content-Type",
		"application/x-www-form-urlencoded")

	res, err := hc.Do(req)

	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	res.Body.Close()
	fmt.Println("Response form server:", string(data))
}
