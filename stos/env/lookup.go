package env

import (
	"fmt"
	"log"
	"os"
)

func Lookup(key string) {
	value, err := os.LookupEnv(key)
	if !err {
		log.Printf("The env variable %s is not set %v.\n", key, err)
		return
	}
	fmt.Println(key, value)
}
