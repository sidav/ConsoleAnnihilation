package main

import cw "TCellConsoleWrapper/tcell_wrapper"

func areCoordsValid(x, y int) bool {
	return (x >= 0) && (x < mapW) && (y >= 0) && (y < mapH)
}

var (
	GAME_IS_RUNNING = true
	log *LOG
	CURRENT_TURN = 0
)

func main() {
	cw.Init_console()
	defer cw.Close_console()

	a := &gameMap{}
	a.init()

	log = &LOG{}

	for GAME_IS_RUNNING {
		CURRENT_TURN += 1
		for _, f := range a.factions {
			renderFactionStats(f)
			plr_control(f, a)
		}
	}

}
