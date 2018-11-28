package main

type pawn struct { // pawn is a building or a unit.
	unitInfo *unit
	buildingInfo *building

}

func (p *pawn) isUnit() bool {
	return p.unitInfo != nil
}

func (p *pawn) isBuilding() bool {
	return p.buildingInfo != nil
}
