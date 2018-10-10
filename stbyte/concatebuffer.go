package stbyte

import (
	"bytes"
	"log"
)

func ConcatBuffer(strings []string) {
	buffer := bytes.Buffer{}

	for _, val := range strings {
		buffer.WriteString(val)
	}

	log.Println(buffer.String())
}
