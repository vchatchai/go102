package charmap

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func Charset() {

	word := "This is sample text with runes Š \n"

	fmt.Printf("word: %s \n", word)
	encoder := charmap.Windows1252.NewEncoder()
	s, e := encoder.String(word)
	if e != nil {
		panic(e)
	}
	fmt.Printf("word is encoded: %s", s)

	ioutil.WriteFile("example.txt", []byte(s), os.ModePerm)

	f, e := os.Open("example.txt")
	if e != nil {
		panic(e)
	}

	defer f.Close()

	decoder := charmap.Windows1252.NewDecoder()
	reader := decoder.Reader(f)

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
