package ststring

import (
	"log"
	"strings"
)

const selectBase = "SELECT * FROM user WHERE %s"

func Join(values []string) {
	word := strings.Join(values, " AND ")
	log.Printf(selectBase+"\n", word)
}
