package ststring

import (
	"fmt"
	"strings"
)

func Contain(word string) {
	lookfor := "lamb"
	contain := strings.Contains(word, lookfor)
	fmt.Printf(`This "%s" contains "%s": %t %c`, word, lookfor, contain, '\n')

	lookfor = "wolf"
	contain = strings.Contains(word, lookfor)
	fmt.Printf(`The "%s" contains "%s": %t %c`, word, lookfor, contain, '\n')

	startWith := "Mary"
	starts := strings.HasPrefix(word, startWith)
	fmt.Printf(`The "%s" starts with "%s": %t 
	`, word, startWith, starts)

	endWith := "lamb"
	ends := strings.HasSuffix(word, endWith)
	fmt.Printf(`The "%s" ends with "%s": %t %c`, word, endWith, ends, '\n')
}
