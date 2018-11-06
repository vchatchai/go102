package stnet

import (
	"fmt"
	"net/http"
)

func Header() {
	header := http.Header{}

	header.Set("Auth-X", "abcdef1234")
	header.Add("Auth-X", "defghijkl")

	fmt.Println(header)

	resSlice := header["Auth-X"]
	fmt.Println(resSlice)

	resFist := header.Get("Auth-X")
	fmt.Println(resFist)

	//replace all existing values with this one
	header.Set("Auth-X", "newValue")
	fmt.Println(header)

	header.Del("Auth-X")
	fmt.Println(header)
}
