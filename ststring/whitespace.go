package ststring

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func WhiteSpace() {
	stringToTrime := "\t\t\n Go\tis\t Awesome \t\t"
	trimeResult := strings.TrimSpace(stringToTrime)
	fmt.Println(trimeResult)

	stringWithSpace := "\t\t\n Go\tis\t Awesome \t\t"
	r := regexp.MustCompile("\\s+")
	replace := r.ReplaceAllString(stringWithSpace, " ")
	fmt.Println(replace)

	needSpace := "need space"
	fmt.Println(pad(needSpace, 14, "CENTER"))
	fmt.Println(strings.TrimLeft(needSpace, " "))
	fmt.Println(pad(needSpace, 14, "LEFT"))

}

func pad(input string, padLen int, align string) string {
	inputLen := len(input)
	if inputLen >= padLen {
		return input
	}
	repeat := padLen - inputLen
	var output string

	switch align {
	case "RIGHT":
		output = fmt.Sprintf("% "+strconv.Itoa(-padLen)+"s", input)
	case "LEFT":
		output = fmt.Sprintf("% "+strconv.Itoa(padLen)+"s", input)
	case "CENTER":
		bothRepeat := float64(repeat) / float64(2)
		left := int(math.Floor(bothRepeat)) + inputLen
		right := int(math.Ceil(bothRepeat))
		output = fmt.Sprintf("% "+strconv.Itoa(left)+"s%"+strconv.Itoa(right)+"s", input, "")
	}

	return output
}
