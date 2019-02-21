package main

import (
	"SomeTBSGame/routines"
	"strconv"
)

const (
	AI_WRITE_DEBUG_TO_LOG = true
)

var (
	// in turns, not ticks
	AI_RECALCULATE_PERIODS_EACH = 50
	AI_CONTROL_PERIOD           = 50
	AI_MIN_CONTROL_PERIOD       = 10
	AI_CONTROL_PERIOD_DECREMENT = 3

	AI_MAX_CONSTRUCTION_ORDERS_AT_A_TIME  = 0
	ai_construction_orders_this_turn      = 0
	ai_max_constr_orders_increment_period = AI_RECALCULATE_PERIODS_EACH * 2
)

func ai_write(text string) {
	if AI_WRITE_DEBUG_TO_LOG {
		log.appendMessage("AI: " + text)
	}
}

func ai_controlFaction(f *faction) {
	ai_recalculateParamsIfNeccessary()

	if getCurrentTurn()%AI_CONTROL_PERIOD != 0 {
		return
	}
	ai_write("assuming direct control over " + f.name)
	ai_construction_orders_this_turn = 0
	for _, p := range CURRENT_MAP.pawns {
		if p.faction == f {
			ai_controlPawn(p)
		}
	}
}

func ai_controlPawn(p *pawn) {
	if p.order != nil {
		return
	}
	if p.canMove() && p.hasWeapons() {
		enemyCommander := ai_getEnemyCommander(p.faction)
		if enemyCommander != nil {
			p.order = &order{orderType: order_attack_move, x: enemyCommander.x, y: enemyCommander.y}
			return
		}
	}
	if p.canConstructUnits() && ai_construction_orders_this_turn < AI_MAX_CONSTRUCTION_ORDERS_AT_A_TIME {
		productionVariants := &p.nanolatherInfo.allowedUnits
		pawnToProduce := ai_decideProduction(productionVariants, p.faction)
		if pawnToProduce != nil {
			p.order = &order{orderType: order_construct, constructingQueue: []*pawn{pawnToProduce}}
			ai_construction_orders_this_turn++
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

func ai_decideProduction(variants *[]string, f *faction) *pawn {
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

func ai_recalculateParamsIfNeccessary() {
	if getCurrentTurn()%AI_RECALCULATE_PERIODS_EACH == 0 {
		if AI_CONTROL_PERIOD-AI_CONTROL_PERIOD_DECREMENT >= AI_MIN_CONTROL_PERIOD {
			AI_CONTROL_PERIOD -= AI_CONTROL_PERIOD_DECREMENT
			ai_write("CONTROL PERIOD changed to " + strconv.Itoa(AI_CONTROL_PERIOD))
		}
	}

	if getCurrentTurn()%ai_max_constr_orders_increment_period == 0 {
		AI_MAX_CONSTRUCTION_ORDERS_AT_A_TIME++
		ai_write("MAX CONSTRUCTION PER TURN changed to " + strconv.Itoa(AI_MAX_CONSTRUCTION_ORDERS_AT_A_TIME))
	}
}
