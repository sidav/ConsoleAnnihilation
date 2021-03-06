package main

const BUILD_MAX_DISTANCE = 2

func doAllNanolathes(m *gameMap) { // does the building itself
	for _, u := range m.pawns {
		// buildings construction
		if u.order != nil && u.order.orderType == order_build {
			tBld := u.order.buildingToConstruct

			if tBld.buildingInfo.hasBeenPlaced == false && tBld.IsCloseupToCoords(u.x, u.y, BUILD_MAX_DISTANCE)  { // place the carcass
				u.reportOrderCompletion("Starts nanolathe")
				tBld.hitpoints = 1
				tBld.buildingInfo.hasBeenPlaced = true
				m.addBuilding(tBld, false)
			}

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

			if u.faction.economy.nanolatheAllowed && tBld.IsCloseupToCoords(u.x, u.y, BUILD_MAX_DISTANCE) {
				tBld.currentConstructionStatus.currentConstructionAmount += u.nanolatherInfo.builderCoeff
				tBld.hitpoints += tBld.maxHitpoints / (tBld.currentConstructionStatus.maxConstructionAmount / u.nanolatherInfo.builderCoeff)
				if tBld.hitpoints > tBld.maxHitpoints {
					tBld.hitpoints = tBld.maxHitpoints
				}
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
					if u.repeatConstructionQueue {
						u.order.constructingQueue = append(u.order.constructingQueue, createUnit(uCnst.codename, 0, 0, u.faction, false))
					}
					u.reportOrderCompletion("Nanolathe completed")
				}
			}
		}
	}
}