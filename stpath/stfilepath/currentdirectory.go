package stfilepath

import (
	"log"
	"os"
	"path/filepath"
)

func CurrentDirectory() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	log.Println("current path", ex)
	exPath := filepath.Dir(ex)
	log.Println("Executable path", exPath)
	realPath, err := filepath.EvalSymlinks(exPath)

	if err != nil {
		panic(err)
	}

	log.Println("Symlink evaluated:", realPath)
}
