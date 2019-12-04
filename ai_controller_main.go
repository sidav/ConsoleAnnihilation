package main

import (
	rnd "github.com/sidav/golibrl/random"
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

	unitLimit, buildingLimit                 int // it MAY be exceeded (dependant on recount units period), be advised.
	currentUnitsCount, currentBuildingsCount int
	currentEngineersCount                    int
	recount_units_period                     int

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

		unitLimit:             25, // it MAY be exceeded (dependant on recount units period), be advised.
		buildingLimit:         15,
		currentUnitsCount:     1,
		currentBuildingsCount: 0,
		currentEngineersCount: 1,
		recount_units_period:  100,
	}
	buildOrderNum := rnd.Random(len(ai_allBuildOrders))
	ai.buildOrder = &(ai_allBuildOrders[buildOrderNum])
	ai_write("SELECTED BUILD ORDER \"" + ai_buildOrderNames[buildOrderNum] + "\"")
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
	// specific Commander orders
	if p.isCommander {
		const RADIUS_FOR_COMMANDER_TO_ATTACK = 10
		// x, y := p.getCenter()
		enemyPawnsInRadius := CURRENT_MAP.getEnemyPawnsInRadiusFromPawn(p, RADIUS_FOR_COMMANDER_TO_ATTACK, p.faction)
		if len(enemyPawnsInRadius) > 0 {
			enemy := enemyPawnsInRadius[0]
			p.order = &order{orderType: order_attack_move, x:enemy.x, y: enemy.y}
			ai_write("Self-protection for the Commander activated.")
			return
		}
	}
	// attack
	if !p.isCommander && p.canMove() && p.hasWeapons() {
		enemyCommander := ai_getEnemyCommander(p.faction)
		if enemyCommander != nil {
			const AMOVE_RADIUS = 15
			x, y := -1, -1
			for !(areCoordsValid(x, y) && CURRENT_MAP.tileMap[x][y].isPassable) {
				x = rnd.RandInRange(enemyCommander.x - AMOVE_RADIUS, enemyCommander.x + AMOVE_RADIUS)
				y = rnd.RandInRange(enemyCommander.y - AMOVE_RADIUS, enemyCommander.y + AMOVE_RADIUS)
			}
			p.order = &order{orderType: order_attack_move, x: x, y: y}
			return
		}
		p.order = &order{orderType: order_attack_move, x: rnd.Random(mapW), y: rnd.Random(mapH)}
	}

	if currAi.construction_orders_this_turn < currAi.MAX_CONSTRUCTION_ORDERS_AT_A_TIME {
		// build
		if p.canConstructBuildings() && currAi.currentBuildingsCount < currAi.buildingLimit {
			currAi.ai_decideConstruction(p)
		}
		// produce
		if p.canConstructUnits() && currAi.currentUnitsCount < currAi.unitLimit {
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
		currAi.recountUnitsAndBuildings(f)
	}
}

func (currAi *aiData) recountUnitsAndBuildings(f *faction) {

	currAi.currentUnitsCount = 0
	currAi.currentBuildingsCount = 0
	currAi.currentEngineersCount = 0

	for _, p := range CURRENT_MAP.pawns {
		if p.faction == f {
			if p.isBuilding() {
				currAi.currentBuildingsCount++
			} else {
				currAi.currentUnitsCount++
			}
			if p.canConstructBuildings() {
				currAi.currentEngineersCount++
			}
		}
	}
	ai_write("I've got " + strconv.Itoa(currAi.currentUnitsCount) + " minions right now.")
	ai_write("I've got " + strconv.Itoa(currAi.currentBuildingsCount) + " buildings right now.")
	ai_write("I've got " + strconv.Itoa(currAi.currentEngineersCount) + " engineers right now.")

}
