package main

func (u *unit) isTimeToAct() bool {
	return u.nextTurnToAct <= CURRENT_TURN
}

func (u *unit) executeOrders(m *gameMap) {
	if !u.isTimeToAct() {
		return
	}

	order := u.order
	if order == nil {
		return
	}

	switch order.order_type {
	case order_move:
		u.doMoveOrder(m)
	}
	// move

}

func (u *unit) doMoveOrder(m *gameMap) { // TODO: rewrite
	order := u.order
	ox, oy := order.x, order.y
	ux, uy := u.getCoords()
	vx, vy := ox-ux, oy-uy
	if vx != 0 {
		vx = vx/abs(vx)
	}
	if vy != 0 {
		vy = vy/abs(vy)
	}
	u.x += vx
	u.y += vy
	u.nextTurnToAct += 10
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
