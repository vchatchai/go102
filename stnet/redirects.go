package stnet

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type RedirectServer struct {
	redirectCount int
}

func (s *RedirectServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.redirectCount++
	fmt.Println("Recieved header:", req.Header)
	http.Redirect(rw, req, fmt.Sprintf("/redirect%d", s.redirectCount), http.StatusTemporaryRedirect)
}

func Redirects() {
	s := http.Server{
		Addr:    addr,
		Handler: &RedirectServer{0},
	}

	go s.ListenAndServe()

	time.Sleep(1 * time.Second)
	client := http.Client{}

	redirectCount := 0

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		fmt.Println("Redirected")
		if redirectCount > 2 {
			return fmt.Errorf("Too many redirects")
		}

		req.Header.Set("Known-redirects", fmt.Sprintf("%d", redirectCount))
		redirectCount++
		for _, prReq := range via {
			fmt.Printf("Previous request: %v\n", prReq.URL)
		}
		return nil
	}

	_, err := client.Get(strings.Join([]string{"http://", addr}, ""))
	if err != nil {
		panic(err)
	}

}
