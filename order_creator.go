package main

func issueDefaultOrder(u *unit, m *gameMap, x, y int) {
	// target = m.getUnitAtCoordinates(x, y)
	u.setOrder(&order{orderType: order_move, x:x, y:y})
}
