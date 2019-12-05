package main

import (
	geometry "github.com/sidav/golibrl/geometry"
	cmenu "github.com/sidav/golibrl/console_menu"
	cw "github.com/sidav/golibrl/console"
	"strconv"
	"time"
)

func areCoordsValid(x, y int) bool {
	return geometry.AreCoordsInRect(x, y, 0, 0, mapW, mapH)
}

var (
	GAME_IS_RUNNING                     = true
	log                                 *LOG
	CURRENT_TICK                        = 1
	CURRENT_MAP                         *gameMap
	CURRENT_FACTION_SEEING_THE_SCREEN   *faction // for various rendering crap
	FIRE_WAS_OPENED_ON_SCREEN_THIS_TURN bool     // for killing pewpews overrender.
	CHEAT_IGNORE_FOW                    bool
	DEBUG_OUTPUT                        bool
)

func getCurrentTurn() int {
	return CURRENT_TICK/10 + 1
}

func debug_write(text string) {
	if DEBUG_OUTPUT {
		log.appendMessage("DEBUG: " + text)
	}
}

func main() {
	cw.Init_console("Console Annihilation", cw.TCellRenderer)
	defer cw.Close_console()

	log = &LOG{}
	initSquadMembersStaticDataMap()
	initBuildingsStaticDataMap()

	CURRENT_MAP = &gameMap{}
	CURRENT_MAP.init()
	r_updateBoundsIfNeccessary(true)

	///////////////////////////////
	// uncomment later
	//r_showTitleScreen()
	//showBriefing()
	// comment later
	// endTurnPeriod = 0
	///////////////////////////////////

	for {
		startTime := time.Now()
		for _, f := range CURRENT_MAP.factions {
			f.recalculateSeenTiles()
			checkWinOrLose()
			if !GAME_IS_RUNNING {
				return
			}
			if f.aiControlled {
				// ai_controlFaction(f)
			}
			if f.playerControlled {
				CURRENT_FACTION_SEEING_THE_SCREEN = f
				renderFactionStats(f)
				plr_control(f, CURRENT_MAP)
				debug_write("TOTAL FLUSHES: " + strconv.Itoa(cw.GetNumberOfRecentFlushes()))
			}
		}
		for i := 0; i < 10; i++ {
			for _, u := range CURRENT_MAP.pawns {
				if !u.isAlive() {
					log.appendMessage(u.getName() + " is destroyed!")
					CURRENT_MAP.removePawn(u)
					continue
				}
				if u.regenPeriod > 0 && CURRENT_TICK%u.regenPeriod == 0 && u.hitpoints < u.maxHitpoints {
					u.hitpoints++
				}
				u.executeOrders(CURRENT_MAP)
				u.openFireIfPossible()
			}
			if FIRE_WAS_OPENED_ON_SCREEN_THIS_TURN {
				cw.Flush_console()
				FIRE_WAS_OPENED_ON_SCREEN_THIS_TURN = false
				time.Sleep(time.Duration(endTurnPeriod/4) * time.Millisecond)
			}
			CURRENT_TICK += 1
		}

		for _, f := range CURRENT_MAP.factions {
			f.recalculateFactionEconomy(CURRENT_MAP)
		}
		doAllNanolathes(CURRENT_MAP)
		timeForTurn := int(time.Since(startTime) / time.Millisecond)
		debug_write("Time for turn: " + strconv.Itoa(timeForTurn) + "ms") // TODO: make it removable
	}

}

func showBriefing() {
	cw.Clear_console()
	text := "Good day, Head Officer #CC-42, and welcome to Thalassean-3. \\n " +
		"That \"Arm\" rebellion was quite of a " +
		"surprise for command, and our forces were drawn away from this planet by surprise attack. Our intel suggested " +
		"that there is a proxy Arm base on the surface. Three hours ago we've confirmed its coordinates. \\n " +
		"There are no corporate forces at the surface for now, but we've found an ancient Commander prototype unit lying in " +
		"conservation from the days of these armored command units development. Energy signatures show that the unit is " +
		"still functional, although you have to keep in mind that this is only an unfinished prototype, so it has no " +
		"nanolathe schematics or Disintegrator Gun. We are starting the process of uploading your brain neural masks " +
		"patterns to the gel neocortex of that machine. \\n " +
		"Your prime directives are the following: \\n " +
		"First, you are to destroy an enemy HQ building, which is supposedly holding an AI patterns to be " +
		"activated by enemy at some point in time. We're thinking that destruction of such a building will immobilize whole Arm vermin at the " +
		"planet. \\n " +
		"Second, you are to keep that prototype Commander unit operational. After we deal with Arm forces on that planet, " +
		"that prototype machine would prove itself useful for continuing the suppression of rebellion on another planets. " +
		"These ACUs are extremely expensive and complex in production, and the Corporation cannot afford producing more " +
		"right now. \\n " +
		"As we speak, our data transfer relays are finishing uploading some basic scematics for the prototype's nanolating " +
		"equipment. That means that you are clear to embark right now. \\n " +
		"You will be dispatched immediately. "
	cmenu.DrawWrappedTextInRect(text, 0, 0, CONSOLE_W, CONSOLE_H)
	cw.Flush_console()
	key := ""
	for key != "ESCAPE" && key != "ENTER" {
		key = cw.ReadKey()
	}
}
