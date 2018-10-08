package ststring

import (
	"log"
	"strings"
)

func FieldFunction(value string) {
	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%,_", r)
	}

	words := strings.FieldsFunc(value, splitFunc)
	for idx, word := range words {
		log.Printf("Word %d is: %s\n", idx, word)
	}
}
