package stos

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func OpenFile() {
	f, err := os.Open("pid.go")

	if err != nil {
		panic(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("### File conent ###\n %s\n", string(c))
	f.Close()

	f, err = os.OpenFile("test.tmp", os.O_CREATE|os.O_RDWR, os.ModePerm)

	if err != nil {
		panic(err)
	}

	io.WriteString(f, string(c))
	f.Close()
}
