package main

type philosopher struct{
	eating bool
	eatingIn chan bool
	eatingOut chan bool
	timesEaten int
}

func newPhilosopher() *philosopher {
	p := philosopher{}
	p.eating = false;
	p.eatingIn = make(chan bool)
	p.eatingOut = make(chan bool)
	p.timesEaten = 0

	return &p
}

func start(p philosopher) {
	for {
		if len(p.eatingIn) > 0 {
			p.eating = <- p.eatingIn
			if p.eating {
				p.timesEaten++
			} 
		}
		p.eatingOut <- p.eating
	}
}

