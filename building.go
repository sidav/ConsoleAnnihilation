package main

type building struct {
	name                      string
	x, y                      int // TOP LEFT CORNER!
	w, h                      int
	appearance                *buildingAppearance
	currentConstructionStatus *constructionInformation
	faction                   *faction
	hasBeenPlaced             bool
}

func (b *building) getCoords() (int, int) {
	return b.x, b.y
}

func (b *building) isOccupyingCoords(x, y int) bool {
	return areCoordsInRect(x, y, b.x, b.y, b.w, b.h)
}

func (b *building) getCenter() (int, int) {
	return b.x + b.w / 2, b.y + b.h / 2
}
