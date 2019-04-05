package main

import (
	geometry "github.com/sidav/golibrl/geometry"
	cw "github.com/sidav/golibrl/console"
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

func r_updateBoundsIfNeccessary(force bool) {
	if cw.WasResized() || force {
		CONSOLE_W, CONSOLE_H = cw.GetConsoleSize()
		VIEWPORT_W           = CONSOLE_W / 2
		VIEWPORT_H           = CONSOLE_H - LOG_HEIGHT - 1
		SIDEBAR_X            = VIEWPORT_W + 1
		SIDEBAR_W            = CONSOLE_W - VIEWPORT_W - 1
		SIDEBAR_H            = CONSOLE_H - LOG_HEIGHT
		SIDEBAR_FLOOR_2      = 7  // y-coord right below resources info
		SIDEBAR_FLOOR_3      = 11 // y-coord right below "floor 2"
	}
}

func r_renderScreenForFaction(f *faction, g *gameMap, selection *[]*pawn, flush bool) {
	r_updateBoundsIfNeccessary(false)
	cw.Clear_console() // TODO: replace with ClearViewportOnly (and create it of course). Prevent overrendering of whole screen.
	renderMapInViewport(f, g)
	renderFactionStats(f)
	renderInfoOnCursor(f, g)
	r_renderUIOutline(f)
	renderPawnsInViewport(f, g)
	if selection != nil && len(*selection) != 0 {
		r_renderSelectedPawns(f, selection)
		if len(*selection) == 1 {
			r_renderPossibleOrdersForPawn((*selection)[0])
		} else {
			r_renderPossibleOrdersForMultiselection(f, selection)
		}
	}
	r_renderCursor(f)
	renderLog(false)
	if flush {
		flushView()
	}
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

func renderMapInViewport(f *faction, g *gameMap) {
	vx, vy := f.cursor.getCameraCoords()
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

func renderPawnsInViewport(f *faction, g *gameMap) {
	vx, vy := f.cursor.getCameraCoords()
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
	vx, vy := f.cursor.getCameraCoords()
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
	if areGlobalCoordsOnScreen(p.x, p.y) && f.areCoordsInSight(p.x, p.y){
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
				if getCurrentTurn() % 2 == 0 {
					colorToRender = cw.GREEN
				}
			}
			if areGlobalCoordsOnScreen(bx+x, by+y) && f.wereCoordsSeen(bx+x, by+y) {
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

func flushView() {
	cw.Flush_console()
}

func renderCharByGlobalCoords(c rune, x, y int) { // TODO: use it everywhere
	vx, vy := CURRENT_FACTION_SEEING_THE_SCREEN.cursor.getCameraCoords()
	if areGlobalCoordsOnScreenForFaction(x, y, CURRENT_FACTION_SEEING_THE_SCREEN) {
		cw.PutChar(c, x-vx, y-vy)
	}
}

func areGlobalCoordsOnScreen(gx, gy int) bool {
	vx, vy := CURRENT_FACTION_SEEING_THE_SCREEN.cursor.getCameraCoords()
	return geometry.AreCoordsInRect(gx, gy, vx, vy, VIEWPORT_W, VIEWPORT_H)
}

func areGlobalCoordsOnScreenForFaction(gx, gy int, f *faction) bool {
	vx, vy := f.cursor.getCameraCoords()
	return geometry.AreCoordsInRect(gx, gy, vx, vy, VIEWPORT_W, VIEWPORT_H)
}
