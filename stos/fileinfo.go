package stos

import (
	"fmt"
	"os"
)

func FileInfo() {
	f, err := os.Open("flatfile.txt")
	if err != nil {
		panic(err)
	}
	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("File name: %v\n", fi.Name())
	fmt.Printf("Is Director: %t\n", fi.IsDir())
	fmt.Printf("Size: %d\n", fi.Size())
	fmt.Printf("Mode: %v\n", fi.Mode())
}
