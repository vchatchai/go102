package env

import (
	"log"
	"os"
)

const GOROOT = "GOROOT"

func GetENV() {
	connStr := os.Getenv(GOROOT)
	log.Println(GOROOT, connStr)
}
