package main

import (
	"SomeTBSGame/routines"
	"strconv"
)

type aiData struct {
	// in turns, not ticks
	RECALCULATE_PERIODS_EACH    int
	CONTROL_PERIOD              int
	AI_MIN_CONTROL_PERIOD       int
	AI_CONTROL_PERIOD_DECREMENT int

	MAX_CONSTRUCTION_ORDERS_AT_A_TIME  int
	construction_orders_this_turn      int
	max_constr_orders_increment_period int

	unit_limit            int // it MAY be exceeded (dependant on recount units period), be advised.
	current_units_count   int // 99999 is for preventing AI production stucking. The var will be recalculated later anyway.
	currentEngineersCount int
	recount_units_period  int

	buildOrder        *[]*ai_buildOrderStep // not meant to be changed
	currentStepNumber int
}

func ai_createAiData() *aiData {

	ai := &aiData{
		RECALCULATE_PERIODS_EACH:           50,
		CONTROL_PERIOD:                     10, // 80
		AI_MIN_CONTROL_PERIOD:              10,
		AI_CONTROL_PERIOD_DECREMENT:        5,
		MAX_CONSTRUCTION_ORDERS_AT_A_TIME:  1,
		construction_orders_this_turn:      0,
		max_constr_orders_increment_period: 100,
		unit_limit:                         25, // it MAY be exceeded (dependant on recount units period), be advised.
		current_units_count:                1,  // 99999 is for preventing AI production stucking. The var will be recalculated later anyway.
		currentEngineersCount:              1,
		recount_units_period:               100,
	}
	buildOrderNum := routines.Random(len(ai_allBuildOrders))
	ai.buildOrder = &(ai_allBuildOrders[buildOrderNum])
	ai_write("SELECTED BUILD ORDER #" + strconv.Itoa(buildOrderNum+1))
	return ai
}

const (
	AI_WRITE_DEBUG_TO_LOG = true
)

func ai_write(text string) {
	if AI_WRITE_DEBUG_TO_LOG {
		log.appendMessage("AI: " + text)
	}
}

func ai_controlFaction(f *faction) {
	currAi := f.aiData
	currAi.ai_recalculateParamsIfNeccessary(f)

	if getCurrentTurn()%currAi.CONTROL_PERIOD != 0 {
		return
	}
	ai_write("assuming direct control over " + f.name)
	currAi.construction_orders_this_turn = 0
	for _, p := range CURRENT_MAP.pawns {
		if p.faction == f {
			ai_controlPawn(currAi, p)
		}
	}
}

func ai_controlPawn(currAi *aiData, p *pawn) {
	if p.order != nil {
		return
	}
	// attack
	if !p.isCommander && p.canMove() && p.hasWeapons() {
		enemyCommander := ai_getEnemyCommander(p.faction)
		if enemyCommander != nil {
			p.order = &order{orderType: order_attack_move, x: enemyCommander.x, y: enemyCommander.y}
			return
		}
		p.order = &order{orderType: order_attack_move, x: routines.Random(mapW), y: routines.Random(mapH)}
	}
	// construct
	if currAi.current_units_count < currAi.unit_limit && currAi.construction_orders_this_turn < currAi.MAX_CONSTRUCTION_ORDERS_AT_A_TIME {
		if p.canConstructBuildings() {
			currAi.ai_decideConstruction(p)
		}
		// produce
		if p.canConstructUnits() {
			currAi.ai_decideProduction(p)
		}
	}
}

func ai_getEnemyCommander(f *faction) *pawn {
	for _, p := range CURRENT_MAP.pawns {
		if p.faction != f && p.isCommander {
			return p
		}
	}
	return nil
}

func (currAi *aiData) ai_recalculateParamsIfNeccessary(f *faction) {
	if getCurrentTurn()%currAi.RECALCULATE_PERIODS_EACH == 0 {
		if currAi.CONTROL_PERIOD-currAi.AI_CONTROL_PERIOD_DECREMENT >= currAi.AI_MIN_CONTROL_PERIOD {
			currAi.CONTROL_PERIOD -= currAi.AI_CONTROL_PERIOD_DECREMENT
			ai_write("CONTROL PERIOD changed to " + strconv.Itoa(currAi.CONTROL_PERIOD))
		}
	}

	if getCurrentTurn()%currAi.max_constr_orders_increment_period == 0 {
		currAi.MAX_CONSTRUCTION_ORDERS_AT_A_TIME++
		ai_write("MAX CONSTRUCTION PER TURN changed to " + strconv.Itoa(currAi.MAX_CONSTRUCTION_ORDERS_AT_A_TIME))
	}

	if getCurrentTurn()%currAi.recount_units_period == 0 {
		currAi.recountUnitsAndCheckBuildsSatisfations(f)
	}
}

func (currAi *aiData) recountUnitsAndCheckBuildsSatisfations(f *faction) {

	currAi.current_units_count = 0
	currAi.currentEngineersCount = 0

	for _, p := range CURRENT_MAP.pawns {
		if p.faction == f {
			currAi.current_units_count++
			if p.canConstructBuildings() {
				currAi.currentEngineersCount++
			}
		}
	}
	ai_write("I've got " + strconv.Itoa(currAi.current_units_count) + " minions right now.")

}
