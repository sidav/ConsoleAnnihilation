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
	r_renderMapAroundCursor(g, f.cursor.x, f.cursor.y)
	renderFactionStats(f)
	renderInfoOnCursor(f, g)
	r_renderCursor(f)
	r_renderUIOutline(f)
	flushView()
}

func r_renderUIOutline(f *faction) {
	cw.SetFgColor(f.getFactionColor())
	for y := 0; y < VIEWPORT_H; y++ {
		cw.PutChar('|', VIEWPORT_W, y)
	}
	for x := 0; x < CONSOLE_W; x++ {
		cw.PutChar('-', x, VIEWPORT_H)
	}
	cw.PutChar('+', VIEWPORT_W, VIEWPORT_H)
	cw.SetBgColor(cw.BLACK)
}

func r_renderMapAroundCursor(g *gameMap, cx, cy int) {
	cw.Clear_console()
	vx := cx - VIEWPORT_W/2
	vy := cy - VIEWPORT_H/2
	renderMapInViewport(g, vx, vy)
	renderPawnsInViewport(g, vx, vy)
	renderLog(false)
}

func renderMapInViewport(g *gameMap, vx, vy int) {
	for x := vx; x < vx+VIEWPORT_W; x++ {
		for y := vy; y < vy+VIEWPORT_H; y++ {
			if areCoordsValid(x, y) {
				tileApp := g.tileMap[x][y].appearance
				r_setFgColorByCcell(tileApp)
				cw.PutChar(tileApp.char, x-vx, y-vy)
			}
		}
	}
}

func renderPawnsInViewport(g *gameMap, vx, vy int) {
	for _, p := range g.pawns {
		if p.isBuilding() {
			renderBuildingsInViewport(p, g, vx, vy)
		} else {
			renderUnitsInViewport(p, g, vx, vy)
		}
	}
}

func renderUnitsInViewport(p *pawn, g *gameMap, vx, vy int) {
	u := p.unitInfo
	if areGlobalCoordsOnScreen(p.x, p.y, vx, vy) {
		tileApp := u.appearance
		// r, g, b := getFactionRGB(u.faction.factionNumber)
		// cw.SetFgColorRGB(r, g, b)
		cw.SetFgColor(p.faction.getFactionColor())
		cw.PutChar(tileApp.char, p.x-vx, p.y-vy)
	}
}

func renderBuildingsInViewport(p *pawn, g *gameMap, vx, vy int) {
	b := p.buildingInfo
	app := b.appearance
	bx, by := p.getCoords()
	for x := 0; x < b.w; x++ {
		for y := 0; y < b.h; y++ {
			if p.currentConstructionStatus == nil {
				color := app.colors[x+b.w*y]
				if color == -1 {
					cw.SetFgColor(p.faction.getFactionColor())
				} else {
					cw.SetFgColor(color)
				}
			} else { // building is under construction
				color := 2
				cw.SetFgColor(color)
			}
			if areGlobalCoordsOnScreen(bx+x, by+y, vx, vy) {
				cw.PutChar(int32(app.chars[x+b.w*y]), bx+x-vx, by+y-vy)
			}
		}
	}
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

func r_renderPossibleOrdersForPawn(p *pawn) {
	orders := make([]string, 0)
	if p.canConstructBuildings() {
		orders = append(orders, "(B)uild")
	}
	if p.canConstructUnits() {
		orders = append(orders, "(C)onstruct units")
	}
	if p.hasWeapons() {
		orders = append(orders, "(A)ttack-move")
	}
	routines.DrawSidebarInfoMenu("Orders for: "+p.name, p.faction.getFactionColor(),
		SIDEBAR_X, SIDEBAR_FLOOR_3, SIDEBAR_W, orders)
}

func flushView() {
	cw.Flush_console()
}

func areGlobalCoordsOnScreen(gx, gy, vx, vy int) bool {
	return areCoordsInRect(gx, gy, vx, vy, VIEWPORT_W, VIEWPORT_H)
}

func areGlobalCoordsOnScreenForFaction(gx, gy int, f *faction) bool {
	vx := f.cursor.x - VIEWPORT_W/2
	vy := f.cursor.y - VIEWPORT_H/2
	return areCoordsInRect(gx, gy, vx, vy, VIEWPORT_W, VIEWPORT_H)
}
