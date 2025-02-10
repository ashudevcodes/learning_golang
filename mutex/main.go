package main

import (
	"sync"
)

type post struct {
	view int
	mu   sync.Mutex
}

func (p *post) viewIncrement(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()
	p.view += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := post{view: 0}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go myPost.viewIncrement(&wg)

	}
	wg.Wait()

}
