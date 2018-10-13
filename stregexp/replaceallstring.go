package stregexp

import (
	"fmt"
	"regexp"
)

func ReplaceAllString(value string) {

	regex := regexp.MustCompile("l[a-z]+")
	out := regex.ReplaceAllString(value, "replacement")
	fmt.Println(out)
}
