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

	switch order.orderType {
	case order_move:
		u.doMoveOrder(m)
	case order_build:
		u.doBuildOrder(m)
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

func (u *unit) doBuildOrder(m* gameMap) {
	order := u.order
	tBld := order.targetBuilding
	ux, uy := u.getCoords()
	ox, oy := tBld.getCenter()

	building_w := tBld.w + 1
	building_h := tBld.h + 1
	sqdistance := (ox-ux)*(ox-ux) + (oy-uy)*(oy-uy)

	if sqdistance <= building_w * building_w || sqdistance <= building_h * building_h {
		log.appendMessage(u.name + " STARTS BUILDING")
		tBld.currentConstructionStatus = &constructionInformation{0, 100}
		m.addBuilding(tBld)
		u.order = nil
	} else {
		u.doMoveOrder(m)
		log.appendMessage(u.name + " MOVES TO BUILD")
		return
	}
}

func (u *unit) reportOrderCompletion(verb string) {
	log.appendMessage(u.name + ": " + verb+".")
}