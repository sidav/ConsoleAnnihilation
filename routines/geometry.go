package routines

func AreCoordsInRect(x, y, rx, ry, w, h int) bool {
	return x >= rx && x < rx+w && y >= ry && y < ry+h
}

func AreCoordsInRange(fx, fy, tx, ty, r int) bool { // border including.
	// uses more wide circle (like in Bresenham's circle) than the real geometric one.
	// It is much more handy for spaces with discrete coords (cells).
	realSqDistanceAndSqRadiusDiff := (fx-tx)*(fx-tx) + (fy-ty)*(fy-ty) - r*r
	return realSqDistanceAndSqRadiusDiff < r
}

func getSqDistanceBetween(x1, y1, x2, y2 int) int {
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}

func AreTwoCellRectsOverlapping(x1, y1, w1, h1, x2, y2, w2, h2 int) bool {
	// WARNING:
	// ALL "-1"s HERE ARE BECAUSE OF WE ARE IN CELLS SPACE
	// I.E. A SINGLE CELL IS 1x1 RECTANGLE
	// SO RECTS (0, 0, 1x1) AND (1, 0, 1x1) ARE NOT OVERLAPPING IN THIS SPACE (BUT SHOULD IN EUCLIDEAN OF COURSE)
	right1 := x1 + w1 - 1
	bot1 := y1 + h1 - 1
	right2 := x2 + w2 -1
	bot2 := y2 + h2 - 1
	return !(x2 > right1 ||
		right2 < x1 ||
		y2 > bot1 ||
		bot2 < y1)
}

func AreCircleAndRectOverlapping(cx, cy, r, rx, ry, w, h int) bool {
	// downleftx, downlefty := rx+w, ry+h
	return true
}
