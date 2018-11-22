package main

type building struct {
	name string
	x, y int // TOP LEFT CORNER!
	w, h int
	appearance *buildingAppearance
	faction *faction
}

func (b *building) getCoords() (int, int) {
	return b.x, b.y
}
