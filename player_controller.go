package main

import (
	cmenu "github.com/sidav/golibrl/console_menu"
	cw "github.com/sidav/golibrl/console"
	"fmt"
	"time"
)

var (
	PLR_LOOP       = true
	IS_PAUSED      = true
	reRenderNeeded = true
	mouseEnabled   = true
	endTurnPeriod  = 700
	last_time      time.Time
)

func plr_control(f *faction, m *gameMap) {
	PLR_LOOP = true
	snapCursorToPawn(f)
	for PLR_LOOP {
		if mouseEnabled {
			selection := plr_selectPawnWithMouse(f, m)
			if selection != nil {
				if len(*selection) == 1 {
					plr_giveOrderWithMouse(selection, f)
				} else if len(*selection) > 1 {
					plr_giveOrderForMultiSelectWithMouse(selection, f)
				}
			}
		} else {
			selection := plr_selectPawn(f, m)
			if selection != nil {
				if len(*selection) == 1 {
					plr_selectOrder(selection, f, m)
				} else if len(*selection) > 1 {
					plr_selectOrderForMultiSelect(selection, f)
				}
			}
		}
	}
}

func plr_selectPawn(f *faction, m *gameMap) *[]*pawn { // returns a pointer to an array of selected pawns.
	f.cursor.currentCursorMode = CURSOR_SELECT
	for {
		if reRenderNeeded {
			r_renderScreenForFaction(f, m, nil) 
		}
		keyPressed := cw.ReadKeyAsync()
		reRenderNeeded = true
		switch keyPressed {

		case "NOTHING", "NON-KEY":
			if !IS_PAUSED && isTimeToAutoEndTurn() {
				last_time = time.Now()
				PLR_LOOP = false // end turn
				return nil
			} else {
				reRenderNeeded = false
			}
		case ".": // end turn without unpausing the game
			if IS_PAUSED {
				PLR_LOOP = false
				return nil
			}
		case "`":
			mouseEnabled = !mouseEnabled
			if mouseEnabled {
				log.appendMessage("Mouse controls enabled.")
			} else {
				log.appendMessage("Mouse controls disabled.")
			}
		case "SPACE", " ":
			IS_PAUSED = !IS_PAUSED
			if IS_PAUSED {
				log.appendMessage("Tactical pause engaged.")
			} else {
				log.appendMessage("Switched to real-time mode.")
			}
		case "=":
			if endTurnPeriod > 100 {
				endTurnPeriod -= 100
				log.appendMessagef("Game speed increased to %d", 10-(endTurnPeriod/100))
			} else {
				log.appendMessage("Can't increase game speed any further.")
			}
		case "-":
			if endTurnPeriod < 2000 {
				endTurnPeriod += 100
				log.appendMessagef("Game speed decreased to %d", 10-(endTurnPeriod/100))
			} else {
				log.appendMessage("Can't decrease game speed any further.")
			}

		case "ENTER", "RETURN":
			u := f.cursor.snappedPawn //m.getUnitAtCoordinates(cx, cy)
			if u == nil {
				return plr_bandboxSelection(f) // select multiple units
			}
			if u.faction.factionNumber != f.factionNumber {
				log.appendMessage("Enemy units can't be selected, Commander.")
				return nil
			}
			return &[]*pawn{f.cursor.snappedPawn}
		case "TAB":
			trySelectNextIdlePawn(f)
		case "C":
			trySnapCursorToCommander(f)
			return &[]*pawn{f.cursor.snappedPawn}
		case "?":
			if f.cursor.snappedPawn != nil {
				renderPawnInfo(f.cursor.snappedPawn)
			}
		case "ESCAPE":
			if cmenu.ShowSimpleYNChoiceModalWindow("Are you sure you want to quit?") {
				GAME_IS_RUNNING = false
				PLR_LOOP = false
				return nil
			}

		case "DELETE": // cheat
			for _, p := range CURRENT_MAP.pawns {
				if p.faction == f && p.isCommander {
					p.res.metalIncome += 10
					p.res.energyIncome += 50
					return nil
				}
			}
		case "INSERT": // cheat
			CURRENT_MAP.addBuilding(createBuilding("wall", f.cursor.x, f.cursor.y, CURRENT_MAP.factions[1]), true)
		case "HOME": // cheat
			// CURRENT_MAP.addBuilding(createBuilding("lturret", f.cursor.x, f.cursor.y, CURRENT_MAP.factions[0]), true)
			endTurnPeriod = 0
		case "END": // cheat
			CHEAT_IGNORE_FOW = !CHEAT_IGNORE_FOW

		default:
			plr_moveCursor(f, keyPressed)
		}
	}
}

