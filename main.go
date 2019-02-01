package main

import (
	"SomeTBSGame/routines"
	cw "TCellConsoleWrapper"
)

func areCoordsValid(x, y int) bool {
	return (x >= 0) && (x < mapW) && (y >= 0) && (y < mapH)
}

func areCoordsInRect(x, y, rx, ry, w, h int) bool {
	return x >= rx && x < rx+w && y >= ry && y < ry+h
}

func areCoordsInRange(fx, fy, tx, ty, r int) bool {
	return (fx-tx)*(fx-tx) + (fy-ty)*(fy-ty) <= r*r
}

func getSqDistanceBetween(x1, y1, x2, y2 int) int {
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}

var (
	GAME_IS_RUNNING = true
	log             *LOG
	CURRENT_TURN    = 0
	CURRENT_MAP     *gameMap
	CURRENT_FACTION_SEEING_THE_SCREEN *faction // for various rendering crap
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

	showBriefing()

	for GAME_IS_RUNNING {
		for _, f := range CURRENT_MAP.factions {
			if !GAME_IS_RUNNING {
				return
			}
			if f.playerControlled {
				CURRENT_FACTION_SEEING_THE_SCREEN = f
				renderFactionStats(f)
				plr_control(f, CURRENT_MAP)
			}
		}
		for i := 0; i < 10; i++ {
			for _, u := range CURRENT_MAP.pawns {
				if u.hitpoints <= 0 {
					log.appendMessage(u.name + " is destroyed!")
					CURRENT_MAP.removePawn(u)
					continue
				}
				u.executeOrders(CURRENT_MAP)
				u.openFireIfPossible()
			}
			CURRENT_TURN += 1
		}

		for _, f := range CURRENT_MAP.factions {
			f.recalculateFactionEconomy(CURRENT_MAP)
		}
		doAllNanolathes(CURRENT_MAP)
	}

}

func showBriefing() {
	text := "Good day, Head Officer #CC-42, and welcome to Thalassean-3. \\n \\n " +
		"That \"Arm\" rebellion was quite of a " +
		"surprise for command, and our forces were drawn away from this planet by surprise attack. Our intel suggested " +
		"that there is a proxy Arm base on the surface. Three hours ago we've confirmed its coordinates. \\n \\n " +
		"There are no corporate forces at the surface for now, but we've found an ancient Commander prototype unit lying in " +
		"conservation from the days of these armored command units development. Energy signatures show that the unit is " +
		"still functional, although you have to keep in mind that this is only an unfinished prototype, so it has no " +
		"nanolathe schematics or Disintegrator Gun. We are starting the process of uploading your brain neural masks " +
		"patterns to the gel neocortex of that machine. \\n \\n " +
		"Your prime directive is to destroy an enemy HQ building, which is supposedly holding an AI patterns to be " +
		"activated by enemy at some point in time. We're thinking that destruction of such a building will immobilize whole Arm vermin at the " +
		"planet. \\n " +
		"As we speak, our data transfer relays are finishing uploading some basic scematics for the prototype's nanolating " +
		"equipment. That means that you are clear to embark right now. \\n \\n " +
		"You will be dispatched immediately. "
	routines.DrawWrappedTextInRect(text, 0, 0, CONSOLE_W, CONSOLE_H)
	key := ""
	for key != "ESCAPE" && key != "ENTER" {
		key = cw.ReadKey()
	}
}
