package main

import cw "GoSdlConsole/GoSdlConsole"

var PLR_LOOP = true

func plr_control(f *faction, m *gameMap) {
	PLR_LOOP = true
	for PLR_LOOP {
		plr_selectUnit(f, m)
	}
}

func plr_selectUnit(f *faction, m *gameMap) {
	r_renderScreenForFaction(f, m)
	renderSelectCursor()
	keyPressed := cw.ReadKey()
	switch keyPressed {
	case "SPACE":
		PLR_LOOP = false
	case "ENTER", "RETURN":
		plr_giveAnOrderToUnit(f, m)
	case "ESCAPE":
		GAME_IS_RUNNING = false
		PLR_LOOP = false
	default:
		plr_moveCursor(f, keyPressed)
	}
}

func plr_giveAnOrderToUnit(f *faction, m *gameMap) {
	u := m.getUnitAtCoordinates(f.cx, f.cy)
	if u == nil {
		// log.appendMessage("SELECTED NIL")
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
			return
		case "ESCAPE":
			return
		default:
			plr_moveCursor(f, keyPressed)
		}
	}
}


func plr_moveCursor(f *faction, keyPressed string) {
	vx, vy := plr_keyToDirection(keyPressed)
	if areCoordsValid(f.cx+vx, f.cy+vy) {
		f.cx += vx
		f.cy += vy
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