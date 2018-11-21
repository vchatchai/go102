package stsync

import (
	"fmt"
	"sync"
)

var namesInterfaces = []interface{}{"Alan", "Joe", "Jack", "Ben", "Ellen", "Lisa", "Carl", "Steve", "Anton", "Yo"}

type Source struct {
	m    *sync.Mutex
	o    *sync.Once
	data []interface{}
}

func (s *Source) Pop() (interface{}, error) {
	s.m.Lock()
	defer s.m.Unlock()
	s.o.Do(func() {
		s.data = namesInterfaces
		fmt.Println("Data has been loaded.")
	})
	if len(s.data) > 0 {
		res := s.data[0]
		s.data = s.data[1:]
		return res, nil
	}
	return nil, fmt.Errorf("No data available")

}

func Once() {
	s := &Source{&sync.Mutex{}, &sync.Once{}, nil}
	wg := &sync.WaitGroup{}
	wg.Add(len(namesInterfaces))

	for i := 0; i < len(namesInterfaces); i++ {
		go func(idx int) {
			if val, err := s.Pop(); err == nil {
				fmt.Printf("Pop %d returned: %s\n", idx, val)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
