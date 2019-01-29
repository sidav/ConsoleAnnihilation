package main

import cw "TCellConsoleWrapper"

func areCoordsValid(x, y int) bool {
	return (x >= 0) && (x < mapW) && (y >= 0) && (y < mapH)
}

func areCoordsInRect(x, y, rx, ry, w, h int) bool {
	return x >= rx && x < rx+w && y >= ry && y < ry+h
}

func getSqDistanceBetween(x1, y1, x2, y2 int) int {
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}

var (
	GAME_IS_RUNNING = true
	log             *LOG
	CURRENT_TURN    = 0
	CURRENT_MAP     *gameMap
)

func main() {
	cw.Init_console()
	defer cw.Close_console()

	CURRENT_MAP = &gameMap{}
	CURRENT_MAP.init()

	//for i:=0; i<1024; i++ {
	//	cw.PutChar(int32(i), i%80, i/80)
	//}
	//cw.Flush_console()
	//for key:=""; key != "ESCAPE"; {
	//	key = cw.ReadKey()
	//}

	log = &LOG{}

	for GAME_IS_RUNNING {
		for _, f := range CURRENT_MAP.factions {
			if !GAME_IS_RUNNING {
				return
			}
			if f.playerControlled {
				renderFactionStats(f)
				plr_control(f, CURRENT_MAP)
			}
		}
		for i := 0; i < 10; i++ {
			for _, u := range CURRENT_MAP.pawns {
				u.executeOrders(CURRENT_MAP)
			}
			CURRENT_TURN += 1
		}

		for _, f := range CURRENT_MAP.factions {
			f.recalculateFactionEconomy(CURRENT_MAP)
		}
		doAllNanolathes(CURRENT_MAP)
	}

}
