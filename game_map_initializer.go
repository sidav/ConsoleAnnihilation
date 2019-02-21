package main

import (
	cw "TCellConsoleWrapper"
)

func (g *gameMap) init() {
	g.pawns = make([]*pawn, 0)
	g.factions = make([]*faction, 0)
	initMapForMission(g, 1)

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
	case '~':
		return &tile{appearance: &ccell{char: '~', r: 32, g: 32, b: 128, color: cw.BLUE}, isPassable: false, isNaval: true}
	}
	return nil
}
