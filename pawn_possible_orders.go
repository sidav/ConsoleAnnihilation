package main

func (p *pawn) canBuild() bool {
	return p.builderInfo != nil
}

func (p *pawn) canMove() bool {
	return true 
}
