package stnet

import (
	"fmt"
	"net"
)

func Lookup() {
	addrs, err := net.LookupAddr("127.0.0.1")
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		fmt.Println("ADDR:", addr)
	}

	ips, err := net.LookupIP("www.google.co.th")
	if err != nil {
		panic(err)
	}

	for _, ip := range ips {
		fmt.Println("IP:", ip.String())
	}
}
