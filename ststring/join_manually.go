package ststring

import (
	"fmt"
	"strings"
)

type JoinFunc func(piece string) string

func JoinManually(value []string) {
	jF := func(p string) string {
		if strings.Contains(p, "INSURANCE") {
			return " OR "
		}

		return " AND "
	}

	result := JoinWithFunc(value, jF)
	fmt.Printf(selectBase, result)
}

func JoinWithFunc(value []string, joinFunc JoinFunc) string {
	concatenate := value[0]
	for _, val := range value[1:] {
		concatenate = concatenate + joinFunc(val) + val
	}
	return concatenate
}
