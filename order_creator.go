package main

func issueDefaultOrderToUnit(u *pawn, m *gameMap, x, y int) {
	// target = m.getUnitAtCoordinates(x, y)
	//tBld := m.getBuildingAtCoordinates(x, y)
	//if tBld != nil {
	//	u.setOrder(&order{orderType: order_build, targetBuilding:tBld})
	//	log.appendMessage(u.name + ": Helps nanolathing")
	//	return
	//}
	u.setOrder(&order{orderType: order_move, x:x, y:y})
}
