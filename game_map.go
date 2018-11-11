package Some4xGame

const (
	mapW = 20
	mapH = 10
)

type gameMap struct {
	tileMap [mapW][mapH] *tile
	units []*unit
}


