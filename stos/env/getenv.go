package env

import (
	"log"
	"os"
)

func GetENV(key string) {
	connStr := os.Getenv(key)
	log.Println(key, connStr)
}
