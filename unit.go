package main

type unit struct {
	faction *faction
	x, y int
	appearance ccell
	name string 
}

func (u *unit) getCoords() (int, int) {
	return u.x, u.y
}
