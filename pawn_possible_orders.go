package main

func (p *pawn) canConstructBuildings() bool {
	return p.getNanolatherInfo() != nil && len(p.getNanolatherInfo().allowedBuildings) > 0
}

func (p *pawn) canConstructUnits() bool {
	return p.getNanolatherInfo() != nil && len(p.getNanolatherInfo().allowedUnits) > 0
}

func (p *pawn) canMove() bool {
	return p.getMovementInfo() != nil
}
