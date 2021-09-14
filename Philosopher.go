package main

import (
	"fmt"
	"math/rand"
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
	go queryP(p)
	for {
		select {
		case rf := <-p.rightFork.inUse:
			if rf {
			} else {
				select {
				case lf := <-p.leftFork.inUse:
					if lf {
					} else {
						p.eating = true
						p.timesEaten++
						//fmt.Printf("%s eating nam nam,  times eaten = %d\n", p.name, p.timesEaten)
						wait := time.Duration(rand.Intn(500)+300) * time.Millisecond
						time.Sleep(wait)
						//fmt.Printf("%s stopped eating\n", p.name)
						p.leftFork.used <- 1
						p.rightFork.used <- 1
						p.leftFork.inUse <- false
					}
				default:
				}
			}
			p.rightFork.inUse <- false
		default:
		}
	}
}

func queryP(p *philosopher) {
	for {
		select {
		case <-queryIn:
			queryOut <- fmt.Sprintf("%s has eaten %d times", p.name, p.timesEaten)
		}
	}
}

// func StartOld(p *philosopher) {
// 	for {
// 		select {
// 		case rf := <-p.rightFork.inUse:
// 			p.rightFork.waiter.Lock()
// 			if !rf {
// 				select {
// 				case lf := <-p.leftFork.inUse:
// 					p.leftFork.waiter.Lock()
// 					if !lf {
// 						p.eating = true
// 						p.timesEaten++
// 						fmt.Printf("%s eating nam nam,  times eaten = %d\n", p.name, p.timesEaten)
// 						wait := time.Duration(rand.Intn(500)+300) * time.Millisecond
// 						time.Sleep(wait)
// 						p.leftFork.used <- 1
// 						p.rightFork.used <- 1
// 						fmt.Printf("%s stopped eating\n", p.name)
// 						p.leftFork.waiter.Unlock()
// 					}
// 				}
// 			} else {
// 				fmt.Println("true for some reason")
// 			}
// 			p.rightFork.waiter.Unlock()
// 		default:
// 			fmt.Println("yoo2")
// 		}
// 	}
// }
