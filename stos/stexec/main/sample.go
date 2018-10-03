package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Simple run")
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}
