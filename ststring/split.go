package ststring

import (
	"log"
	"strings"
)

func Split(value string) {
	words := strings.Split(value, "_")
	for idx, word := range words {
		log.Printf("Word %d is: %s\n", idx, word)
	}

}
