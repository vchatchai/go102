package stsync

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type SearrchSrc struct {
	ID    string
	Delay int
}

func (s *SearrchSrc) Search(ctx context.Context) <-chan string {
	out := make(chan string)
	go func() {
		time.Sleep(time.Duration(s.Delay) * time.Second)
		select {
		case out <- "Result " + s.ID:
		case <-ctx.Done():
			fmt.Println("Search received Done()")
		}
		close(out)
		fmt.Println("Search finished for ID:", s.ID)
	}()

	return out
}

func First() {
	ctx, cancel := context.WithCancel(context.Background())

	src1 := &SearrchSrc{"1", 2}
	src2 := &SearrchSrc{"2", 6}

	r1 := src1.Search(ctx)
	r2 := src2.Search(ctx)

	out := merge(ctx, r1, r2)

	for firstResult := range out {
		cancel()
		fmt.Println("First result is: ", firstResult)
	}
}

func merge(ctx context.Context, results ...<-chan string) <-chan string {
	wg := sync.WaitGroup{}
	out := make(chan string)
	output := func(c <-chan string) {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("Received ctx.Done()")
		case res := <-c:
			out <- res
		}
	}

	wg.Add(len(results))
	for _, c := range results {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
