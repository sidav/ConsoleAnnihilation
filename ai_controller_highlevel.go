package main

import "SomeTBSGame/routines"

func (currAi *aiData) ai_decideProduction(factory *pawn) {
	currFaction := factory.faction
	variants := &(factory.nanolatherInfo.allowedUnits)
	listOfCombatUnits := make([]*pawn, 0)
	for _, variant := range *variants {
		pawnUnderConsideration := createUnit(variant, 0, 0, currFaction, false)
		if pawnUnderConsideration.canMove() && pawnUnderConsideration.hasWeapons() {
			listOfCombatUnits = append(listOfCombatUnits, pawnUnderConsideration)
		}
	}
	if len(listOfCombatUnits) > 0 {
		pawnToProduce := listOfCombatUnits[routines.Random(len(listOfCombatUnits))]
		ai_write("producing " + pawnToProduce.name)
		if pawnToProduce != nil {
			factory.order = &order{orderType: order_construct, constructingQueue: []*pawn{pawnToProduce}}
			currAi.construction_orders_this_turn++
		}
	}
	ai_write("nothing to produce.")
}

func ai_decideConstruction(builder *pawn) {
	variants := &builder.nanolatherInfo.allowedBuildings
	ai_makeBuildOrderForBuilding(builder, (*variants)[routines.Random(len(*variants))])
	//listOfCombatUnits := make([]*pawn, 0)
	//for _, variant := range *variants {
	//	pawnUnderConsideration := createUnit(variant, 0, 0, f, false)
	//	if pawnUnderConsideration.canMove() && pawnUnderConsideration.hasWeapons() {
	//		listOfCombatUnits = append(listOfCombatUnits, pawnUnderConsideration)
	//	}
	//}
	//if len(listOfCombatUnits) > 0 {
	//	pwn := listOfCombatUnits[routines.Random(len(listOfCombatUnits))]
	//	ai_write("producing " + pwn.name)
	//	return pwn
	//}
	//ai_write("nothing to produce.")
	//return nil
}
