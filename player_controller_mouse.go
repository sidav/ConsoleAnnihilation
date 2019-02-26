package main

import (
	"SomeTBSGame/routines"
	cw "TCellConsoleWrapper"
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
			plr_moveCursorWithMouse(f)
			return nil
		}
		if cw.WasMouseMovedSinceLastEvent() {
			cx, cy := cw.GetMouseCoords()
			camx, camy := f.cursor.getCameraCoords()
			if areCoordsValid(camx+cx, camy+cy) {
				f.cursor.x, f.cursor.y = camx+cx, camy+cy
				snapCursorToPawn(f)
				return nil
			}
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

func plr_moveCursorWithMouse(f *faction) {
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
