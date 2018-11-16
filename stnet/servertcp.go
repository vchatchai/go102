package stnet

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func ServerTCP() {
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("Waiting for client...")
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			for {
				msg, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					panic(err)
				}

				if msg == "q" {
					conn.Close()
				}

				_, err = io.WriteString(conn, "Received: "+string(msg))
				if err != nil {
					fmt.Println(err)
				}
			}

		}()

	}
}
