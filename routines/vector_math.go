package routines

import "math"

type Vector struct {
	X, Y float64
}

//func (v *Vector) InitByStartAndEndInt(sx, sy, ex, ey int) {
//	v.X = float64(ex-sx)
//	v.Y = float64(ey-sy)
//}

func CreateVectorByStartAndEndInt(sx, sy, ex, ey int) *Vector {
	var v Vector
	v.X = float64(ex-sx)
	v.Y = float64(ey-sy)
	return &v
}

func CreateVectorByIntegers(x, y int) *Vector{
	return &Vector{float64(x), float64(y)}
}

func (v *Vector) Add(w *Vector){
	v.X += w.X
	v.Y += w.Y
}

func (v *Vector) GetRoundedCoords() (int, int) {
	return int(math.Round(v.X)), int(math.Round(v.Y))
}

func (v *Vector) Rotate(degrees int) {
	x, y := v.X, v.Y
	rads := float64(degrees)*3.14159/180.0
	v.X = x*math.Cos(rads) - y * math.Sin(rads)
	v.Y = x*math.Sin(rads) + y * math.Cos(rads)
}

func (v *Vector) GetUnitVector() *Vector {
	x, y := v.X, v.Y
	length := math.Sqrt(float64(x*x + y*y))
	newx, newy := x/length, y/length
	return &Vector{newx, newy}
}

func (v *Vector) TransformIntoUnitVector() {
	if v.X == v.Y && v.X == 0 {
		return
	}
	length := math.Sqrt(v.X*v.X + v.Y*v.Y)
	v.X /= length
	v.Y /= length
}

func CreateRandomVectorBetweenTwo(a, b *Vector) *Vector {
	const precision = 100.0
	x := float64(RandInRange(int(a.X*precision), int(b.X*precision)))/precision // strict typing sucks
	y := float64(RandInRange(int(a.Y*precision), int(b.Y*precision)))/precision
	return &Vector{x, y}
}
