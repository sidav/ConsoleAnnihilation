package routines

import (
	"time"
)

const (
	a = 513
	c = 313
	m = 262147
)

var (
	x int
)

func Randomize() int {
	x = int(time.Duration(time.Now().UnixNano())/time.Millisecond) % m
	return x
}

func SetSeed(val int) {
	x = val
}

func Random(modulo int) int {
	x = (x*a + c) % m
	if modulo != 0 {
		return x % modulo
	} else {
		return x
	}
}

func RollDice(dnum, dval, dmod int) int {
	var result int
	for i := 0; i < dnum; i++ {
		result += Random(dval) + 1
	}
	return result + dmod
}

func RandomUnitVectorInt() (int, int) {
	var vx, vy int
	for vx == 0 && vy == 0 {
		vx, vy = Random(3)-1, Random(3)-1
	}
	return vx, vy
}

func RandInRange(from, to int) int { //should be inclusive
	if to < from {
		t := from
		from = to
		to = t
	}
	if from == to {
		return from
	}
	return Random(to-from+1) + from // TODO: replace routines.random usage with package own implementation
}

func RandomPercent() int {
	return Random(100)
}
