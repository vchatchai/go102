package stsync

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type SyncWrite struct {
	m      sync.Mutex
	Writer io.Writer
}

func (w *SyncWrite) Write(b []byte) (n int, err error) {
	w.m.Lock()
	defer w.m.Unlock()
	return w.Writer.Write(b)
}

var data = []string{
	"Hello!",
	"Ola!",
	"Ahoj",
}

func Syncwrite() {
	f, err := os.Create("simple.file")
	if err != nil {
		panic(err)
	}

	wr := &SyncWrite{sync.Mutex{}, f}
	wg := sync.WaitGroup{}
	for _, val := range data {
		wg.Add(1)
		go func(greetings string) {
			fmt.Fprintln(wr, greetings)
			wg.Done()
		}(val)
	}
	wg.Wait()

}
