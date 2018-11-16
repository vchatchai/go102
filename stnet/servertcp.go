package stnet

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
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

		channel := make(chan string, 1)

		go func() {
			for {
				msg, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					panic(err)
				}

				channel <- msg
				fmt.Printf("%+q :\n", strings.TrimSpace(msg))
				fmt.Printf("%+q :\n", strings.TrimRight(msg, "\r\n"))
				fmt.Printf("%+q :\n", "q")
				if strings.TrimSpace(msg) == "q" {
					fmt.Println("quiting...")
					close(channel)
					conn.Close()

					return
				}

			}

		}()
		go func() {
			for msg := range channel {

				_, err = io.WriteString(conn, "Received: "+string(msg))
				if err != nil {
					fmt.Println(err)
				}
			}

		}()
	}
}
