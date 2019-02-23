package main

import (
	"SomeTBSGame/routines"
	"strconv"
)

type aiData struct {
	// in turns, not ticks
	AI_RECALCULATE_PERIODS_EACH int
	AI_CONTROL_PERIOD           int
	AI_MIN_CONTROL_PERIOD       int
	AI_CONTROL_PERIOD_DECREMENT int

	AI_MAX_CONSTRUCTION_ORDERS_AT_A_TIME  int
	ai_construction_orders_this_turn      int
	ai_max_constr_orders_increment_period int

	ai_unit_limit           int    // it MAY be exceeded (dependant on recount units period), be advised.
	ai_current_units_count  int // 99999 is for preventing AI production stucking. The var will be recalculated later anyway.
	ai_recount_units_period int
}

func ai_createAiData() *aiData {
	return &aiData{
		AI_RECALCULATE_PERIODS_EACH: 50, 
		AI_CONTROL_PERIOD          : 80, 
		AI_MIN_CONTROL_PERIOD      : 10, 
		AI_CONTROL_PERIOD_DECREMENT: 5,
		AI_MAX_CONSTRUCTION_ORDERS_AT_A_TIME : 0, 
		ai_construction_orders_this_turn     : 0, 
		ai_max_constr_orders_increment_period: 100,
		ai_unit_limit          : 25,    // it MAY be exceeded (dependant on recount units period), be advised.
		ai_current_units_count : 99999, // 99999 is for preventing AI production stucking. The var will be recalculated later anyway.
		ai_recount_units_period: 100,
	}
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

	if getCurrentTurn()%currAi.AI_CONTROL_PERIOD != 0 {
		return
	}
	ai_write("assuming direct control over " + f.name)
	currAi.ai_construction_orders_this_turn = 0
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
	if p.canMove() && p.hasWeapons() {
		enemyCommander := ai_getEnemyCommander(p.faction)
		if enemyCommander != nil {
			p.order = &order{orderType: order_attack_move, x: enemyCommander.x, y: enemyCommander.y}
			return
		}
		p.order = &order{orderType: order_attack_move, x: routines.Random(mapW), y: routines.Random(mapH)}
	}
	// produce
	if p.canConstructUnits() && currAi.ai_current_units_count < currAi.ai_unit_limit &&
		currAi.ai_construction_orders_this_turn < currAi.AI_MAX_CONSTRUCTION_ORDERS_AT_A_TIME {
		productionVariants := &p.nanolatherInfo.allowedUnits
		pawnToProduce := ai_decideProduction(currAi, productionVariants, p.faction)
		if pawnToProduce != nil {
			p.order = &order{orderType: order_construct, constructingQueue: []*pawn{pawnToProduce}}
			currAi.ai_construction_orders_this_turn++
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

func ai_decideProduction(currAi *aiData, variants *[]string, f *faction) *pawn {
	listOfCombatUnits := make([]*pawn, 0)
	for _, variant := range *variants {
		pawnUnderConsideration := createUnit(variant, 0, 0, f, false)
		if pawnUnderConsideration.canMove() && pawnUnderConsideration.hasWeapons() {
			listOfCombatUnits = append(listOfCombatUnits, pawnUnderConsideration)
		}
	}
	if len(listOfCombatUnits) > 0 {
		pwn := listOfCombatUnits[routines.Random(len(listOfCombatUnits))]
		ai_write("producing " + pwn.name)
		return pwn
	}
	ai_write("nothing to produce.")
	return nil
}

func (currAi *aiData) ai_recalculateParamsIfNeccessary(f *faction) {
	if getCurrentTurn()%currAi.AI_RECALCULATE_PERIODS_EACH == 0 {
		if currAi.AI_CONTROL_PERIOD-currAi.AI_CONTROL_PERIOD_DECREMENT >= currAi.AI_MIN_CONTROL_PERIOD {
			currAi.AI_CONTROL_PERIOD -= currAi.AI_CONTROL_PERIOD_DECREMENT
			ai_write("CONTROL PERIOD changed to " + strconv.Itoa(currAi.AI_CONTROL_PERIOD))
		}
	}

	if getCurrentTurn()%currAi.ai_max_constr_orders_increment_period == 0 {
		currAi.AI_MAX_CONSTRUCTION_ORDERS_AT_A_TIME++
		ai_write("MAX CONSTRUCTION PER TURN changed to " + strconv.Itoa(currAi.AI_MAX_CONSTRUCTION_ORDERS_AT_A_TIME))
	}

	if getCurrentTurn()%currAi.ai_recount_units_period == 0 {
		currAi.ai_current_units_count = 0
		for _, p := range CURRENT_MAP.pawns {
			if p.faction == f {
				currAi.ai_current_units_count++
			}
		}
		ai_write("I've got " + strconv.Itoa(currAi.ai_current_units_count) + " minions right now.")
	}
}
