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

	if builder.faction.economy.metalIncome < ai.getCurrentOrderStep().desiredMIncome {
		ai.ai_buildMetalIncome(builder)
		return
	}

	if builder.faction.economy.energyIncome < ai.getCurrentOrderStep().desiredEIncome {
		ai.ai_buildEnergyIncome(builder)
		return
	}

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

func (ai *aiData) ai_buildEnergyIncome(builder *pawn) {
	variants := builder.nanolatherInfo.allowedBuildings
	final_build_variant := ""
	for _, variant := range variants {
		candidate := createBuilding(variant, 0, 0, builder.faction)
		if candidate.res != nil && candidate.res.energyIncome > 0 {
			final_build_variant = variant
		}
	}
	ai_makeBuildOrderForBuilding(builder, final_build_variant)
}

func (ai *aiData) ai_buildMetalIncome(builder *pawn) {
	variants := builder.nanolatherInfo.allowedBuildings
	final_build_variant := ""
	for _, variant := range variants {
		candidate := createBuilding(variant, 0, 0, builder.faction)
		if candidate.res != nil && candidate.res.isMetalExtractor {
			final_build_variant = variant
		}
	}
	ai_tryBuildMetalExtractor(builder, final_build_variant)
}
