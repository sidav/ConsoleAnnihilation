package main

import (
	"SomeTBSGame/routines"
	cw "TCellConsoleWrapper"
	"fmt"
)

var (
	CONSOLE_W, CONSOLE_H = 80, 25
	VIEWPORT_W           = 40
	VIEWPORT_H           = CONSOLE_H - LOG_HEIGHT
	SIDEBAR_X            = VIEWPORT_W + 1
	SIDEBAR_W            = CONSOLE_W - VIEWPORT_W - 1
	SIDEBAR_H            = CONSOLE_H - LOG_HEIGHT
	SIDEBAR_FLOOR_2      = 7  // y-coord right below resources info
	SIDEBAR_FLOOR_3      = 11 // y-coord right below "floor 2"
)

func r_setFgColorByCcell(c *ccell) {
	cw.SetFgColor(c.color)
	// cw.SetFgColorRGB(c.r, c.g, c.b)
}

func r_updateBoundsIfNeccessary() {
	if cw.WasResized() {
		CONSOLE_W, CONSOLE_H = cw.GetConsoleSize()
		VIEWPORT_W           = cw.CONSOLE_WIDTH / 2
		VIEWPORT_H           = CONSOLE_H - LOG_HEIGHT - 1
		SIDEBAR_X            = VIEWPORT_W + 1
		SIDEBAR_W            = CONSOLE_W - VIEWPORT_W - 1
		SIDEBAR_H            = CONSOLE_H - LOG_HEIGHT
		SIDEBAR_FLOOR_2      = 7  // y-coord right below resources info
		SIDEBAR_FLOOR_3      = 11 // y-coord right below "floor 2"
	}
}

func r_renderScreenForFaction(f *faction, g *gameMap) {
	r_updateBoundsIfNeccessary()
	cx, cy := f.cursor.x, f.cursor.y
	cw.Clear_console()
	vx := cx - VIEWPORT_W/2
	vy := cy - VIEWPORT_H/2
	renderMapInViewport(f, g, vx, vy)
	renderFactionStats(f)
	renderInfoOnCursor(f, g)
	r_renderUIOutline(f)
	renderPawnsInViewport(f, g, vx, vy)
	r_renderCursor(f)
	renderLog(false)
	flushView()
}

func r_renderUIOutline(f *faction) {
	if IS_PAUSED {
		cw.SetBgColor(f.getFactionColor())
	} else {
		cw.SetFgColor(f.getFactionColor())
	}
	for y := 0; y < VIEWPORT_H; y++ {
		cw.PutChar('|', VIEWPORT_W, y)
	}
	for x := 0; x < CONSOLE_W; x++ {
		cw.PutChar('-', x, VIEWPORT_H)
	}
	cw.PutChar('+', VIEWPORT_W, VIEWPORT_H)
	if IS_PAUSED {
		cw.SetBgColor(cw.BLACK)
		cw.SetFgColor(cw.YELLOW)
		cw.PutString("TACTICAL PAUSE", VIEWPORT_W / 2 - 7, VIEWPORT_H)
	}
	cw.SetBgColor(cw.BLACK)
}

func renderMapInViewport(f *faction, g *gameMap, vx, vy int) {
	for x := vx; x < vx+VIEWPORT_W; x++ {
		for y := vy; y < vy+VIEWPORT_H; y++ {
			if areCoordsValid(x, y) {
				if f.wereCoordsSeen(x, y) {
					tileApp := g.tileMap[x][y].appearance
					if f.areCoordsInSight(x, y) {
						r_setFgColorByCcell(tileApp)
					} else {
						cw.SetFgColor(cw.DARK_BLUE)
					}
					cw.PutChar(tileApp.char, x-vx, y-vy)
				} else {
					cw.PutChar(' ', x-vx, y-vy)
				}
			}
		}
	}
}

func renderPawnsInViewport(f *faction, g *gameMap, vx, vy int) {
	for _, p := range g.pawns {
		cx, cy := p.getCenter()
		if f.areCoordsInRadarRadius(cx, cy) {
			cw.SetFgColor(cw.RED)
			renderCharByGlobalCoords('?', cx,cy)
		}
		if p.isBuilding() {
			renderBuilding(f, p, g, vx, vy, false)
		} else {
			renderUnit(f, p, g, vx, vy, false)
		}
	}
}

func r_renderSelectedPawns(f* faction, selection *[]*pawn) {
	vx := f.cursor.x - VIEWPORT_W/2
	vy := f.cursor.y - VIEWPORT_H/2
	for _, p := range *selection {
		if p.isUnit() {
			renderUnit(f, p, CURRENT_MAP, vx, vy, true)
		} else if p.isBuilding() {
			renderBuilding(f, p, CURRENT_MAP, vx, vy, true)
		}
	}
}

