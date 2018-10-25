package stos

import (
	"fmt"
	"io"
	"os"
)

func Stdouterr() {
	io.WriteString(os.Stdout, "This is string to standard out.\n")

	io.WriteString(os.Stderr, "This is string to standard error output.\n")
	buf := []byte{0xAF, 0xFF, 0xFF}
	for i := 0; i < 200; i++ {
		if _, e := os.Stdout.Write(buf); e != nil {
			panic(e)
		}
	}
	fmt.Fprintf(os.Stdout, "\n")
}
