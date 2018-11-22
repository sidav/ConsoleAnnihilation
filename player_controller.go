package main

import cw "TCellConsoleWrapper/tcell_wrapper"

var PLR_LOOP = true

func plr_control(f *faction, m *gameMap) {
	PLR_LOOP = true
	for PLR_LOOP {
		plr_selectEntity(f, m)
	}
}

func plr_selectEntity(f *faction, m *gameMap) {
	r_renderScreenForFaction(f, m)
	renderSelectCursor()
	keyPressed := cw.ReadKey()
	switch keyPressed {
	case "SPACE", " ":
		PLR_LOOP = false // end turn
	case "ENTER", "RETURN":
		plr_giveDefaultOrderToUnit(f, m)
	case "ESCAPE":
		GAME_IS_RUNNING = false
		PLR_LOOP = false
	default:
		plr_moveCursor(m, f, keyPressed)
	}
}

func plr_giveDefaultOrderToUnit(f *faction, m *gameMap) {
	cx, cy := f.cursor.getCoords()
	u := m.getUnitAtCoordinates(cx, cy)
	if u == nil {
		// log.appendMessage("SELECTED NIL")
		return
	}
	if u.faction.factionNumber != f.factionNumber {
		log.appendMessage("Enemy units can't be selected, Commander.")
		return
	} else {
		log.appendMessage(u.name + " is awaiting orders.")
	}

	for {
		r_renderScreenForFaction(f, m)
		renderMoveCursor()
		keyPressed := cw.ReadKey()
		switch keyPressed {
		case "ENTER", "RETURN":
			issueDefaultOrder(u, m, cx, cy)
			return
		case "ESCAPE":
			return
		default:
			plr_moveCursor(m, f, keyPressed)
		}
	}
}

func plr_moveCursor(g *gameMap, f *faction, keyPressed string) {
	vx, vy := plr_keyToDirection(keyPressed)
	cx, cy := f.cursor.getCoords()
	if areCoordsValid(cx+vx, cy+vy) {
		f.cursor.moveByVector(vx, vy)
	}

	snapB := f.cursor.snappedBuilding
	if snapB != nil { // unsnap cursor
		for snapB.isOccupyingCoords(f.cursor.x, f.cursor.y) {
			f.cursor.moveByVector(vx, vy)
		}
		f.cursor.snappedBuilding = nil
	}
	b := g.getBuildingAtCoordinates(f.cursor.x, f.cursor.y)
	if b != nil {
		// snap cursor
		f.cursor.x, f.cursor.y = b.getCenter()
		f.cursor.snappedBuilding = b
	}
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
