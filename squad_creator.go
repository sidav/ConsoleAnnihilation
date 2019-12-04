package main 

// func createEmptySquad(x, y int, f *faction, alreadyConstructed bool) *pawn {
// 	sqpawn := &pawn{x: x, y: y, squadInfo: &squad{}}
// 	return sqpawn
// }

func createSquadOfSingleMember(code string, x, y int, f *faction, alreadyConstructed bool) *pawn {
	sqpawn := &pawn{x: x, y: y, 
		faction: f,
		squadInfo: &squad{}}
	return sqpawn
}
