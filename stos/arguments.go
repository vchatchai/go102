package stos

import (
	"fmt"
	"os"
)

func Arguments() {
	args := os.Args

	fmt.Println(args)

	//The first argument, zero item from slice
	programNme := args[0]
	fmt.Println("The binary name is: %s \n", programNme)

	//The rest of the  argments could be obtained
	//by omitting the first arguments.
	otherArgs := args[1:]
	fmt.Println(otherArgs)
	for idx, arg := range otherArgs {
		fmt.Printf("Arg %d = %s\n", idx, arg)
	}
}
