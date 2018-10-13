package stregexp

import (
	"fmt"
	"regexp"
)

func FindSubString(refString string) {
	emailRegexp := regexp.MustCompile("[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\\.[a-z]{1,}")
	first := emailRegexp.FindString(refString)
	fmt.Printf("First:")
	fmt.Println(first)

	all := emailRegexp.FindAllString(refString, -1)
	fmt.Println("All:")
	for _, val := range all {
		fmt.Println(val)
	}
}
