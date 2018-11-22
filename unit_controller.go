package main

import (
	"SomeTBSGame/routines"
)

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

	vector := routines.CreateVectorByStartAndEndInt(ux, uy, ox, oy)
	vector.TransformIntoUnitVector()
	vx, vy := vector.GetRoundedCoords()

	u.x += vx
	u.y += vy
	u.nextTurnToAct = CURRENT_TURN + u.ticksForMoveOneCell

	if u.x == ox && u.y == oy {
		u.reportOrderCompletion("Arrived")
		u.order = nil
		return
	}
}

func (u *unit) reportOrderCompletion(verb string) {
	log.appendMessage(u.name + ": " + verb+".")
}