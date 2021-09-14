package main

import (
	"fmt"
	"sync"
)

var number int

type fork struct {
	waiter    sync.Mutex
	timesUsed int
	used      chan int
	inUse     chan bool
	id        int
}

func NewFork() *fork {
	f := fork{}
	f.timesUsed = 0
	f.inUse = make(chan bool, 1)
	f.inUse <- false
	f.used = make(chan int, 10)
	number++
	f.id = number
	return &f
}

func Fork(f *fork) {
	go queryF(f)
	for {
		select {
		case input := <-f.used:
			f.timesUsed += input
		default:
		}

	}
}

func queryF(f *fork) {
	for {
		select {
		case <-queryIn:
			queryOut <- fmt.Sprintf("Fork number %d has been used %d times", f.id, f.timesUsed)
		}
	}
}
