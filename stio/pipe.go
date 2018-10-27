package stio

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func Pipe() {
	pReader, pWriter := io.Pipe()

	cmd := exec.Command("echo", "Hello Go!\n This is example")
	cmd.Stdout = pWriter

	go func() {
		defer pReader.Close()
		if _, err := io.Copy(os.Stdout, pReader); err != nil {
			log.Fatal(err)
		}
	}()

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
