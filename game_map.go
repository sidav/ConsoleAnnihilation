package main

import (
	"SomeTBSGame/routines"
	"astar/astar"
)

var (
	mapW int
	mapH int
)

type gameMap struct {
	tileMap  [][]*tile
	factions []*faction
	pawns    []*pawn
}

func (g *gameMap) addPawn(p *pawn) {
	g.pawns = append(g.pawns, p)
}

func (g *gameMap) addBuilding(b *pawn, asAlreadyConstructed bool) {
	if asAlreadyConstructed {
		b.currentConstructionStatus = nil
		b.buildingInfo.hasBeenPlaced = true
	}

	if b.nanolatherInfo != nil && len(b.nanolatherInfo.allowedUnits) > 0 { // sets default rally point for build units.
		b.nanolatherInfo.defaultOrderForUnitBuilt = &order{orderType: order_move, x: b.x + b.buildingInfo.w/2, y: b.y + b.buildingInfo.h + 1}
	}

	g.addPawn(b)
}

func (g *gameMap) removePawn(p *pawn) {
	for i := 0; i < len(g.pawns); i++ {
		if p == g.pawns[i] {
			g.pawns = append(g.pawns[:i], g.pawns[i+1:]...) // ow it's fucking... magic!
		}
	}
}

func (g *gameMap) getPawnAtCoordinates(x, y int) *pawn {
	for _, b := range g.pawns {
		if b.isOccupyingCoords(x, y) {
			return b
		}
	}
	return nil
}

//func (g *gameMap) getUnitAtCoordinates(x, y int) *pawn { // TODO: remove (as a duplicate of getPawnAtCoordinates)
//	for _, b := range g.pawns {
//		if b.isOccupyingCoords(x, y) {
//			return b
//		}
//	}
//	return nil
//}

func (g *gameMap) getPawnsInRadiusFrom(x, y, radius int) []*pawn {
	var arr []*pawn
	for _, p := range g.pawns {
		px, py := p.getCenter()
		if routines.GetSqDistanceBetween(x, y, px, py) <= radius*radius {
			arr = append(arr, p)
		}
	}
	return arr
}

func (g *gameMap) getPawnsInRect(x, y, w, h int) []*pawn {
	var arr []*pawn
	for _, p := range g.pawns {
		px, py := p.getCenter()
		if (px >= x) && (px < x+w) && (py >= y) && (py < y+h) { // TODO: pawns bigger than one cell
			arr = append(arr, p)
		}
	}
	return arr
}

func (g *gameMap) getEnemyPawnsInRadiusFrom(x, y, radius int, f *faction) []*pawn {
	var arr []*pawn
	for _, p := range g.pawns {
		px, py := p.getCenter()
		if p.faction != f && routines.GetSqDistanceBetween(x, y, px, py) <= radius*radius {
			arr = append(arr, p)
		}
	}
	return arr
}

func (g *gameMap) getBuildingAtCoordinates(x, y int) *pawn {
	for _, b := range g.pawns {
		if b.isOccupyingCoords(x, y) {
			return b
		}
	}
	return nil
}

func (g *gameMap) getNumberOfMetalDepositsInRect(x, y, w, h int) int {
	total := 0
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if areCoordsValid(x+i, y+j) {
				total += g.tileMap[x+i][y+j].metalAmount
			}
		}
	}
	return total
}

func (g *gameMap) getNumberOfThermalDepositsInRect(x, y, w, h int) int {
	total := 0
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if areCoordsValid(x+i, y+j) {
				total += g.tileMap[x+i][y+j].thermalAmount
			}
		}
	}
	return total
}

func (g *gameMap) getNumberOfMetalDepositsUnderBuilding(b *pawn) int {
	return g.getNumberOfMetalDepositsInRect(b.x, b.y, b.buildingInfo.w, b.buildingInfo.h)
}

func (g *gameMap) getNumberOfThermalDepositsUnderBuilding(b *pawn) int {
	return g.getNumberOfThermalDepositsInRect(b.x, b.y, b.buildingInfo.w, b.buildingInfo.h)
}

func (g *gameMap) canBuildingBeBuiltAt(b *pawn, cx, cy int) bool {
	b.x = cx - b.buildingInfo.w/2
	b.y = cy - b.buildingInfo.h/2
	if b.x < 0 || b.y < 0 || b.x+b.buildingInfo.w > mapW || b.y+b.buildingInfo.h > mapH {
		return false
	}
	if b.buildingInfo.canBeBuiltOnMetalOnly && g.getNumberOfMetalDepositsUnderBuilding(b) == 0 {
		return false
	}
	if b.buildingInfo.canBeBuiltOnThermalOnly && g.getNumberOfThermalDepositsUnderBuilding(b) == 0 {
		return false
	}
	if len(g.getPawnsInRect(b.x, b.y, b.buildingInfo.w, b.buildingInfo.h)) > 0 {
		return false
	}
	return true
}

func (g *gameMap) createCostMapForPathfinding() *[][]int {
	width, height := len(g.tileMap), len((g.tileMap)[0])

	costmap := make([][]int, width)
	for j := range costmap {
		costmap[j] = make([]int, height)
	}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			// TODO: optimize by iterating through pawns separately
			if !(g.tileMap[i][j].isPassable) || g.getPawnAtCoordinates(i, j) != nil {
				costmap[i][j] = -1
			}
		}
	}
	return &costmap
}

func (g *gameMap) getPathFromTo(fx, fy, tx, ty int) *astar.Cell {
	return astar.FindPath(g.createCostMapForPathfinding(), fx, fy, tx, ty, true, true)
}
