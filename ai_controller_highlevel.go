package main

import "SomeTBSGame/routines"

func (currAi *aiData) ai_decideProduction(factory *pawn) {
	currFaction := factory.faction
	variants := &(factory.nanolatherInfo.allowedUnits)
	listOfCombatUnits := make([]*pawn, 0)
	listOfEngineerUnits := make([]*pawn, 0)
	for _, variant := range *variants {
		pawnUnderConsideration := createUnit(variant, 0, 0, currFaction, false)
		if pawnUnderConsideration.canMove() && pawnUnderConsideration.hasWeapons() {
			listOfCombatUnits = append(listOfCombatUnits, pawnUnderConsideration)
		}
		if pawnUnderConsideration.canConstructBuildings() {
			listOfEngineerUnits = append(listOfEngineerUnits, pawnUnderConsideration)
		}
	}
	var pawnToProduce *pawn
	if len(listOfEngineerUnits) > 0 && currAi.shouldProduceEngineers() {
		pawnToProduce = listOfEngineerUnits[routines.Random(len(listOfEngineerUnits))]
		ai_write("producing engineer " + pawnToProduce.name)
	} else {
		if len(listOfCombatUnits) > 0 {
			pawnToProduce = listOfCombatUnits[routines.Random(len(listOfCombatUnits))]
			ai_write("producing " + pawnToProduce.name)
		}
	}

	if pawnToProduce != nil {
		factory.order = &order{orderType: order_construct, constructingQueue: []*pawn{pawnToProduce}}
		currAi.construction_orders_this_turn++
	} else {
		ai_write("nothing to produce.")
	}
}

func (ai *aiData) ai_decideConstruction(builder *pawn) {
	variants := builder.nanolatherInfo.allowedBuildings
	step := ai.getCurrentOrderStep()
	final_build_variant := ""
	for _, variant := range variants {
		if variant == step.buildCode || variant == step.buildCodeAlt {
			final_build_variant = variant
			ai.orderStepSatisfied()
		}
	}
	if final_build_variant == "" {
		final_build_variant = variants[routines.Random(len(variants))]
	}
	ai_makeBuildOrderForBuilding(builder, final_build_variant)
}
