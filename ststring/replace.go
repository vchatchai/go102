package ststring

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little lamb"
const refString2 = "lamb lamb lamb lamb"

func Repalce() {
	out := strings.Replace(refString, "lamb", "wolf", -1)
	fmt.Println(out)
	out = strings.Replace(refString2, "lamb", "wolf", 2)
	fmt.Println(out)

}
