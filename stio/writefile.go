package stio

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func WriteFile() {
	f, err := ioutil.TempFile("", "test")
	fmt.Println(f.Name())
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString("Go is awesome!\n")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, strings.NewReader("Yeah! Go is great.\n"))
	if err != nil {
		panic(err)
	}
}
