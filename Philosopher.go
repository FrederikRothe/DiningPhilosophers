package main

import (
	"fmt"
	"time"
)

type philosopher struct {
	name       string
	eating     bool
	timesEaten int
	leftFork   *fork
	rightFork  *fork
}

func NewPhilosopher(name string, leftFork *fork, rightFork *fork) *philosopher {
	p := philosopher{name: name, leftFork: leftFork, rightFork: rightFork}
	p.eating = false
	p.timesEaten = 0
	return &p
}

func Start(p *philosopher) {
	for {
		select {
		case rf := <-p.rightFork.inUse:
			p.rightFork.waiter.Lock()
			if !rf {
				lf := <-p.leftFork.inUse
				if !lf {
					p.leftFork.waiter.Lock()
					p.eating = true
					p.timesEaten++
					fmt.Printf("%s eating nam nam,  times eaten = %d\n", p.name, p.timesEaten)
					time.Sleep(300 * time.Millisecond)
					p.leftFork.used <- 1
					p.rightFork.used <- 1
					fmt.Printf("%s stopped eating\n", p.name)
					p.leftFork.waiter.Unlock()
				}
			} else {
				fmt.Println("true for some reason")
			}
			p.rightFork.waiter.Unlock()
		default:
			fmt.Println("yoo2")
		}
	}
}
