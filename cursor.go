package main

type CURSOR_MODE_ENUM uint8

const (
	CURSOR_SELECT      CURSOR_MODE_ENUM = iota
	CURSOR_MULTISELECT CURSOR_MODE_ENUM = iota
	CURSOR_MOVE        CURSOR_MODE_ENUM = iota
	CURSOR_ATTACK      CURSOR_MODE_ENUM = iota
	CURSOR_AMOVE       CURSOR_MODE_ENUM = iota
	CURSOR_BUILD       CURSOR_MODE_ENUM = iota // maybe merge build and repair?
)

type cursor struct {
	cameraX, cameraY                     int // global space coords for upper-left (0, 0) point of screen space
	x, y                                 int
	xorig, yorig                         int // for bandbox selection
	snappedPawn                          *pawn
	currentCursorMode                    CURSOR_MODE_ENUM
	w, h, radius                         int  // Used for certain modes only.
	buildOnMetalOnly, buildOnThermalOnly bool // for build mode only
	lastSelectedIdlePawnIndex            int  // for selecting the next idle unit
}

func (c *cursor) getCoords() (int, int) {
	return c.x, c.y
}

func (c *cursor) getOnScreenCoords() (int, int) {
	return c.x - c.cameraX, c.y - c.cameraY
}

func (c *cursor) getCameraCoords() (int, int) {
	return c.cameraX, c.cameraY
}

func (c *cursor) centralizeCamera() { // changes cameraX, cameraY so that cursorX, cursorY will appear at the center of screen
	c.cameraX = c.x - VIEWPORT_W/2
	c.cameraY = c.y - VIEWPORT_H/2
}

func (c *cursor) moveByVector(x, y int) {
	c.x += x
	c.y += y
}
