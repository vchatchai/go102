package stfmt

import (
	"bufio"
	"fmt"
	"os"
)

func Scanner() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		txt := sc.Text()
		fmt.Printf("Echo: %s\n", txt)
	}
}
