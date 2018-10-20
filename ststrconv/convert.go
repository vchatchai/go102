package strconv

import (
	"fmt"
	"strconv"
)

const convbin = "10111"
const convhex = "1A"
const convoct = "12"
const convdec = "10"
const convfloatNum = 16.123557

func Convert() {
	v, _ := ConvertInt(convbin, 2, 16)
	fmt.Printf("Binary value %s converted to hex: %s\n", convbin, v)

	v, _ = ConvertInt(convhex, 16, 10)
	fmt.Printf("Hex value %s converted to dec: %s\n", convhex, v)

	v, _ = ConvertInt(convoct, 8, 16)
	fmt.Printf("Oct value %s converted to hex: %s\n", convoct, v)

	v, _ = ConvertInt(convdec, 10, 8)
	fmt.Printf("Dec value %s converted to oct: %s\n", convdec, v)
}

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}
