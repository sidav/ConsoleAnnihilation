package main

type cursor struct {
	x, y            int
	snappedBuilding *building

}

func (c *cursor) getCoords() (int, int) {
	return c.x, c.y
}

func (c *cursor) moveByVector(x, y int) {
	c.x += x
	c.y += y
}
