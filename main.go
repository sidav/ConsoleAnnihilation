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

	//for i:=0; i<1024; i++ {
	//	cw.PutChar(int32(i), i%80, i/80)
	//}
	//cw.Flush_console()
	//for key:=""; key != "ESCAPE"; {
	//	key = cw.ReadKey()
	//}

	log = &LOG{}

	for GAME_IS_RUNNING {
		CURRENT_TURN += 1
		for _, f := range a.factions {
			if !GAME_IS_RUNNING {
				return 
			}
			renderFactionStats(f)
			plr_control(f, a)
		}
	}

}
