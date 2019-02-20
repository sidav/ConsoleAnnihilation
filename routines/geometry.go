package routines

func AreCoordsInRect(x, y, rx, ry, w, h int) bool {
	return x >= rx && x < rx+w && y >= ry && y < ry+h
}

func AreCoordsInRange(fx, fy, tx, ty, r int) bool {
	return (fx-tx)*(fx-tx) + (fy-ty)*(fy-ty) <= r*r
}

func GetSqDistanceBetween(x1, y1, x2, y2 int) int {
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}

func AreCircleAndRectOverlapping(cx, cy, r, rx, ry, w, h int) bool {
	// downleftx, downlefty := rx+w, ry+h
	return true
}
