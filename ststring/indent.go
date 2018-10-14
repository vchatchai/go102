package ststring

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func IndentALL() {
	text := "Hi! Go is awesome."
	text = Indent(text, 6)
	fmt.Println(text)

	text = UnIndent(text, 3)
	fmt.Println(text)

	text = UnIndent(text, 10)
	fmt.Println(text)

	text = IndentByRune(text, 10, '.')
	fmt.Println(text)
}

func IndentByRune(input string, indent int, r rune) string {
	buffer := bytes.Buffer{}
	buffer.WriteString(strings.Repeat(string(r), indent))
	buffer.WriteString(input)
	return buffer.String()
}

func Indent(input string, indent int) string {
	padding := indent + len(input)
	buffer := bytes.Buffer{}
	buffer.WriteRune('%')
	buffer.WriteString(strconv.Itoa(padding))
	buffer.WriteString("s")
	return fmt.Sprintf(buffer.String(), input)
}

func UnIndent(input string, indent int) string {
	count := 0
	for _, val := range input {
		if unicode.IsSpace(val) {
			count++
		}
		if count == indent || !unicode.IsSpace(val) {
			break
		}
	}
	return input[count:]
}
