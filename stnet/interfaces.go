package stnet

import (
	"fmt"
	"net"
)

func Interfaces() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			panic(err)
		}
		fmt.Println(i.Name)
		for _, add := range addrs {
			if ip, ok := add.(*net.IPNet); ok {
				fmt.Printf("\t%v\n", ip)
			}
		}
	}
}
