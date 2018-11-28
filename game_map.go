package main

const (
	mapW = 40
	mapH = 20
)

type gameMap struct {
	tileMap [mapW][mapH] *tile
	factions []*faction
	pawns []*pawn
	// buildings []*building
}

func (g *gameMap) addPawn(p *pawn) {
	g.pawns = append(g.pawns, p)
}

func (g *gameMap) addBuilding(b *pawn, asAlreadyConstructed bool) {
	if asAlreadyConstructed {
		b.currentConstructionStatus = nil
		b.buildingInfo.hasBeenPlaced = true
	}
	g.addPawn(b)
}

func (g *gameMap) getPawnAtCoordinates(x, y int) *pawn {
	for _, b := range g.pawns {
		if b.isOccupyingCoords(x, y) {
			return b
		}
	}
	return nil
}

func (g *gameMap) getUnitAtCoordinates(x, y int) *pawn {
	for _, b := range g.pawns {
		if b.isOccupyingCoords(x, y) {
			return b
		}
	}
	return nil
}

func (g *gameMap) getBuildingAtCoordinates(x, y int) *pawn {
	for _, b := range g.pawns {
		if b.isOccupyingCoords(x, y) {
			return b
		}
	}
	return nil
}

func (g *gameMap) init() {
	g.pawns = make([]*pawn, 0)
	g.factions = make([]*faction, 0)
	for i:=0; i < mapW; i++ {
		for j:=0; j < mapH; j++ {
			g.tileMap[i][j] = &tile{appearance: &ccell{char: '.', r: 64, g: 128, b: 64, color: 3}}
		}
	}

	g.factions = append(g.factions, createFaction("The Core Contingency", 0, true))
	g.addPawn(createUnit("commander", 3, 5, g.factions[0]))
	// g.addBuilding(createBuilding("metalmaker", 5, 1, g.factions[0]), true)
	// g.addUnit(createUnit("weasel", 3, 6, g.factions[0]))
	// g.addUnit(createUnit("thecan", 3, 4, g.factions[0]))
	g.addBuilding(createBuilding("corekbotlab", 5, 1, g.factions[0]), true)
	g.addBuilding(createBuilding("corevehfactory", 5, 5, g.factions[0]), true)

	g.factions = append(g.factions, createFaction("The Arm Rebellion", 1, false))
	g.addPawn(createUnit("commander", mapW-10, 5, g.factions[1]))
	g.addPawn(createUnit("ak", mapW-1, 4, g.factions[1]))
	g.addBuilding(createBuilding("armkbotlab", mapW-5, 1, g.factions[1]), true )
	g.addBuilding(createBuilding("armvehfactory", mapW-5, 5, g.factions[1]), true)
}
