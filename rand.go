package main

import (
	"math/rand"
	"time"
)

type Rand struct {
	rndSeed  rand.Source
	rndKlass *rand.Rand
}

func newRand() Rand {
	r := Rand{}
	r.rndSeed = rand.NewSource(time.Now().UnixNano())
	r.rndKlass = rand.New(r.rndSeed)
	return r
}

func (r Rand) rnd(n int) int {
	return r.rndKlass.Intn(n)
}

// returns random number from 0 to n-1 with new seed everytime
func rndNum(n int) int {
	return newRand().rndKlass.Intn(n)
}
