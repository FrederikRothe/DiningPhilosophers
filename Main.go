package main

import (
	"fmt"
	"time"
)

var queryIn chan string
var queryOut chan string

func main() {
	queryIn = make(chan string, 10)
	queryOut = make(chan string, 10)

	f1 := NewFork()
	f2 := NewFork()
	f3 := NewFork()
	f4 := NewFork()
	f5 := NewFork()

	go Fork(f1)
	go Fork(f2)
	go Fork(f3)
	go Fork(f4)
	go Fork(f5)

	p1 := NewPhilosopher("Caspar", f1, f2)
	p2 := NewPhilosopher("Rasmus", f2, f3)
	p3 := NewPhilosopher("Freddy", f3, f4)
	p4 := NewPhilosopher("The-is", f4, f5)
	p5 := NewPhilosopher("Nadija", f5, f1)

	go Start(p1)
	go Start(p2)
	go Start(p3)
	go Start(p4)
	go Start(p5)

	for {
		time.Sleep(5000 * time.Millisecond)
		queries := 10
		for i := 0; i < queries; i++ {
			queryIn <- "query"
		}
		for i := 0; i < queries; i++ {
			fmt.Println(<-queryOut)
		}
	}
}
