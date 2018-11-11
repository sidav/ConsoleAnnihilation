package SomeTBSGame

type unit struct {
	faction int
	x, y int
	appearance ccell
	name string 
}

func (u *unit) getCoords() (int, int) {
	return u.x, u.y
}
