package routines

type point struct {
	X, Y int
}

func (p *point) GetCoords() (int, int) {
	return p.X, p.Y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GetLine(fromx, fromy, tox, toy int) []point {
	line := make([]point, 0)
	deltax := abs(tox - fromx)
	deltay := abs(toy - fromy)
	xmod := 1
	ymod := 1
	if tox < fromx {
		xmod = -1
	}
	if toy < fromy {
		ymod = -1
	}
	error := 0
	if deltax >= deltay {
		y := fromy
		deltaerr := deltay
		for x := fromx; x != tox+xmod; x += xmod {
			line = append(line, point{x, y})
			error += deltaerr
			if 2*error >= deltax {
				y += ymod
				error -= deltax
			}
		}
	} else {
		x := fromx
		deltaerr := deltax
		for y := fromy; y != toy+ymod; y += ymod {
			line = append(line, point{x, y})
			error += deltaerr
			if 2*error >= deltay {
				x += xmod
				error -= deltay
			}
		}
	}
	return line
}

func GetLineOver(fromx, fromy, tox, toy, length int) []point { // returns line of fixed length which does not stop at (tox, toy)
	line := make([]point, 0)
	deltax := abs(tox - fromx)
	deltay := abs(toy - fromy)
	xmod := 1
	ymod := 1
	if tox < fromx {
		xmod = -1
	}
	if toy < fromy {
		ymod = -1
	}
	error := 0
	if deltax >= deltay {
		y := fromy
		deltaerr := deltay
		for x := fromx; len(line) < length; x += xmod {
			line = append(line, point{x, y})
			error += deltaerr
			if 2*error >= deltax {
				y += ymod
				error -= deltax
			}
		}
	} else {
		x := fromx
		deltaerr := deltax
		for y := fromy; len(line) < length; y += ymod {
			line = append(line, point{x, y})
			error += deltaerr
			if 2*error >= deltay {
				x += xmod
				error -= deltay
			}
		}
	}
	return line
}
//
//func GetLastPointOfLineOver(fromx, fromy, tox, toy, length int) (int, int) { // returns last point of the line of fixed length which does not stop at (tox, toy)
//	currLength := 1
//	deltax := abs(tox - fromx)
//	deltay := abs(toy - fromy)
//	xmod := 1
//	ymod := 1
//	if tox < fromx {
//		xmod = -1
//	}
//	if toy < fromy {
//		ymod = -1
//	}
//	error := 0
//	x, y := 0, 0
//	if deltax >= deltay {
//		y = fromy
//		deltaerr := deltay
//		for x = fromx; currLength < length; x += xmod {
//			currLength++
//			error += deltaerr
//			if 2*error >= deltax {
//				y += ymod
//				error -= deltax
//			}
//		}
//	} else {
//		x = fromx
//		deltaerr := deltax
//		for y = fromy; currLength < length; y += ymod {
//			currLength++
//			error += deltaerr
//			if 2*error >= deltay {
//				x += xmod
//				error -= deltay
//			}
//		}
//	}
//	return x, y
//}
