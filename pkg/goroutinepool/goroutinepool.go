package goroutinepool

import "sync"

// how to use: https://github.com/EricChiou/goroutinepool

// Pool goroutine pool struct
type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

// Add add a goroutine
func (p *Pool) Add() {
	p.queue <- 1
	p.wg.Add(1)
}

// Done one of goroutine done
func (p *Pool) Done() {
	<-p.queue
	p.wg.Done()
}

// Wait wait all goroutine finish
func (p *Pool) Wait() {
	p.wg.Wait()
}

// New creat new goroutine pool
func New(size int) *Pool {
	if size < 1 {
		size = 1
	}
	return &Pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}
