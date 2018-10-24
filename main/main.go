package main

import (
	"log"

	"github.com/rhymond/go-money"
	"github.com/vchatchai/go102/stos"
)

func main() {
	// stos.Arguments()
	// stflag.Flag()
	// stsignal.Signal()

	// stexec.Exec()

	// stexec.Asynchronous()
	// stexec.ReadBufferChildProcess()
	// stexec.WriteChileProcess()
	// stsignal.Shutdown()
	// Cal()
	// stfmt.Scanner()
	stos.Reader()
}

func Cal() {
	baht := money.New(9900, "THB")
	log.Println(baht.Display())
	// baht2 := money.New(10, "THB")
	parties, err := baht.Allocate(3, 1)
	if err != nil {
		log.Fatal(err)
	}

	for _, party := range parties {
		log.Println(party.Display(), party.Amount())
	}

	mul := baht.Multiply(7)
	log.Println(mul.Display())

	div := baht.Divide(7)
	log.Println(div.Display())

	vat := Vat(baht)
	log.Println(baht.Display(), vat.Display())
	var count int64

	for count = 0; count <= 100; count++ {

		baht = money.New(count, baht.Currency().Code)

		vat := Vat(baht)
		log.Println(baht.Display(), vat.Display())
	}

}

func Vat(amount *money.Money) *money.Money {
	vat := amount.Multiply(7)
	vat = vat.Divide(100)
	return vat
}
