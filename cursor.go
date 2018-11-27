package main

type CURSOR_MODE_ENUM uint8

const (
	CURSOR_SELECT CURSOR_MODE_ENUM = iota
	CURSOR_MOVE   CURSOR_MODE_ENUM = iota
	CURSOR_ATTACK CURSOR_MODE_ENUM = iota
	CURSOR_AMOVE  CURSOR_MODE_ENUM = iota
	CURSOR_BUILD  CURSOR_MODE_ENUM = iota // maybe merge build and repair?
)

type cursor struct {
	x, y              int
	snappedBuilding   *building
	currentCursorMode CURSOR_MODE_ENUM
	w, h int // Used for certain modes only.
}

func (c *cursor) getCoords() (int, int) {
	return c.x, c.y
}

func (c *cursor) moveByVector(x, y int) {
	c.x += x
	c.y += y
}
