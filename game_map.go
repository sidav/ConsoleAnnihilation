package main

const (
	mapW = 20
	mapH = 11
)

type gameMap struct {
	tileMap [mapW][mapH] *tile
	factions []*faction
	units []*unit
	buildings []*building
}

func (g *gameMap) addUnit(u *unit) {
	g.units = append(g.units, u)
}

func (g *gameMap) addBuilding(b *building) {
	g.buildings = append(g.buildings, b)
}

func (g *gameMap) getBuildingAtCoordinates(x, y int) *building {
	for _, b := range g.buildings {
		if b.isOccupyingCoords(x, y) {
			return b
		}
	}
	return nil
}

func (g *gameMap) getUnitAtCoordinates(x, y int) *unit {
	for _, u := range g.units {
		if u.x == x && u.y == y {
			return u
		}
	}
	return nil
}

func (g *gameMap) init() {
	g.units = make([]*unit, 0)
	g.factions = make([]*faction, 0)
	for i:=0; i < mapW; i++ {
		for j:=0; j < mapH; j++ {
			g.tileMap[i][j] = &tile{appearance: &ccell{char: '.', r: 64, g: 128, b: 64, color: 3}}
		}
	}

	g.factions = append(g.factions, createFaction("The Core Contingency", 0))
	g.addUnit(createUnit("commander", 3, 5, g.factions[0]))
	g.addUnit(createUnit("weasel", 3, 6, g.factions[0]))
	g.addUnit(createUnit("thecan", 3, 4, g.factions[0]))
	g.addBuilding(createBuilding("kbotlab", 5, 1, g.factions[0]))
	g.addBuilding(createBuilding("vehfactory", 5, 5, g.factions[0]))

	//g.factions = append(g.factions, createFaction("The Arm Rebellion", 1))
	//g.addUnit(createUnit("commander", MAP_WIDTH-3, 5, g.factions[1]))
}