func plr_bandboxSelection(f *faction) *[]*pawn {
	f.cursor.currentCursorMode = CURSOR_MULTISELECT
	f.cursor.xorig, f.cursor.yorig = f.cursor.getCoords()
	for {
		r_renderScreenForFaction(f, CURRENT_MAP, nil)
		keyPressed := cw.ReadKey()
		switch keyPressed {
		case "ESCAPE":
			return nil
		case "ENTER":
			fromx, fromy := f.cursor.xorig, f.cursor.yorig
			tox, toy := f.cursor.getCoords()
			if fromx > tox {
				t := fromx
				fromx = tox
				tox = t
			}
			if fromy > toy {
				t := fromy
				fromy = toy
				toy = t
			}
			unitsInSelection := CURRENT_MAP.getPawnsInRect(fromx, fromy, tox-fromx+1, toy-fromy+1)
			unitsToReturn := make([]*pawn, 0)
			for _, p := range unitsInSelection {
				// select only the pawns if current faction which are capable to move AND attack and are not commanders.
				if p.faction != nil && p.faction == f && p.hasWeapons() && p.canMove() && !p.isCommander {
					unitsToReturn = append(unitsToReturn, p)
				}
			}
			log.appendMessage(fmt.Sprintf("%d units selected from %d", len(unitsToReturn), len(unitsInSelection)))
			return &unitsToReturn
		default:
			plr_moveCursor(f, keyPressed)
		}
	}
}

func plr_selectOrder(selection *[]*pawn, f *faction, m *gameMap) {
	selectedPawn := (*selection)[0] //m.getUnitAtCoordinates(cx, cy)
	log.appendMessage(selectedPawn.name + " is awaiting orders.")
	f.cursor.currentCursorMode = CURSOR_MOVE
	for {
		cx, cy := f.cursor.getCoords()
		r_renderScreenForFaction(f, m, selection)

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

func plr_selectOrderForMultiSelect(selection *[]*pawn, f *faction) {
	log.appendMessage(fmt.Sprintf("%d units are awaiting orders.", len(*selection)))
	f.cursor.currentCursorMode = CURSOR_MOVE
	for {
		cx, cy := f.cursor.getCoords()
		r_renderScreenForFaction(f, CURRENT_MAP, selection)

		keyPressed := cw.ReadKey()
		switch keyPressed {
		case "ENTER", "RETURN":
			for _, p := range *selection {
				issueDefaultOrderToUnit(p, CURRENT_MAP, cx, cy)
			}
			return
		case "a": // attack-move
			f.cursor.currentCursorMode = CURSOR_AMOVE
		case "m": // move
			f.cursor.currentCursorMode = CURSOR_MOVE
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

	indicesQueue := cmenu.ShowSidebarCreateQueueMenu("CONSTRUCT:", p.faction.getFactionColor(),
		SIDEBAR_X, SIDEBAR_FLOOR_2, SIDEBAR_W, SIDEBAR_H-SIDEBAR_FLOOR_2, names, descriptions, presetValues)

	if indicesQueue != nil {
		if len(indicesQueue) > 0 {
			p.setOrder(&order{orderType: order_construct})
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

	index := cmenu.ShowSidebarSingleChoiceMenu("BUILD:", p.faction.getFactionColor(),
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
		cursor := f.cursor
		cx, cy := cursor.getCoords()
		cursor.currentCursorMode = CURSOR_BUILD

		if b.buildingInfo.allowsTightPlacement {
			cursor.w = b.buildingInfo.w
			cursor.h = b.buildingInfo.h
		} else {
			cursor.w = b.buildingInfo.w + 2
			cursor.h = b.buildingInfo.h + 2
		}

		cursor.buildOnMetalOnly = b.buildingInfo.canBeBuiltOnMetalOnly
		cursor.buildOnThermalOnly = b.buildingInfo.canBeBuiltOnThermalOnly
		cursor.radius = b.getMaxRadiusToFire()
		r_renderScreenForFaction(f, m, nil)

		keyPressed := cw.ReadKey()
		switch keyPressed {
		case "ENTER", "RETURN":
			if m.canBuildingBeBuiltAt(b, cx, cy) {
				b.x = cx - b.buildingInfo.w/2
				b.y = cy - b.buildingInfo.h/2
				p.setOrder(&order{orderType: order_build, x: cx, y: cy, buildingToConstruct: b})
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
	f.cursor.centralizeCamera()

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
		snapCursorToPawn(f)
	}
}

func snapCursorToPawn(f *faction) {
	if !(f.areCoordsInSight(f.cursor.x, f.cursor.y) || f.areCoordsInRadarRadius(f.cursor.x, f.cursor.y)) {
		return
	}
	b := CURRENT_MAP.getPawnAtCoordinates(f.cursor.x, f.cursor.y)
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
			f.cursor.centralizeCamera()
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
	case "2", "DOWN":
		return 0, 1
	case "8", "UP":
		return 0, -1
	case "4", "LEFT":
		return -1, 0
	case "6", "RIGHT":
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
	return time.Since(last_time) >= time.Duration(endTurnPeriod)*time.Millisecond
}
