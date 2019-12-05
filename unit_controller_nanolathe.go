package main

const BUILD_MAX_DISTANCE = 2

func doAllNanolathes(m *gameMap) { // does the building itself
	for _, builder := range m.pawns {
		// buildings construction
		if builder.order != nil && builder.order.orderType == order_build {
			tBld := builder.order.buildingToConstruct

			if tBld.buildingInfo.hasBeenPlaced == false && tBld.IsCloseupToCoords(builder.x, builder.y, BUILD_MAX_DISTANCE) { // place the carcass
				builder.reportOrderCompletion("Starts nanolathe")
				tBld.hitpoints = 1
				tBld.buildingInfo.hasBeenPlaced = true
				m.addBuilding(tBld, false)
			}

			if tBld.currentConstructionStatus == nil {
				builder.reportOrderCompletion("Nanolathe interrupted")
				builder.order = nil
				continue
			}
			if tBld.hitpoints <= 0 {
				builder.reportOrderCompletion("Nanolathe interrupted by hostile action")
				builder.order = nil
				continue
			}

			if builder.faction.economy.nanolatheAllowed && tBld.IsCloseupToCoords(builder.x, builder.y, BUILD_MAX_DISTANCE) {
				tBld.currentConstructionStatus.currentConstructionAmount += builder.getNanolatherInfo().builderCoeff
				tBld.hitpoints += tBld.getMaxHitpoints() / (tBld.currentConstructionStatus.maxConstructionAmount / builder.getNanolatherInfo().builderCoeff)
				if tBld.hitpoints > tBld.getMaxHitpoints() {
					tBld.hitpoints = tBld.getMaxHitpoints()
				}
				if tBld.currentConstructionStatus.isCompleted() {
					tBld.currentConstructionStatus = nil
					builder.order = nil
					builder.reportOrderCompletion("Nanolathe completed")
				}
			}
		}

		// units construction
		if builder.order != nil && builder.order.orderType == order_construct {
			currentConstructionCode := builder.order.constructingQueue[0]
			if builder.order.currentPawnUnderConstruction == nil {
				builder.order.currentPawnUnderConstruction = createSquadOfSingleMember(*currentConstructionCode, 0, 0, builder.faction, false)
			}
			uCnst := builder.order.currentPawnUnderConstruction
			ux, _ := builder.getCenter()

			if builder.faction.economy.nanolatheAllowed {
				if uCnst.currentConstructionStatus == nil {
					builder.reportOrderCompletion("WTF CONSTRUCTION STATUS IS NIL FOR " + uCnst.getName())
					continue
				}
				uCnst.currentConstructionStatus.currentConstructionAmount += builder.getNanolatherInfo().builderCoeff
				if uCnst.currentConstructionStatus.isCompleted() {
					uCnst.currentConstructionStatus = nil
					_, building_h := builder.getSize()
					uCnst.x, uCnst.y = ux, builder.y+building_h
					uCnst.order = &order{}
					uCnst.order.cloneFrom(builder.getNanolatherInfo().defaultOrderForUnitBuilt)
					m.addPawn(uCnst)
					builder.order.constructingQueue = builder.order.constructingQueue[1:]
					if builder.repeatConstructionQueue {
						builder.order.constructingQueue = append(builder.order.constructingQueue, currentConstructionCode)
					}
					builder.reportOrderCompletion("Nanolathe completed")
				}
			}
		}
	}
}
