package main

import cw "GoSdlConsole/GoSdlConsole"

func areCoordsValid(x, y int) bool {
	return (x >= 0) && (x < mapW) && (y >= 0) && (y < mapH)
}

var (
	GAME_IS_RUNNING = true
	log *LOG
)

func main() {
	cw.Init_console()
	defer cw.Close_console()

	a := &gameMap{}
	a.init()

	log = &LOG{}

	f := &faction{factionNumber: 0}

	for GAME_IS_RUNNING {
		// r_renderMapAroundCursor(a, 0, 0)
		plr_control(f, a)
	}

}
