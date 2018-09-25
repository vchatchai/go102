package struntime

import (
	"log"
	"runtime"
)

const info = `
Application %s starting.
The binary was build by GO: %s`

func Version() {
	log.Printf(info, "Example", runtime.Version())
}
