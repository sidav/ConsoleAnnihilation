package main

import (
	cw "TCellConsoleWrapper"
)

func (g *gameMap) init() {
	g.pawns = make([]*pawn, 0)
	g.factions = make([]*faction, 0)
	g.initTileMap(MIS1_MAP)

	g.factions = append(g.factions, createFaction("The Core Corporation", 0, true))
	g.addPawn(createUnit("protocommander", 3, 9, g.factions[0], true))
	g.factions[0].cursor.x = 3
	g.factions[0].cursor.y = 9

	g.addPawn(createUnit("flash", 4, 5, g.factions[0], true))
	g.addPawn(createUnit("flash", 4, 6, g.factions[0], true))
	g.addPawn(createUnit("flash", 5, 5, g.factions[0], true))
	g.addPawn(createUnit("flash", 5, 6, g.factions[0], true))
	g.addPawn(createUnit("weasel", 7, 7, g.factions[0], true))


	g.factions = append(g.factions, createFaction("The rogue Arm AI", 1, false))
	// g.addPawn(createUnit("armcommander", mapW-10, 5, g.factions[1], true))
	g.addBuilding(createBuilding("armhq", mapW-5, 9, g.factions[1]), true)
	g.addBuilding(createBuilding("mstorage", 20, 10, g.factions[1]), true)
	g.addBuilding(createBuilding("lturret", mapW-10, 1, g.factions[1]), true)
	g.addBuilding(createBuilding("lturret", mapW-10, 4, g.factions[1]), true)
	g.addBuilding(createBuilding("guardian", mapW-7, 3, g.factions[1]), true)
	g.addBuilding(createBuilding("lturret", mapW-10, 8, g.factions[1]), true)
	g.addBuilding(createBuilding("lturret", mapW-10, 12, g.factions[1]), true)
	g.addBuilding(createBuilding("lturret", mapW-10, 16, g.factions[1]), true)
	g.addBuilding(createBuilding("guardian", mapW-7, 14, g.factions[1]), true)
	g.addBuilding(createBuilding("lturret", mapW-10, 19, g.factions[1]), true)

	for _, f := range g.factions {
		f.recalculateFactionEconomy(g)
	}
}

func (g *gameMap) initTileMap(strmap *[]string) {
	mapH = len(*strmap)
	mapW = len((*strmap)[0])
	g.tileMap = make([][]*tile, mapW)
	for i := range g.tileMap {
		g.tileMap[i] = make([]*tile, mapH)
	}

	for y, str := range *strmap{
		for x, chr := range str {
			g.tileMap[x][y] = mapinit_getTileByChar(chr)
		}
	}
}

func mapinit_getTileByChar(char rune) *tile {
	switch char {
	case '.':
		return &tile{appearance: &ccell{char: '.', r: 64, g: 128, b: 64, color: cw.DARK_YELLOW}, isPassable: true}
	case ';':
		return &tile{appearance: &ccell{char: ';', r: 64, g: 64, b: 128, color: cw.DARK_GRAY}, metalAmount: 1, isPassable: true}
	case '$':
		return &tile{appearance: &ccell{char: '$', r: 64, g: 64, b: 128, color: cw.DARK_GRAY}, thermalAmount: 1, isPassable: true}
	case '^':
		return &tile{appearance: &ccell{char: '^', r: 64, g: 64, b: 128, color: cw.BEIGE}, isPassable: false}
	}
	return nil
}
