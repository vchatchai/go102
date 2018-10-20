package stmessage

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const num = 100000.5678

func Localized() {
	p := message.NewPrinter(language.English)
	p.Printf(" %.2f \n", num)

	p = message.NewPrinter(language.German)
	p.Printf(" %.2f \n", num)

	p = message.NewPrinter(language.Thai)
	p.Printf(" %.2f \n", num)
}
