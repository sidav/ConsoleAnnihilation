package main

const (
	mapW = 20
	mapH = 10
)

type gameMap struct {
	tileMap [mapW][mapH] *tile
	units []*unit
}

func (g *gameMap) addUnit(u *unit) {
	g.units = append(g.units, u)
}

func (g *gameMap) init() {
	g.units = make([]*unit, 0)
	for i:=0; i < mapW; i++ {
		for j:=0; j < mapH; j++ {
			g.tileMap[i][j] = &tile{appearance: &ccell{char: '.', r: 64, g: 128, b: 64}}
		}
	}

	g.addUnit(createUnit("commander", 0, 0, 0))
}
