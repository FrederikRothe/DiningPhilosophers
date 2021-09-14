package main

import (
	"sync"
)

type fork struct {
	waiter    sync.Mutex
	timesUsed int
	used      chan int
	inUse     chan bool
}

func NewFork() *fork {
	f := fork{}
	f.timesUsed = 0
	f.inUse = make(chan bool, 1)
	f.inUse <- false
	f.used = make(chan int, 10)
	return &f
}

func Fork(f *fork) {
	for {
		select {
		case input := <-f.used:
			f.timesUsed += input
		default:
		}

	}
}
