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
	r_renderCursor(f)
	r_renderUIOutline(f)
	flushView()
}

func r_renderUIOutline(f *faction) {
	cw.SetFgColor(f.getFactionColor())
	for y := 0; y < VIEWPORT_H; y++ {
		cw.PutChar('|', VIEWPORT_W, y)
	}
	for x:=0; x < CONSOLE_W; x++ {
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
		b:= p.buildingInfo
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
			if sp.res != nil {
				res = sp.res
			}
		}
	}

	if len(details) > 0 {
		if res != nil {
			economyInfo := fmt.Sprintf("METAL: (+%d / -%d) ENERGY: (+%d / -%d)",
				res.metalIncome, res.metalSpending, res.energyIncome, res.energySpending+res.energyReqForConditionalMetalIncome)
			details = append(details, economyInfo)
		}
		routines.DrawSidebarInfoMenu(title, color, SIDEBAR_X, 7, SIDEBAR_W, details)
	}
}

//func r_renderPawnPossibleOrders(p *pawn) {
//	orders := make([]string, 0)
//	if p.canBuild() {
//		orders = append(orders, "(B)uild")
//	}
//	// routines.DrawSidebarInfoMenu(p.name, p.faction.getFactionColor(), )// move to render?
//}

func renderFactionStats(f *faction) {
	eco := f.economy
	statsx := VIEWPORT_W + 1

	// fr, fg, fb := getFactionRGB(f.factionNumber)
	// cw.SetFgColorRGB(fr, fg, fb)
	cw.SetFgColor(f.getFactionColor())
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
	var filledCells int
	if maxvalue > 0 {
		filledCells = barWidth * curvalue / maxvalue
	} else {
		filledCells = 0
	}
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
