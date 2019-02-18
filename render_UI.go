package main

import (
	"SomeTBSGame/routines"
	"fmt"
)
import cw "TCellConsoleWrapper"

func renderFactionStats(f *faction) {
	eco := f.economy
	statsx := VIEWPORT_W + 1

	// fr, fg, fb := getFactionRGB(f.factionNumber)
	// cw.SetFgColorRGB(fr, fg, fb)
	if IS_PAUSED {
		cw.SetFgColor(f.getFactionColor())
		cw.PutString(f.name+": ", statsx, 0)
		cw.SetFgColor(cw.YELLOW)
		cw.PutString(fmt.Sprintf("turn %d (PAUSED)", CURRENT_TURN/10+1), statsx+len(f.name)+2, 0)
	} else {
		cw.SetFgColor(f.getFactionColor())
		cw.PutString(fmt.Sprintf("%s: turn %d", f.name, CURRENT_TURN/10+1), statsx, 0)
	}

	metal, maxmetal, metalFlow := eco.currentMetal, eco.maxMetal, eco.metalIncome-eco.metalSpending
	cw.SetFgColor(cw.DARK_CYAN)
	renderStatusbar("METAL", metal, maxmetal, statsx, 1, CONSOLE_W-VIEWPORT_W-3, cw.DARK_CYAN)
	cw.SetFgColor(cw.DARK_CYAN)
	metalDetails := fmt.Sprintf("%d/%d stored  %d (+%d / -%d) per turn", eco.currentMetal, eco.maxMetal,
		metalFlow, eco.metalIncome, eco.metalSpending)
	cw.PutString(metalDetails, statsx, 2)

	if metal+metalFlow < 0 {
		cw.SetFgColor(cw.RED)
		cw.PutString("STALL", statsx+SIDEBAR_W/3, 1)
	}

	energy, maxenergy, energyFlow := eco.currentEnergy, eco.maxEnergy, eco.energyIncome-eco.energySpending
	cw.SetFgColor(cw.DARK_YELLOW)
	renderStatusbar("ENERGY", energy, maxenergy, statsx, 4, CONSOLE_W-VIEWPORT_W-3, cw.DARK_YELLOW)
	cw.SetFgColor(cw.DARK_YELLOW)
	energyDetails := fmt.Sprintf("%d/%d stored  %d (+%d / -%d) per turn", eco.currentEnergy, eco.maxEnergy,
		energyFlow, eco.energyIncome, eco.energySpending)
	cw.PutString(energyDetails, statsx, 5)

	if energy+energyFlow < 0 {
		cw.SetFgColor(cw.RED)
		cw.PutString("STALL", statsx+SIDEBAR_W/3, 4)
	}
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
		cw.PutString(log.last_msgs[i].getText(), 0, VIEWPORT_H+i+1)
	}
	if flush {
		flushView()
	}
}

func r_renderAttackRadius(p *pawn) {
	if len(p.weapons) > 0 {
		// (p.x, p.y, p.weapons[0].attackRadius, false, CURRENT_FACTION_SEEING_THE_SCREEN.cursor.x-VIEWPORT_W/2, CURRENT_FACTION_SEEING_THE_SCREEN.cursor.y-VIEWPORT_H/2)
		vx, vy := CURRENT_FACTION_SEEING_THE_SCREEN.cursor.x-VIEWPORT_W/2, CURRENT_FACTION_SEEING_THE_SCREEN.cursor.y-VIEWPORT_H/2
		line := routines.GetCircle(p.x, p.y, p.weapons[0].attackRadius)
		for _, point := range *line {
			x, y := point.X, point.Y
			if areCoordsInRect(x-vx, y-vy, 0, 0, VIEWPORT_W, VIEWPORT_H) {
				cw.PutChar('X', x-vx, y-vy)
			}
		}
	}
}

func renderOrderLine(p *pawn) {
	var ordr *order
	if p.order != nil {
		if p.order.canBeDrawnAsLine() {
			ordr = p.order
		}
	}
	if ordr == nil {
		if p.canConstructUnits() {
			ordr = p.nanolatherInfo.defaultOrderForUnitBuilt
		}
	}
	if ordr != nil {
		if ordr.orderType == order_attack {
			cw.SetFgColor(cw.RED)
		} else {
			cw.SetFgColor(cw.YELLOW)
		}
		f := p.faction
		cx, cy := p.getCenter()
		renderLine(cx, cy, ordr.x, ordr.y, false, f.cursor.x-VIEWPORT_W/2, f.cursor.y-VIEWPORT_H/2)
	}
}

func renderLine(fromx, fromy, tox, toy int, flush bool, vx, vy int) {
	line := routines.GetLine(fromx, fromy, tox, toy)
	char := '?'
	if len(line) > 1 {
		dirVector := routines.CreateVectorByStartAndEndInt(fromx, fromy, tox, toy)
		dirVector.TransformIntoUnitVector()
		dirx, diry := dirVector.GetRoundedCoords()
		char = getTargetingChar(dirx, diry)
	}
	//if fromx == tox && fromy == toy {
	//	renderPawn(d.player, true)
	//}
	for i := 1; i < len(line); i++ {
		// x, y := line[i].X, line[i].Y
		//if d.isPawnPresent(x, y) {
		//	renderPawn(d.getPawnAt(x, y), true)
		//} else {
		// cw.SetFgColor(cw.YELLOW)
		if i == len(line)-1 {
			char = 'X'
		}
		viewx, viewy := line[i].X-vx, line[i].Y-vy
		if areCoordsInRect(viewx, viewy, 0, 0, VIEWPORT_W, VIEWPORT_H) {
			cw.PutChar(char, viewx, viewy)
		}
		// }
	}
	if flush {
		cw.Flush_console()
	}
}

func renderCircle(fromx, fromy, radius int, flush bool, vx, vy int) {
	if radius == 0 {
		return
	} else {
		line := routines.GetCircle(fromx, fromy, radius)
		for _, point := range *line {
			x, y := point.X, point.Y
			if areCoordsInRect(x-vx, y-vy, 0, 0, VIEWPORT_W, VIEWPORT_H) {
				cw.PutChar('X', x-vx, y-vy)
			}
		}
	}
	if flush {
		cw.Flush_console()
	}
}

func getTargetingChar(x, y int) rune {
	if abs(x) > 1 {
		x /= abs(x)
	}
	if abs(y) > 1 {
		y /= abs(y)
	}
	if x == 0 {
		return '|'
	}
	if y == 0 {
		return '-'
	}
	if x*y == 1 {
		return '\\'
	}
	if x*y == -1 {
		return '/'
	}
	return '?'
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
