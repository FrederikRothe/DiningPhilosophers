package main

import (
	"sync"
)

type fork struct {
	waiter    sync.Mutex
	timesUsed int
}

func Fork(f fork) {
	output := make(chan bool, 1)
	input := make(chan int, 1)

	for {
		f.timesUsed += <-input
		output <- false
	}
}
