package main

import cw "GoSdlConsole/GoSdlConsole"

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
		for _, f := range a.factions {
			renderFactionStats(f)
			CURRENT_TURN += 1
			log.appendMessagef("TURN %d", (CURRENT_TURN))
			plr_control(f, a)
		}
	}

}
