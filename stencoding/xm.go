package stencoding

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Book struct {
	Title  string `xml:"title"`
	Author string `xml:"author"`
}

func XML() {
	f, err := os.Open("data.xml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	books := make([]Book, 0)
	decoder := xml.NewDecoder(f)
	// decoder.Decode(&book)
	for {
		tok, _ := decoder.Token()
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "book" {
				var b Book
				decoder.DecodeElement(&b, &tp)
				books = append(books, b)
			}
		}
	}
	fmt.Println(books)

}
