package main

func (p *pawn) canBuild() bool {
	return p.nanolatherInfo != nil
}

func (p *pawn) canMove() bool {
	return true // temp
}
