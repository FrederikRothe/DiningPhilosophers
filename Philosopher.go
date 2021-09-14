package main

type philosopher struct{
	eating bool
	timesEaten int
	leftFork *fork
	rightFork *fork
}

func NewPhilosopher(leftFork *fork, rightFork *fork) *philosopher {
	p := philosopher{leftFork: leftFork, rightFork: rightFork}
	p.eating = false
	p.timesEaten = 0
	return &p
}

func Start(p philosopher) {
	for {
		var rf, lf bool
		select {
		case rf <- rightFork.output:
			rightFork.waiter.Lock()
			select {
			case lf <- leftFork.output:
				leftFork.waiter.Lock()
				p.eating = true
				p.timesEaten++
				fmt.Println("Eating nam nam")
				time.Sleep(300 * time.Millisecond)
				leftFork.input <- 1
				rightFork.input <- 1
				leftFork.waiter.Unlock()
			}
			rightFork.waiter.Unlock()
		}	
	}
}

