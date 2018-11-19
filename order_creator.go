package main

func issueDefaultOrder(u *unit, m *gameMap, x, y int) {
	// target = m.getUnitAtCoordinates(x, y)
	u.setOrder(&order{order_type: order_move, x:x, y:y})
}
