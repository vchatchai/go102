package stnet

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func URL() {
	u := &url.URL{}
	u.Scheme = "http"
	u.Host = "localhost"
	u.Path = "index.html"
	u.RawPath = "id=1&name=John"
	u.User = url.UserPassword("admin", "1234")
	fmt.Printf("Assembled URL:\n%v\n\n\n", u)
	fmt.Println("u.String():", u.String())
	parseURL, err := url.Parse(u.String())
	if err != nil {
		panic(err)
	}
	jsonURL, err := json.Marshal(parseURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Parse URL:", jsonURL)
	fmt.Println(string(jsonURL))
}
