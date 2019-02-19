package main

import (
	"SomeTBSGame/routines"
	cw "TCellConsoleWrapper"
	"time"
)

var (
	PLR_LOOP       = true
	IS_PAUSED      = true
	reRenderNeeded = true
	endTurnEachMs  = 700
	last_time      time.Time
)

func plr_control(f *faction, m *gameMap) {
	PLR_LOOP = true
	snapCursorToPawn(f, m)
	for PLR_LOOP {
		if plr_selectPawn(f, m) {
			// plr_selectOrder(f, m)
			plr_selectOrder(f, m)
		}
	}
}

func plr_selectPawn(f *faction, m *gameMap) bool { // true if pawn was selected
	f.cursor.currentCursorMode = CURSOR_SELECT
	for {
		if reRenderNeeded {
			r_renderScreenForFaction(f, m) // TODO: think what to do with all that rendering overkill.
		}
		keyPressed := cw.ReadKeyAsync()
		reRenderNeeded = true
		switch keyPressed {

		case "NOTHING":
			if !IS_PAUSED && isTimeToAutoEndTurn() {
				last_time = time.Now()
				PLR_LOOP = false // end turn
				return false
			} else {
				reRenderNeeded = false
			}
		case ".": // end turn without unpausing the game
			if IS_PAUSED {
				PLR_LOOP = false
				return false 
			}
		case "SPACE", " ":
			IS_PAUSED = !IS_PAUSED
			if IS_PAUSED {
				log.appendMessage("Tactical pause engaged.")
			} else {
				log.appendMessage("Switched to real-time mode.")
			}
		case "=":
			if endTurnEachMs > 100 {
				endTurnEachMs -= 100
				log.appendMessagef("Game speed increased to %d", 10-(endTurnEachMs/100))
			} else {
				log.appendMessage("Can't increase game speed any further.")
			}
		case "-":
			if endTurnEachMs < 2000 {
				endTurnEachMs += 100
				log.appendMessagef("Game speed decreased to %d", 10-(endTurnEachMs/100))
			} else {
				log.appendMessage("Can't decrease game speed any further.")
			}

		case "ENTER", "RETURN":
			u := f.cursor.snappedPawn //m.getUnitAtCoordinates(cx, cy)
			if u == nil {
				plr_bandboxSelection(f) // select multiple units
				return false
			}
			if u.faction.factionNumber != f.factionNumber {
				log.appendMessage("Enemy units can't be selected, Commander.")
				return false
			}
			return true
		case "TAB":
			trySelectNextIdlePawn(f)
		case "C":
			return trySnapCursorToCommander(f)
		case "ESCAPE":
			GAME_IS_RUNNING = false
			PLR_LOOP = false
			return false

		case "DELETE": // cheat
			for _, p := range CURRENT_MAP.pawns {
				if p.faction == f && p.isCommander {
					p.res.metalIncome += 10
					p.res.energyIncome += 50
					return false
				}
			}

		default:
			plr_moveCursor(f, keyPressed)
		}
	}
}

func plr_bandboxSelection(f *faction) {
	f.cursor.currentCursorMode = CURSOR_MULTISELECT
	f.cursor.xorig, f.cursor.yorig = f.cursor.getCoords()
	for {
		r_renderScreenForFaction(f, CURRENT_MAP)
		keyPressed := cw.ReadKey()
		switch keyPressed {
		case "ESCAPE":
			return
		case "ENTER":
			return
		default:
			plr_moveCursor(f, keyPressed)
		}
	}
}

