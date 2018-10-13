package ststring

import (
	"fmt"
	"strings"
)

func Replacer() {
	replacer := strings.NewReplacer("lamb", "wolf", "Mary", "Jack")
	out := replacer.Replace(refString)
	fmt.Println(out)
}
