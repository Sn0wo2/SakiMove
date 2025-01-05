package main

import (
	"fmt"
	"golang.org/x/exp/rand"
	"time"
)

type Saki struct {
	Name string
	X    int
	Y    int
	Z    int
}

type Cucumber struct {
	Water int
}

func (c *Cucumber) eat() {
	c = nil
}

func (s *Saki) move() Saki {
	s.X = rand.Intn(10)
	s.Y = rand.Intn(1)
	s.Z = rand.Intn(10)
	return *s
}

func initSaki() *Saki {
	return &Saki{Name: "丰川祥子", X: 0, Y: 100, Z: 0}
}

func muzimi(cucumber *Cucumber) (brokenSaki Saki) {
	cucumber.eat()
	return initSaki().move()
}

func main() {
	rand.Seed(uint64(time.Now().UnixNano()))
	fmt.Println(muzimi(&Cucumber{Water: 100}))
}