func plr_selectOrder(f *faction, m *gameMap) {
	selectedPawn := f.cursor.snappedPawn //m.getUnitAtCoordinates(cx, cy)
	log.appendMessage(selectedPawn.name + " is awaiting orders.")
	f.cursor.currentCursorMode = CURSOR_MOVE
	for {
		cx, cy := f.cursor.getCoords()
		r_renderScreenForFaction(f, m)
		r_renderPossibleOrdersForPawn(selectedPawn)
		flushView()

		keyPressed := cw.ReadKey()
		switch keyPressed {
		case "ENTER", "RETURN":
			issueDefaultOrderToUnit(selectedPawn, m, cx, cy)
			return
		case "a": // attack-move
			if selectedPawn.hasWeapons() || selectedPawn.canConstructUnits() {
				f.cursor.currentCursorMode = CURSOR_AMOVE
			}
		case "m": // move
				f.cursor.currentCursorMode = CURSOR_MOVE
		case "b": // build
			if selectedPawn.canConstructBuildings() {
				code := plr_selectBuidingToConstruct(selectedPawn)
				if code != "" {
					plr_selectBuildingSite(selectedPawn, createBuilding(code, cx, cy, f), m)
					return
				}
			}
		case "c": // construct units
			if selectedPawn.canConstructUnits() {
				plr_selectUnitsToConstruct(selectedPawn)
			}
		case "r": // repeat construction queue
			if selectedPawn.canConstructUnits() {
				selectedPawn.repeatConstructionQueue = !selectedPawn.repeatConstructionQueue
			}
		case "ESCAPE":
			return
		default:
			plr_moveCursor(f, keyPressed)
		}
	}
}

func plr_selectUnitsToConstruct(p *pawn) {
	availableUnitCodes := p.nanolatherInfo.allowedUnits

	names := make([]string, 0)
	descriptions := make([]string, 0)

	// descriptions := make([]string, 0)
	for _, code := range availableUnitCodes {
		name, desc := getUnitNameAndDescription(code)
		names = append(names, name)
		descriptions = append(descriptions, desc)
	}

	presetValues := make([]int, 0)
	// init values array for already existing queue
	if p.order != nil && p.order.constructingQueue != nil {
		for _, pawnInQueue := range p.order.constructingQueue {
			for i, name := range names {
				if pawnInQueue.name == name {
					presetValues = append(presetValues, i)
				}
			}
		}
	}

	indicesQueue := routines.ShowSidebarCreateQueueMenu("CONSTRUCT:", p.faction.getFactionColor(),
		SIDEBAR_X, SIDEBAR_FLOOR_2, SIDEBAR_W, SIDEBAR_H-SIDEBAR_FLOOR_2, names, descriptions, presetValues)

	if indicesQueue != nil {
		if len(indicesQueue) > 0 {
			p.order = &order{orderType: order_construct}
			for _, i := range indicesQueue {
				p.order.constructingQueue = append(p.order.constructingQueue,
					createUnit(availableUnitCodes[i], p.x, p.y, p.faction, false))
			}
			log.appendMessagef("Construction of %d units initiated.", len(p.order.constructingQueue))
		} else {
			p.order = nil
			log.appendMessage("Construction orders cancelled.")
		}
	}
}

func plr_selectBuidingToConstruct(p *pawn) string {
	availableBuildingCodes := p.nanolatherInfo.allowedBuildings

	names := make([]string, 0)
	descriptions := make([]string, 0)
	for _, code := range availableBuildingCodes {
		name, desc := getBuildingNameAndDescription(code)
		names = append(names, name)
		descriptions = append(descriptions, desc)
	}

	index := routines.ShowSidebarSingleChoiceMenu("BUILD:", p.faction.getFactionColor(),
		SIDEBAR_X, SIDEBAR_FLOOR_2, SIDEBAR_W, SIDEBAR_H-SIDEBAR_FLOOR_2, names, descriptions)
	if index != -1 {
		return availableBuildingCodes[index]
	}
	return ""
}

