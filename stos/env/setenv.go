package env

import (
	"log"
	"os"
)

func SetENV(key, value string) {
	os.Setenv(key, value)

	val, status := os.LookupEnv(key)
	if !status {
		log.Println("Not has value")
	}

	log.Println("This value is", val)

}
