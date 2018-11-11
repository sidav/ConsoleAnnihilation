package Some4xGame

type unit struct {
	x, y int
	appearance ccell
}

func (u *unit) getCoords() (int, int) {
	return u.x, u.y
}
