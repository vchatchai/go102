package stos

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile() {
	fmt.Println("### Read as reader")
	f, err := os.Open("pid.go")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}

	fmt.Println(wr.String())

	fmt.Println("### Read File ###")
	fContent, err := ioutil.ReadFile("pid.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fContent))

}
