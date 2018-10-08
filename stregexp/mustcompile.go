package stregexp

import (
	"log"
	"regexp"
)

func MustCompile(value string) {
	words := regexp.MustCompile("[*,%_]{1}").Split(value, -1)
	for idx, word := range words {
		log.Printf("Word %d is: %s\n", idx, word)
	}
}
