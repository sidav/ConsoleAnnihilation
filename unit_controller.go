package main

import (
	"SomeTBSGame/routines"
	"TCellConsoleWrapper"
	"time"
)

func (p *pawn) isTimeToAct() bool {
	return p.nextTurnToAct <= CURRENT_TURN
}

func (u *pawn) executeOrders(m *gameMap) {
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
	case order_attack:
		u.doAttackOrder()
	case order_build:
		u.doBuildOrder(m)
	case order_construct:
		u.doConstructOrder(m)
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

	if u.coll_canMoveByVector(vx, vy) {

		u.x += vx
		u.y += vy

		u.nextTurnToAct = CURRENT_TURN + u.moveInfo.ticksForMoveSingleCell

		if u.x == ox && u.y == oy {
			u.reportOrderCompletion("Arrived")
			u.order = nil
			return
		}
	}
}

func (p *pawn) doAttackOrder() { // Only moves the unit to a firing position. The firing itself is in openFireIfPossible()
	order := p.order

	ux, uy := p.getCoords()
	if order.targetPawn.hitpoints <= 0 {
		p.reportOrderCompletion("target destroyed. Now standing by")
		p.order = nil
	}
	targetX, targetY := order.targetPawn.getCenter()

	if getSqDistanceBetween(ux, uy, targetX, targetY) > p.getMaxRadiusToFire()*p.getMaxRadiusToFire() {
		order.x = targetX
		order.y = targetY
		p.doMoveOrder(CURRENT_MAP)
		return
	}
}

func (p *pawn) openFireIfPossible() { // does the firing, does NOT necessary mean execution of attack order (but can be)
	if p.currentConstructionStatus != nil || !p.hasWeapons() || p.order != nil && p.order.orderType == order_build {
		return
	}
	var pawnInOrder *pawn
	if p.order != nil && p.order.targetPawn != nil {
		pawnInOrder = p.order.targetPawn
	}
	for _, wpn := range p.weapons {
		if (wpn.canBeFiredOnMove && wpn.nextTurnToFire > CURRENT_TURN) || (!wpn.canBeFiredOnMove && !p.isTimeToAct()) {
			// log.appendMessage(fmt.Sprintf("Skipping fire: TtA:%b CBFoM:%b TRN: %b", p.isTimeToAct() ,wpn.canBeFiredOnMove, wpn.nextTurnToFire > CURRENT_TURN))
			continue
		}
		var target *pawn
		radius := wpn.attackRadius
		if pawnInOrder != nil && areCoordsInRange(p.x, p.y, pawnInOrder.x, pawnInOrder.y, radius) {
			target = pawnInOrder
		} else {
			potential_targets := CURRENT_MAP.getEnemyPawnsInRadiusFrom(p.x, p.y, radius, p.faction)
			if len(potential_targets) > 0 {
				target = potential_targets[0]
			}
		}
		if target != nil {
			if wpn.canBeFiredOnMove {
				wpn.nextTurnToFire = CURRENT_TURN + wpn.attackDelay
			} else {
				p.nextTurnToAct = CURRENT_TURN + wpn.attackDelay
			}
			// draw the pew pew laser TODO: move this crap somewhere already 
			if areGlobalCoordsOnScreenForFaction(p.x, p.y, CURRENT_FACTION_SEEING_THE_SCREEN) || areGlobalCoordsOnScreenForFaction(target.x, target.y, CURRENT_FACTION_SEEING_THE_SCREEN) {
				tcell_wrapper.SetFgColor(tcell_wrapper.RED)
				renderLine(p.x, p.y, target.x, target.y, true, CURRENT_FACTION_SEEING_THE_SCREEN.cursor.x-VIEWPORT_W/2, CURRENT_FACTION_SEEING_THE_SCREEN.cursor.y-VIEWPORT_H/2)
				// tcell_wrapper.Flush_console()
				time.Sleep(250 * time.Millisecond)
			}
			dealDamageToTarget(p, wpn, target)
		}
	}
}

