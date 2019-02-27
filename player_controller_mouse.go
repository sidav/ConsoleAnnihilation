package main

import (
	"SomeTBSGame/routines"
	cw "TCellConsoleWrapper"
	"fmt"
	"time"
)

func plr_selectPawnWithMouse(f *faction, m *gameMap) *[]*pawn { // returns a pointer to an array of selected pawns.
	f.cursor.currentCursorMode = CURSOR_SELECT
	for {
		if reRenderNeeded {
			r_renderScreenForFaction(f, m) // TODO: think what to do with all that rendering overkill.
		}
		keyPressed := cw.ReadKeyAsync()
		reRenderNeeded = true

		if cw.IsMouseHeld() && cw.WasMouseMovedSinceLastEvent() && cw.GetMouseButton() == "RIGHT" {
			plr_moveCameraWithMouse(f)
			return nil
		}
		if cw.WasMouseMovedSinceLastEvent() {
			plr_moveCursorWithMouse(f)
			return nil
		}
		if !cw.IsMouseHeld() && cw.GetMouseButton() == "LEFT" {
			u := f.cursor.snappedPawn //m.getUnitAtCoordinates(cx, cy)
			if u == nil {
				return plr_bandboxSelectionWithMouse(f) // select multiple units
			}
			if u.faction.factionNumber != f.factionNumber {
				log.appendMessage("Enemy units can't be selected, Commander.")
				return nil
			}
			return &[]*pawn{f.cursor.snappedPawn}
		}

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
				return plr_bandboxSelectionWithMouse(f) // select multiple units
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
			if routines.ShowSimpleYNChoiceModalWindow("Are you sure you want to quit?") {
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

func plr_bandboxSelectionWithMouse(f *faction) *[]*pawn {
	f.cursor.currentCursorMode = CURSOR_MULTISELECT
	f.cursor.xorig, f.cursor.yorig = f.cursor.getCoords()
	reRenderNeeded = true
	for {
		if reRenderNeeded {
			r_renderScreenForFaction(f, CURRENT_MAP)
		}
		keyPressed := cw.ReadKeyAsync()
		if keyPressed == "ESCAPE" {
			return nil
		}
		if !cw.IsMouseHeld() {
			reRenderNeeded = true
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
		}
		if cw.WasMouseMovedSinceLastEvent() {
			plr_moveCursorWithMouse(f)
		} else {
			reRenderNeeded = false
		}
	}
}

func plr_giveOrderWithMouse(selection *[]*pawn, f *faction) {
	selectedPawn := (*selection)[0] //m.getUnitAtCoordinates(cx, cy)
	log.appendMessage(selectedPawn.name + " is awaiting orders.")
	f.cursor.currentCursorMode = CURSOR_MOVE
	reRenderNeeded = true
	for {
		cx, cy := f.cursor.getCoords()
		if reRenderNeeded {
			r_renderScreenForFaction(f, CURRENT_MAP)
			r_renderSelectedPawns(f, selection)
			r_renderPossibleOrdersForPawn(selectedPawn)
			flushView()
		}

		keyPressed := cw.ReadKeyAsync()
		if cw.IsMouseHeld() && cw.WasMouseMovedSinceLastEvent() && cw.GetMouseButton() == "RIGHT" {
			plr_moveCameraWithMouse(f)
			continue
		}
		if cw.WasMouseMovedSinceLastEvent() {
			plr_moveCursorWithMouse(f)
			continue
		}
		if !cw.IsMouseHeld() && cw.GetMouseButton() == "LEFT" {
			reRenderNeeded = true
			issueDefaultOrderToUnit(selectedPawn, CURRENT_MAP, cx, cy)
			return
		}
		switch keyPressed {
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
					plr_selectBuildingSite(selectedPawn, createBuilding(code, cx, cy, f), CURRENT_MAP)
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
			reRenderNeeded = false
		}
	}
}

func plr_giveOrderForMultiSelectWithMouse(selection *[]*pawn, f *faction) {
	log.appendMessage(fmt.Sprintf("%d units are awaiting orders.", len(*selection)))
	f.cursor.currentCursorMode = CURSOR_MOVE
	reRenderNeeded = true
	for {
		cx, cy := f.cursor.getCoords()

		if reRenderNeeded {
			r_renderScreenForFaction(f, CURRENT_MAP)
			r_renderSelectedPawns(f, selection)
			r_renderPossibleOrdersForMultiselection(f, selection)
			flushView()
		}

		keyPressed := cw.ReadKeyAsync()
		if cw.IsMouseHeld() && cw.WasMouseMovedSinceLastEvent() && cw.GetMouseButton() == "RIGHT" {
			plr_moveCameraWithMouse(f)
			continue
		}
		if cw.WasMouseMovedSinceLastEvent() {
			plr_moveCursorWithMouse(f)
			continue
		}
		if !cw.IsMouseHeld() && cw.GetMouseButton() == "LEFT" {
			for _, p := range *selection {
				issueDefaultOrderToUnit(p, CURRENT_MAP, cx, cy)
			}
			reRenderNeeded = true
			return
		}

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
			reRenderNeeded = false
		}
	}
}

func plr_moveCameraWithMouse(f *faction) {
	vx, vy := cw.GetMouseMovementVector()
	if vx == 0 && vy == 0 {
		return
	}
	cx, cy := f.cursor.getCoords()
	if areCoordsValid(cx+vx, cy+vy) {
		f.cursor.cameraX += vx
		f.cursor.cameraY += vy
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
		snapCursorToPawn(f)
	}
}

func plr_moveCursorWithMouse(f *faction) {
	cx, cy := cw.GetMouseCoords()
	camx, camy := f.cursor.getCameraCoords()

	reRenderNeeded = !(f.cursor.x == camx+cx && f.cursor.y == camy+cy) // rerender is needed if cursor was _actually_ moved

	if areCoordsValid(camx+cx, camy+cy) {
		f.cursor.x, f.cursor.y = camx+cx, camy+cy
		snapCursorToPawn(f)
	}
}
