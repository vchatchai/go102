package strconv

import (
	"fmt"
	"strconv"
)

const bin = "100001"
const hex = "2f"
const intString = "12"
const floatString = "12.3"

func StringToNumber() {

	res, err := strconv.Atoi(intString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed integer: %d\n", res)

	res64, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed heaxdecima: %d\n", res64)

	resBin, err := strconv.ParseInt(bin, 2, 32)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Parsed bin: %d\n", resBin)

	resFloat, err := strconv.ParseFloat(floatString, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed float: %.5f\n", resFloat)
}