func renderUnit(f *faction, p *pawn, g *gameMap, vx, vy int, inverse bool) {
	u := p.unitInfo
	if areGlobalCoordsOnScreen(p.x, p.y, vx, vy) && f.areCoordsInSight(p.x, p.y){
		tileApp := u.appearance
		// r, g, b := getFactionRGB(u.faction.factionNumber)
		// cw.SetFgColorRGB(r, g, b)\
		colorToRender := p.faction.getFactionColor()
		if inverse {
			cw.SetBgColor(colorToRender)
			cw.SetFgColor(cw.BLACK)
		} else {
			cw.SetFgColor(colorToRender)
		}
		cw.PutChar(tileApp.char, p.x-vx, p.y-vy)
		cw.SetBgColor(cw.BLACK)
	}
}

func renderBuilding(f *faction, p *pawn, g *gameMap, vx, vy int, inverse bool) {
	b := p.buildingInfo
	app := b.appearance
	bx, by := p.getCoords()
	colorToRender := 0
	for x := 0; x < b.w; x++ {
		for y := 0; y < b.h; y++ {
			if p.currentConstructionStatus == nil {
				color := app.colors[x+b.w*y]
				if f.areCoordsInSight(bx+x,by+y) {
					if color == -1 {
						colorToRender = p.faction.getFactionColor()
					} else {
						colorToRender = color
					}
				} else {
					colorToRender = cw.DARK_BLUE
				}
			} else { // building is under construction
				colorToRender = cw.DARK_GREEN
				if CURRENT_TURN/10 % 2 == 0 {
					colorToRender = cw.GREEN
				}
			}
			if areGlobalCoordsOnScreen(bx+x, by+y, vx, vy) && f.wereCoordsSeen(bx+x, by+y) {
				if inverse {
					cw.SetBgColor(colorToRender)
					cw.SetFgColor(cw.BLACK)
				} else {
					cw.SetFgColor(colorToRender)
				}
				cw.PutChar(int32(app.chars[x+b.w*y]), bx+x-vx, by+y-vy)
			}
		}
	}
	cw.SetBgColor(cw.BLACK)
}

func renderInfoOnCursor(f *faction, g *gameMap) {

	title := "nothing"
	color := 2
	details := make([]string, 0)
	var res *pawnResourceInformation
	sp := f.cursor.snappedPawn

	if sp != nil {

		renderOrderLine(sp)

		color = sp.faction.getFactionColor()
		title = sp.name
		if sp.faction != f {
			if sp.isBuilding() {
				details = append(details, "(Enemy building)")
			} else {
				details = append(details, "(Enemy unit)")
			}
		} else {
			details = append(details, sp.getCurrentOrderDescription())
			if sp.res != nil && sp.currentConstructionStatus == nil {
				res = sp.res
			}
		}
		r_renderAttackRadius(sp)
	}

	if len(details) > 0 {
		details = append(details, sp.getArmorDescriptionString())
		if sp.hasWeapons() {
			for _, wpn := range sp.weapons {
				details = append(details, wpn.getDescriptionString())
			}
		}
		if res != nil {
			economyInfo := fmt.Sprintf("METAL: (+%d / -%d) ENERGY: (+%d / -%d)",
				res.metalIncome, res.metalSpending, res.energyIncome, res.energySpending+res.energyReqForConditionalMetalIncome)
			details = append(details, economyInfo)
		}
		routines.DrawSidebarInfoMenu(title, color, SIDEBAR_X, SIDEBAR_FLOOR_2, SIDEBAR_W, details)
	}
}

func flushView() {
	cw.Flush_console()
}

func renderCharByGlobalCoords(c rune, x, y int) { // TODO: use it everywhere
	vx := CURRENT_FACTION_SEEING_THE_SCREEN.cursor.x - VIEWPORT_W/2
	vy := CURRENT_FACTION_SEEING_THE_SCREEN.cursor.y - VIEWPORT_H/2
	if areGlobalCoordsOnScreenForFaction(x, y, CURRENT_FACTION_SEEING_THE_SCREEN) {
		cw.PutChar(c, x-vx, y-vy)
	}
}

func areGlobalCoordsOnScreen(gx, gy, vx, vy int) bool {
	return routines.AreCoordsInRect(gx, gy, vx, vy, VIEWPORT_W, VIEWPORT_H)
}

func areGlobalCoordsOnScreenForFaction(gx, gy int, f *faction) bool {
	vx := f.cursor.x - VIEWPORT_W/2
	vy := f.cursor.y - VIEWPORT_H/2
	return routines.AreCoordsInRect(gx, gy, vx, vy, VIEWPORT_W, VIEWPORT_H)
}
