package ststring

import (
	"fmt"
	"strings"
)

func Fields(value string) {

	words := strings.Fields(value)
	for idx, word := range words {
		fmt.Printf("Word %d is : %s\n", idx, word)
	}

}
