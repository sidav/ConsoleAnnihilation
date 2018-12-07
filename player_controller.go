package main

import (
	"SomeTBSGame/routines"
	cw "TCellConsoleWrapper/tcell_wrapper"
)

var PLR_LOOP = true

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
		r_renderScreenForFaction(f, m)
		keyPressed := cw.ReadKey()
		switch keyPressed {
		case "SPACE", " ":
			PLR_LOOP = false // end turn
			return false
		case "ENTER", "RETURN":
			u := f.cursor.snappedPawn //m.getUnitAtCoordinates(cx, cy)
			if u == nil {
				// log.appendMessage("SELECTED NIL")
				return false
			}
			if u.faction.factionNumber != f.factionNumber {
				log.appendMessage("Enemy units can't be selected, Commander.")
				return false
			}
			return true
		case "C":
			return trySnapCursorToCommander(f)
		case "ESCAPE":
			GAME_IS_RUNNING = false
			PLR_LOOP = false
			return false
		default:
			plr_moveCursor(m, f, keyPressed)
		}
	}
}

func plr_selectOrder(f *faction, m *gameMap) {
	u := f.cursor.snappedPawn //m.getUnitAtCoordinates(cx, cy)
	log.appendMessage(u.name + " is awaiting orders.")

	for {
		cx, cy := f.cursor.getCoords()
		f.cursor.currentCursorMode = CURSOR_MOVE
		r_renderScreenForFaction(f, m)
		r_renderPossibleOrdersForPawn(u)
		flushView()

		keyPressed := cw.ReadKey()
		switch keyPressed {
		case "ENTER", "RETURN":
			issueDefaultOrderToUnit(u, m, cx, cy)
			return
		case "b": // build
			if u.canBuild() {
				code := plr_selectBuidingToConstruct(u)
				if code != "" {
					plr_selectBuildingSite(u, createBuilding(code, cx, cy, f), m)
					return
				}
			}
		case "ESCAPE":
			return
		default:
			plr_moveCursor(m, f, keyPressed)
		}
	}
}

func plr_selectBuidingToConstruct(p *pawn) string {
	avail_buildings := p.nanolatherInfo.allowedBuildings
	index := routines.ShowSidebarSingleSelectMenu("BUILD:", p.faction.getFactionColor(),
		SIDEBAR_X, SIDEBAR_FLOOR_2,  SIDEBAR_W,  SIDEBAR_H - SIDEBAR_FLOOR_2, avail_buildings)
	if index != -1 {
		return avail_buildings[index]
	}
	return ""
}

func plr_selectBuildingSite(p *pawn, b *pawn, m *gameMap) {
	log.appendMessage("Select construction site for "+b.name)
	for {
		f := p.faction
		cx, cy := f.cursor.getCoords()
		f.cursor.currentCursorMode = CURSOR_BUILD
		f.cursor.w = b.buildingInfo.w
		f.cursor.h = b.buildingInfo.h
		r_renderScreenForFaction(f, m)
		flushView()

		keyPressed := cw.ReadKey()
		switch keyPressed {
		case "ENTER", "RETURN":
			b.x = cx - b.buildingInfo.w/2
			b.y = cy - b.buildingInfo.h/2
			p.order = &order{orderType: order_build, x: cx, y: cy, targetBuilding: b}
			return
		case "ESCAPE":
			log.appendMessage("Construction cancelled: "+b.name)
			return
		default:
			plr_moveCursor(m, f, keyPressed)
		}
	}
}

func plr_moveCursor(g *gameMap, f *faction, keyPressed string) {
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
		snapCursorToPawn(f, g)
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
		if p.faction == f && p.name == "Commander" {
			f.cursor.x, f.cursor.y = p.getCoords()
			f.cursor.snappedPawn = p
			return true
		}
	}
	return false
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
