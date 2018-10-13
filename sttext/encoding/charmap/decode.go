package charmap

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func Decode() {
	f, err := os.Open("win1250.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	content := string(b)

	fmt.Println("Without decode:" + content)
	decoder := charmap.Windows1250.NewDecoder()
	reader := decoder.Reader(strings.NewReader(content))
	b, err = ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded:" + string(b))
}
