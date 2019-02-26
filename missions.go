package main

import (
	"SomeTBSGame/routines"
	"strconv"
)

//var MIS1_MAP = &[]string {
//	"~~~.........................^............................................",
//	"~~..;;......................^............................................",
//	"~~..;..............;;.......^^......;;;..................................",
//	"~.....;.....................^^^.....;;;;.................................",
//	"~~..........................^^^.....;;;..................................",
//	"~~~..........................^...........................................",
//	"~~.......................................................................",
//	"~~.......................................................................",
//	"~.....................................^..................................",
//	"~..................;............^^^.^^^^^................................",
//	"~..........$........;;......^^^^^.^^^^..^^^^^^...........................",
//	"~..........$........;.......^^..........^................................",
//	"~~...........................^^^.....$...................................",
//	"~..............................^....$$$..................................",
//	"~~..................................$$...................................",
//	"~~~.........................^^...........................................",
//	"~~~.........;................^^..........................................",
//	"~~~~.......;;;................^^.........................................",
//	"~~~~........................^^^..........................................",
//	"~~~~~.....................^^^^...........................................",
//}

var MIS1_MAP = &[]string {
	"........................................^^^^^^^^^^........................................",
	"........................................^^^^^^^^^^........................................",
	"........................................^^^^^^^^^^........................................",
	"..;;;...................................^^^^.......................................;;;....",
	"..;;;...................................^^^^.^^^^^.................................;;;....",
	"........................................^^^^.^^^^^........................................",
	".............................................^^^^^........................................",
	"........................................^^^^^^^^^^........................................",
	"........................................^^^^^^^^^^........................................",
	"..;;;...................................^^^^^^^^^^.................................;;;....",
	"........................................^^^^^^^^^^........................................",
	"........................................^^^^^^^^^^........................................",
	"..............................................^^^^........................................",
	"........................................^^^^^.^^^^........................................",
	"........................................^^^^^.^^^^........................................",
	"........................................^^^^^.^^^^........................................",
	"........................................^^^^^.............................................",
	"........................................^^^^^^^^^^........................................",
	"........................................^^^^^^^^^^........................................",
	"........................................^^^^^^^^^^........................................",
}

func initMapForMission(g *gameMap, missionNumber int) {
	g.initTileMap(MIS1_MAP)

	ai_write("Seed is " + strconv.Itoa(routines.Randomize()))

	g.factions = append(g.factions, createFaction("AI 1", 0,true, true))
	g.addPawn(createUnit("armcommander", 7, mapH/2, g.factions[0], true))
	g.factions[0].cursor.centralizeCamera()
	//g.factions[0].cursor.x = 7
	//g.factions[0].cursor.y = mapH/2

	// g.addPawn(createUnit("coreck", 3, 3, g.factions[0], true))

	g.factions = append(g.factions, createFaction("AI 2", 1, false, true))
	g.addPawn(createUnit("corecommander", mapW - 10, mapH/2, g.factions[1], true))
	// g.addPawn(createUnit("coreck", 3, 3, g.factions[0], true))

	// g.factions = append(g.factions, createFaction("OBSERVER", 0, true, false))
	CHEAT_IGNORE_FOW = true
}

func checkWinOrLose() { // TEMPORARY
	//if getCurrentTurn() % 10 != 0 {
	//	return
	//}
	//plrAlive := false
	//enemyAlive := false
	//for _, p := range CURRENT_MAP.pawns {
	//	if p.isCommander {
	//		if p.faction.playerControlled {
	//			plrAlive = true
	//		}
	//		if p.faction.aiControlled {
	//			enemyAlive = true
	//		}
	//	}
	//}
	//if !plrAlive {
	//	GAME_IS_RUNNING = false
	//	r_gamelostScreen()
	//	return
	//}
	//if !enemyAlive {
	//	GAME_IS_RUNNING = false
	//	r_gameWonScreen()
	//	return
	//}
}
