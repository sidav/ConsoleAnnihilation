package routines

func GetCircle(x, y, r int) *[]point {
	points := make([]point, 0)
	if r < 0 {
		return nil
	}
	// Bresenham algorithm
	x1, y1, err := -r, 0, 2-2*r
	for {
		points = append(points, point{x-x1, y+y1})
		points = append(points, point{x-y1, y-x1})
		points = append(points, point{x+x1, y-y1})
		points = append(points, point{x+y1, y+x1})
		r = err
		if r > x1 {
			x1++
			err += x1*2 + 1
		}
		if r <= y1 {
			y1++
			err += y1*2 + 1
		}
		if x1 >= 0 {
			break
		}
	}
	return &points
}
