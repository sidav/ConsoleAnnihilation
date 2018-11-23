package main

import (
	"SomeTBSGame/routines"
	cw "TCellConsoleWrapper/tcell_wrapper"
	"fmt"
)

const (
	CONSOLE_W  = 80
	CONSOLE_H  = 25
	VIEWPORT_W = 40
	VIEWPORT_H = 20
	SIDEBAR_X  = VIEWPORT_W + 1
	SIDEBAR_W  = CONSOLE_W - VIEWPORT_W - 1
)

func r_setFgColorByCcell(c *ccell) {
	cw.SetFgColor(c.color)
	// cw.SetFgColorRGB(c.r, c.g, c.b)
}

func r_renderScreenForFaction(f *faction, g *gameMap) {
	r_renderMapAroundCursor(g, f.cursor.x, f.cursor.y)
	renderFactionStats(f)
	renderInfoOnCursor(f, g)
	flushView()
}

func r_renderMapAroundCursor(g *gameMap, cx, cy int) {
	cw.Clear_console()
	vx := cx - VIEWPORT_W/2
	vy := cy - VIEWPORT_H/2
	renderMapInViewport(g, vx, vy)
	renderBuildingsInViewport(g, vx, vy)
	renderUnitsInViewport(g, vx, vy)
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

func renderUnitsInViewport(g *gameMap, vx, vy int) {
	for _, u := range g.units {
		if areGlobalCoordsOnScreen(u.x, u.y, vx, vy) {
			tileApp := u.appearance
			// r, g, b := getFactionRGB(u.faction.factionNumber)
			// cw.SetFgColorRGB(r, g, b)
			cw.SetFgColor(getFactionColor(u.faction.factionNumber))
			cw.PutChar(tileApp.char, u.x-vx, u.y-vy)
		}
	}

}

func renderBuildingsInViewport(g *gameMap, vx, vy int) {
	for _, b := range g.buildings {
		app := b.appearance
		bx, by := b.getCoords()
		for x := 0; x < b.w; x++ {
			for y := 0; y < b.h; y++ {
				if b.currentConstructionStatus == nil {
					color := app.colors[x+b.w*y]
					if color == -1 {
						cw.SetFgColor(getFactionColor(b.faction.factionNumber))
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

}

func renderInfoOnCursor(f *faction, g *gameMap) {

	if f.cursor.snappedBuilding != nil {
		b := f.cursor.snappedBuilding
		str := make([]string, 0)
		if b.faction != f {
			str = append(str, "(Enemy building)")
		} else {
			str = append(str, "Your building, Commander")
		}
		routines.DrawSidebarInfoMenu(b.name, getFactionColor(b.faction.factionNumber), SIDEBAR_X, 7, SIDEBAR_W, str)
		return
	}

	cx, cy := f.cursor.getCoords()
	u := g.getUnitAtCoordinates(cx, cy)
	if u != nil {
		str := make([]string, 0)
		if u.faction != f {
			str = append(str, "(Enemy unit)")
		} else {
			str = append(str, "Your loyal unit, Commander")
		}
		routines.DrawSidebarInfoMenu(u.name, getFactionColor(u.faction.factionNumber), SIDEBAR_X, 7, SIDEBAR_W, str)
		return
	}
}

func renderSelectCursor() {
	x := VIEWPORT_W / 2
	y := VIEWPORT_H / 2
	// cw.SetFgColorRGB(128, 128, 128)
	cw.SetFgColor(cw.WHITE)

	cw.PutChar('[', x-1, y)
	cw.PutChar(']', x+1, y)

	// outcommented for non-SDL console
	//cw.PutChar(16*13+10, x-1, y-1)
	//cw.PutChar(16*11+15, x+1, y-1)
	//cw.PutChar(16*12, x-1, y+1)
	//cw.PutChar(16*13+9, x+1, y+1)
	flushView()
}

func renderMoveCursor() {
	x := VIEWPORT_W / 2
	y := VIEWPORT_H / 2
	// cw.SetFgColorRGB(128, 255, 128)
	cw.SetFgColor(cw.GREEN)

	cw.PutChar('>', x-1, y)
	cw.PutChar('<', x+1, y)

	//cw.PutChar('\\', x-1, y-1)
	//cw.PutChar('/', x+1, y-1)
	//cw.PutChar('/', x-1, y+1)
	//cw.PutChar('\\', x+1, y+1)

	flushView()
}

func renderFactionStats(f *faction) {
	eco := f.economy
	statsx := VIEWPORT_W + 1

	// fr, fg, fb := getFactionRGB(f.factionNumber)
	// cw.SetFgColorRGB(fr, fg, fb)
	cw.SetFgColor(getFactionColor(f.factionNumber))
	cw.PutString(fmt.Sprintf("%s: turn %d", f.name, CURRENT_TURN/10+1), statsx, 0)

	metal, maxmetal := eco.currentMetal, eco.maxMetal
	cw.SetFgColor(cw.DARK_CYAN)
	renderStatusbar("METAL", metal, maxmetal, statsx, 1, CONSOLE_W-VIEWPORT_W-3, cw.DARK_CYAN)
	cw.SetFgColor(cw.DARK_CYAN)
	metalDetails := fmt.Sprintf("%d/%d stored  %d (+%d / -%d) per turn",eco.currentMetal, eco.maxMetal,
		eco.metalIncome - eco.metalSpending, eco.metalIncome, eco.metalSpending)
	cw.PutString(metalDetails, statsx, 2)

	energy, maxenergy := eco.currentEnergy, eco.maxEnergy
	cw.SetFgColor(cw.DARK_YELLOW)
	renderStatusbar("ENERGY", energy, maxenergy, statsx, 4, CONSOLE_W-VIEWPORT_W-3, cw.DARK_YELLOW)
	cw.SetFgColor(cw.DARK_YELLOW)
	energyDetails := fmt.Sprintf("%d/%d stored  %d (+%d / -%d) per turn",eco.currentEnergy, eco.maxEnergy,
		eco.energyIncome - eco.energySpending, eco.energyIncome, eco.energySpending)
	cw.PutString(energyDetails, statsx, 5)
}

func renderStatusbar(name string, curvalue, maxvalue, x, y, width, barColor int) {
	barTitle := name
	cw.PutString(barTitle, x, y)
	barWidth := width - len(name)
	filledCells := barWidth * curvalue / maxvalue
	barStartX := x + len(barTitle) + 1
	for i := 0; i < barWidth; i++ {
		if i < filledCells {
			cw.SetFgColor(barColor)
			cw.PutChar('=', i+barStartX, y)
		} else {
			cw.SetFgColor(cw.DARK_BLUE)
			cw.PutChar('-', i+barStartX, y)
		}
	}
}

func renderLog(flush bool) {
	cw.SetFgColor(cw.WHITE)
	for i := 0; i < LOG_HEIGHT; i++ {
		cw.PutString(log.last_msgs[i].getText(), 0, VIEWPORT_H+i)
	}
	if flush {
		flushView()
	}
}

func flushView() {
	cw.Flush_console()
}

func areGlobalCoordsOnScreen(gx, gy, vx, vy int) bool {
	return areCoordsInRect(gx, gy, vx, vy, VIEWPORT_W, VIEWPORT_H)
}
