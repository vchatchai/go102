package stos

import (
	"fmt"
	"os"
)

func FileMode() {
	f, err := os.Create("test.file")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Printf("File permissions %v\n", fi.Mode())

	err = f.Chmod(0777)
	if err != nil {
		panic(err)
	}

	fi, err = f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Printf("File permissions %v\n", fi.Mode())
}
