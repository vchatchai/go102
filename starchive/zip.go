package starchive

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func Zip() {
	var buff bytes.Buffer
	zipW := zip.NewWriter(&buff)

	f, err := zipW.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte("This is my file content"))
	if err != nil {
		panic(err)
	}

	err = zipW.Close()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("data.zip", buff.Bytes(), os.ModePerm)

	if err != nil {
		panic(err)
	}

	zipR, err := zip.OpenReader("data.zip")
	if err != nil {
		panic(err)
	}

	for _, file := range zipR.File {
		fmt.Println("File " + file.Name + " contains:")
		r, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(os.Stdout, r)
		if err != nil {
			panic(err)
		}

		err = r.Close()

		if err != nil {
			panic(err)
		}
		fmt.Println()
	}
}
