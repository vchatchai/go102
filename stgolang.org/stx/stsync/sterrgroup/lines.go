package sterrgroup

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"strings"

	"golang.org/x/sync/errgroup"
)

const data = `line one 
line two with more words
error: this iserrorneous line
line three
`

func Lines() {
	log.Printf("Application %s starting.", "Error Detection")
	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanLines)

	g, _ := errgroup.WithContext(context.Background())
	for scanner.Scan() {
		row := scanner.Text()
		g.Go(func() error {
			return func(s string) error {
				if strings.Contains(s, "error:") {
					return fmt.Errorf(s)
				}
				return nil
			}(row)
		})

		fmt.Println(row)
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Error while waiting:", err.Error())
	}

}
