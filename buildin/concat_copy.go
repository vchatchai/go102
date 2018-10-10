package buildin

import "fmt"

func ConcatCopy(values []string) {
	bs := make([]byte, 100)
	b1 := 0

	for _, val := range values {
		b1 += copy(bs[b1:], []byte(val))
	}
	fmt.Println(string(bs[:]))
}
