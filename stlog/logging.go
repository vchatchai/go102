package stlog

import (
	"log"
	"os"
)

func Logging() {
	custLogger := log.New(os.Stdout, "custom1: ", log.Ldate|log.Ltime)
	custLogger.Println("Hello I'm customized")
	custLoggerEnh := log.New(os.Stdout, "custom2: ", log.Ldate|log.Lshortfile)
	custLoggerEnh.Println("Hello I'm customized logger 2")
}
