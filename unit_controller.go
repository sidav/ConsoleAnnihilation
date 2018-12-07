package main

import (
	"SomeTBSGame/routines"
)

func (p *pawn) isTimeToAct() bool {
	return p.nextTurnToAct <= CURRENT_TURN
}

func (u *pawn) executeOrdersAsUnit(m *gameMap) {
	if !u.isTimeToAct() {
		return
	}

	if u.res != nil {
		u.res.resetSpendings()
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

func (u *pawn) doMoveOrder(m *gameMap) { // TODO: rewrite
	order := u.order

	ox, oy := order.x, order.y
	ux, uy := u.getCoords()

	vector := routines.CreateVectorByStartAndEndInt(ux, uy, ox, oy)
	vector.TransformIntoUnitVector()
	vx, vy := vector.GetRoundedCoords()

	u.x += vx
	u.y += vy
	u.nextTurnToAct = CURRENT_TURN + u.unitInfo.ticksForMoveOneCell

	if u.x == ox && u.y == oy {
		u.reportOrderCompletion("Arrived")
		u.order = nil
		return
	}
}

func (u *pawn) doBuildOrder(m *gameMap) { // only moves to location and/or sets the spendings. Building itself is in doAllNanolathes()
	order := u.order
	tBld := order.targetBuilding
	ux, uy := u.getCoords()
	ox, oy := tBld.getCenter()

	building_w := tBld.buildingInfo.w + 1
	building_h := tBld.buildingInfo.h + 1
	sqdistance := (ox-ux)*(ox-ux) + (oy-uy)*(oy-uy)

	if sqdistance <= building_w*building_w || sqdistance <= building_h*building_h { // is in building range
		u.res.metalSpending = u.builderInfo.builderCoeff * tBld.currentConstructionStatus.costM / tBld.currentConstructionStatus.maxConstructionAmount
		u.res.energySpending = u.builderInfo.builderCoeff * tBld.currentConstructionStatus.costE / tBld.currentConstructionStatus.maxConstructionAmount
	} else { // out of range, move to the construction site
		order.x, order.y = tBld.getCenter()
		u.doMoveOrder(m)
		log.appendMessage(u.name + " MOVES TO BUILD")
		return
	}
}

func doAllNanolathes(m *gameMap) { // does the building itself
	for _, u := range m.pawns {
		if u.order != nil && u.order.orderType == order_build {
			tBld := u.order.targetBuilding

			ux, uy := u.getCoords()
			ox, oy := tBld.getCenter()
			building_w := tBld.buildingInfo.w + 1
			building_h := tBld.buildingInfo.h + 1
			sqdistance := (ox-ux)*(ox-ux) + (oy-uy)*(oy-uy)

			if tBld.buildingInfo.hasBeenPlaced == false && (sqdistance <= building_w*building_w || sqdistance <= building_h*building_h) { // place the carcass
				u.reportOrderCompletion("Starts nanolathe")
				tBld.buildingInfo.hasBeenPlaced = true
				m.addBuilding(tBld, false)
			}

			if u.faction.economy.nanolatheAllowed && (sqdistance <= building_w*building_w || sqdistance <= building_h*building_h){
				if tBld.currentConstructionStatus == nil {
					u.reportOrderCompletion("Nanolathe interrupted")
				}
				tBld.currentConstructionStatus.currentConstructionAmount += u.builderInfo.builderCoeff
				if tBld.currentConstructionStatus.isCompleted() {
					tBld.currentConstructionStatus = nil
					u.order = nil
					u.reportOrderCompletion("Nanolathe completed")
				}
			}
		}
	}
}

func (u *pawn) reportOrderCompletion(verb string) {
	log.appendMessage(u.name + ": " + verb + ".")
}
