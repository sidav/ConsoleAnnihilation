package main

func issueDefaultOrderToUnit(u *pawn, m *gameMap, x, y int) {
	target := m.getPawnAtCoordinates(x, y)
	if target != nil {
		if target.isBuilding() && target.currentConstructionStatus.isCompleted() == false {
			u.setOrder(&order{orderType: order_build, buildingToConstruct: target})
			log.appendMessage(u.name + ": Helps nanolathing")
			return
		}
	}
	u.setOrder(&order{orderType: order_move, x:x, y:y})
}
