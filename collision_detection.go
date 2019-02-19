package main

func (p *pawn) coll_canStandAt(x, y int) bool {
	if !areCoordsValid(x, y) {
		return false
	}
	mov := p.moveInfo
	// first, let's check if the pawn can move there at all
	if !(mov.movesOnLand && CURRENT_MAP.tileMap[x][y].isPassable || mov.movesOnSea && CURRENT_MAP.tileMap[x][y].isNaval) {
		return false
	}
	//second, check if the tile is not occupied
	return CURRENT_MAP.getPawnAtCoordinates(x, y) == nil
}

func (p *pawn) coll_canMoveByVector(x, y int) bool {
	return p.coll_canStandAt(p.x + x, p.y + y)
}
