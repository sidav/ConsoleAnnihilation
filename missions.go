package main

var MIS1_MAP = &[]string {
	"~~~.........................^............................................",
	"~~..;;......................^............................................",
	"~~..;..............;;.......^^......;;;..................................",
	"~.....;.....................^^^.....;;;;.................................",
	"~~..........................^^^.....;;;..................................",
	"~~~..........................^...........................................",
	"~~.......................................................................",
	"~~.......................................................................",
	"~.....................................^..................................",
	"~..................;............^^^.^^^^^................................",
	"~..........$........;;......^^^^^.^^^^..^^^^^^...........................",
	"~..........$........;.......^^..........^................................",
	"~~...........................^^^.....$...................................",
	"~..............................^....$$$..................................",
	"~~..................................$$...................................",
	"~~~.........................^^...........................................",
	"~~~.........;................^^..........................................",
	"~~~~.......;;;................^^.........................................",
	"~~~~........................^^^..........................................",
	"~~~~~.....................^^^^...........................................",
}

func initMapForMission(g *gameMap, missionNumber int) {
	g.initTileMap(MIS1_MAP)

	g.factions = append(g.factions, createFaction("The Core Corporation", 0, true, false))
	g.addPawn(createUnit("protocommander", 3, 9, g.factions[0], true))
	g.factions[0].cursor.x = 3
	g.factions[0].cursor.y = 9
	//
	//g.addPawn(createUnit("coreraider", 4, 5, g.factions[0], true))
	//g.addPawn(createUnit("coreraider", 4, 6, g.factions[0], true))
	//g.addPawn(createUnit("coreraider", 5, 5, g.factions[0], true))
	//g.addPawn(createUnit("coreraider", 5, 6, g.factions[0], true))
	//g.addPawn(createUnit("coreweasel", 7, 7, g.factions[0], true))


	g.factions = append(g.factions, createFaction("The rogue Arm AI", 1, false, true))
	// g.addPawn(createUnit("armcommander", mapW-10, 5, g.factions[1], true))
	g.addBuilding(createBuilding("armhq", mapW-5, 9, g.factions[1]), true)
	g.addBuilding(createBuilding("armkbotlab", 33, 1, g.factions[1]), true)
	g.addBuilding(createBuilding("armvehfactory", mapW-15, 3, g.factions[1]), true)
	g.addBuilding(createBuilding("armkbotlab", mapW-15, mapH-5, g.factions[1]), true)
	g.addBuilding(createBuilding("mstorage", 20, 10, g.factions[1]), true)
	g.addBuilding(createBuilding("lturret", mapW-10, 1, g.factions[1]), true)
	g.addBuilding(createBuilding("lturret", mapW-10, 4, g.factions[1]), true)
	g.addBuilding(createBuilding("guardian", mapW-7, 3, g.factions[1]), true)
	g.addBuilding(createBuilding("lturret", mapW-10, 8, g.factions[1]), true)
	g.addBuilding(createBuilding("lturret", mapW-10, 12, g.factions[1]), true)
	g.addBuilding(createBuilding("lturret", mapW-10, 16, g.factions[1]), true)
	g.addBuilding(createBuilding("guardian", mapW-7, 14, g.factions[1]), true)
	g.addBuilding(createBuilding("lturret", mapW-10, 19, g.factions[1]), true)
}

func checkWinOrLose() { // TEMPORARY
	if getCurrentTurn() % 10 != 0 {
		return
	}
	plrAlive := false
	enemyAlive := false
	for _, p := range CURRENT_MAP.pawns {
		if p.isCommander {
			if p.faction.playerControlled {
				plrAlive = true
			}
			if p.faction.aiControlled {
				enemyAlive = true
			}
		}
	}
	if !plrAlive {
		GAME_IS_RUNNING = false
		r_gamelostScreen()
		return
	}
	if !enemyAlive {
		GAME_IS_RUNNING = false
		r_gameWonScreen()
		return
	}
}