func plr_selectBuildingSite(p *pawn, b *pawn, m *gameMap) {
	log.appendMessage("Select construction site for " + b.name)
	for {
		f := p.faction
		cx, cy := f.cursor.getCoords()
		f.cursor.currentCursorMode = CURSOR_BUILD
		f.cursor.w = b.buildingInfo.w
		f.cursor.h = b.buildingInfo.h
		f.cursor.buildOnMetalOnly = b.buildingInfo.canBeBuiltOnMetalOnly
		f.cursor.buildOnThermalOnly = b.buildingInfo.canBeBuiltOnThermalOnly
		r_renderScreenForFaction(f, m)
		flushView()

		keyPressed := cw.ReadKey()
		switch keyPressed {
		case "ENTER", "RETURN":
			if m.canBuildingBeBuiltAt(b, cx, cy) {
				b.x = cx - b.buildingInfo.w/2
				b.y = cy - b.buildingInfo.h/2
				p.order = &order{orderType: order_build, x: cx, y: cy, buildingToConstruct: b}
				return
			} else {
				log.appendMessage("This building can't be placed here!")
			}
		case "ESCAPE":
			log.appendMessage("Construction cancelled: " + b.name)
			return
		default:
			plr_moveCursor(f, keyPressed)
		}
	}
}

func plr_moveCursor(f *faction, keyPressed string) {
	vx, vy := plr_keyToDirection(keyPressed)
	if vx == 0 && vy == 0 {
		return
	}
	cx, cy := f.cursor.getCoords()
	if areCoordsValid(cx+vx, cy+vy) {
		f.cursor.moveByVector(vx, vy)
	}

	snapB := f.cursor.snappedPawn
	if snapB != nil { // unsnap cursor
		for snapB.isOccupyingCoords(f.cursor.x, f.cursor.y) {
			if areCoordsValid(f.cursor.x+vx, f.cursor.y+vy) {
				f.cursor.moveByVector(vx, vy)
			} else {
				break
			}
		}
		f.cursor.snappedPawn = nil
	}
	if f.cursor.currentCursorMode != CURSOR_BUILD {
		snapCursorToPawn(f, CURRENT_MAP)
	}
}

func snapCursorToPawn(f *faction, g *gameMap) {
	b := g.getPawnAtCoordinates(f.cursor.x, f.cursor.y)
	if b == nil {
		f.cursor.snappedPawn = nil
	} else {
		f.cursor.x, f.cursor.y = b.getCenter()
		f.cursor.snappedPawn = b
	}
}

func trySnapCursorToCommander(f *faction) bool {
	for _, p := range CURRENT_MAP.pawns {
		if p.faction == f && p.isCommander {
			f.cursor.x, f.cursor.y = p.getCoords()
			f.cursor.snappedPawn = p
			return true
		}
	}
	return false
}

func trySelectNextIdlePawn(f *faction) {
	totalPawnsOnMap := len(CURRENT_MAP.pawns)
	for offset := 0; offset < totalPawnsOnMap; offset++ {
		index_to_select := (offset + f.cursor.lastSelectedIdlePawnIndex) % totalPawnsOnMap

		p := CURRENT_MAP.pawns[index_to_select]
		if p.faction == f && (p.order == nil || p.order.orderType == order_hold) {
			log.appendMessage("Next idle unit selected.")
			f.cursor.lastSelectedIdlePawnIndex = index_to_select + 1
			f.cursor.x, f.cursor.y = p.getCenter()
			f.cursor.snappedPawn = p
			return
		}
	}
	log.appendMessage("There currently are no idle units in your army, Commander.")
}

func plr_keyToDirection(keyPressed string) (int, int) {
	switch keyPressed {
	case "2":
		return 0, 1
	case "8":
		return 0, -1
	case "4":
		return -1, 0
	case "6":
		return 1, 0
	case "7":
		return -1, -1
	case "9":
		return 1, -1
	case "1":
		return -1, 1
	case "3":
		return 1, 1
	default:
		return 0, 0
	}
}

func isTimeToAutoEndTurn() bool {
	return time.Since(last_time) >= time.Duration(endTurnEachMs)*time.Millisecond
}
