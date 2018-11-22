package main

import "fmt"

type unit struct {
	faction *faction
	x, y int
	appearance ccell
	name string
	order *order
	nextTurnToAct int
	ticksForMoveOneCell int
}

func (u *unit) getCoords() (int, int) {
	return u.x, u.y
}

func (u *unit) setOrder(o *order) {
	u.order = o
	log.appendMessage(fmt.Sprintf("Order for %d, %d received!", o.x, o.y))
}