func (u *pawn) doBuildOrder(m *gameMap) { // only moves to location and/or sets the spendings. Building itself is in doAllNanolathes()
	order := u.order
	tBld := order.buildingToConstruct
	ux, uy := u.getCoords()
	ox, oy := tBld.getCenter()

	building_w := tBld.buildingInfo.w + 1
	building_h := tBld.buildingInfo.h + 1
	sqdistance := getSqDistanceBetween(ox, oy, ux, uy) //(ox-ux)*(ox-ux) + (oy-uy)*(oy-uy)

	if tBld == nil {
		log.appendMessage(u.name + " NIL BUILD")
		return
	}

	if tBld.currentConstructionStatus == nil {
		u.reportOrderCompletion("Construction is finished by another unit")
		u.order = nil
		return 
	}

	if sqdistance <= building_w*building_w || sqdistance <= building_h*building_h { // is in building range
		u.res.metalSpending = u.nanolatherInfo.builderCoeff * tBld.currentConstructionStatus.costM / tBld.currentConstructionStatus.maxConstructionAmount
		u.res.energySpending = u.nanolatherInfo.builderCoeff * tBld.currentConstructionStatus.costE / tBld.currentConstructionStatus.maxConstructionAmount
	} else { // out of range, move to the construction site
		order.x, order.y = tBld.getCenter()
		u.doMoveOrder(m)
		log.appendMessage(u.name + " MOVES TO BUILD")
		return
	}
}

func (p *pawn) doConstructOrder(m *gameMap) {
	order := p.order

	if len(p.order.constructingQueue) == 0 {
		p.reportOrderCompletion("Construction queue finished")
		p.order = nil
		return
	}

	uCnst := order.constructingQueue[0]

	p.res.metalSpending = p.nanolatherInfo.builderCoeff * uCnst.currentConstructionStatus.costM / uCnst.currentConstructionStatus.maxConstructionAmount
	p.res.energySpending = p.nanolatherInfo.builderCoeff * uCnst.currentConstructionStatus.costE / uCnst.currentConstructionStatus.maxConstructionAmount
}

func doAllNanolathes(m *gameMap) { // does the building itself
	for _, u := range m.pawns {
		// buildings construction
		if u.order != nil && u.order.orderType == order_build {
			tBld := u.order.buildingToConstruct

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

			if u.faction.economy.nanolatheAllowed && (sqdistance <= building_w*building_w || sqdistance <= building_h*building_h) {
				if tBld.currentConstructionStatus == nil {
					u.reportOrderCompletion("Nanolathe interrupted")
					u.order = nil
					continue
				}
				if tBld.hitpoints <= 0 {
					u.reportOrderCompletion("Nanolathe interrupted by hostile action")
					u.order = nil
					continue
				}
				tBld.currentConstructionStatus.currentConstructionAmount += u.nanolatherInfo.builderCoeff
				if tBld.currentConstructionStatus.isCompleted() {
					tBld.currentConstructionStatus = nil
					u.order = nil
					u.reportOrderCompletion("Nanolathe completed")
				}
			}
		}

		// units construction
		if u.order != nil && u.order.orderType == order_construct {
			uCnst := u.order.constructingQueue[0]

			ux, _ := u.getCenter()

			if u.faction.economy.nanolatheAllowed {
				if uCnst.currentConstructionStatus == nil {
					u.reportOrderCompletion("WTF CONSTRUCTION STATUS IS NIL FOR " + uCnst.name)
					continue
				}
				uCnst.currentConstructionStatus.currentConstructionAmount += u.nanolatherInfo.builderCoeff
				if uCnst.currentConstructionStatus.isCompleted() {
					uCnst.currentConstructionStatus = nil
					uCnst.x, uCnst.y = ux, u.y+u.buildingInfo.h
					uCnst.order = &order{}
					uCnst.order.cloneFrom(u.nanolatherInfo.defaultOrderForUnitBuilt)
					m.addPawn(uCnst)
					u.order.constructingQueue = u.order.constructingQueue[1:]
					u.reportOrderCompletion("Nanolathe completed")
				}
			}
		}
	}
}

func (u *pawn) reportOrderCompletion(verb string) {
	log.appendMessage(u.name + ": " + verb + ".")
}
