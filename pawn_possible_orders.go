package main

func (p *pawn) canConstructBuildings() bool {
	return len(p.nanolatherInfo.allowedBuildings) > 0
}

func (p *pawn) canConstructUnits() bool {
	return len(p.nanolatherInfo.allowedUnits) > 0
}

func (p *pawn) canMove() bool {
	return true // temp
}
