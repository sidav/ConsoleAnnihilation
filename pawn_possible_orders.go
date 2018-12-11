package main

func (p *pawn) canConstructBuildings() bool {
	return p.nanolatherInfo != nil && len(p.nanolatherInfo.allowedBuildings) > 0
}

func (p *pawn) canConstructUnits() bool {
	return p.nanolatherInfo != nil && len(p.nanolatherInfo.allowedUnits) > 0
}

func (p *pawn) canMove() bool {
	return true // temp
}
