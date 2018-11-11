package main

import cw "GoSdlConsole/GoSdlConsole"

const (
	VIEWPORT_W = 25
	VIEWPORT_H = 20
)

func r_setFgColorByCcell(c *ccell) {
	cw.SetFgColorRGB(c.r, c.g, c.b)
}

func r_renderMapAroundCursor(g *gameMap, cx, cy int) {
	cw.Clear_console()
	vx := cx - VIEWPORT_W / 2
	vy := cy - VIEWPORT_H / 2
	renderMapInViewport(g, vx, vy)
	renderUnitsInViewport(g, vx, vy)
	renderLog(false)
	cw.Flush_console()
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
		tileApp := u.appearance
		r, g, b := getFactionRGB(u.faction)
		cw.SetFgColorRGB(r, g, b)
		cw.PutChar(tileApp.char, u.x-vx, u.y-vy)
	}

}

func renderSelectCursor() {
	x := VIEWPORT_W / 2
	y := VIEWPORT_H / 2
	cw.SetFgColorRGB(128, 128, 128)
	cw.PutChar(16*13+10, x-1, y-1)
	cw.PutChar(16*11+15, x+1, y-1)
	cw.PutChar(16*12, x-1, y+1)
	cw.PutChar(16*13+9, x+1, y+1)
	cw.Flush_console()
}

func renderMoveCursor() {
	x := VIEWPORT_W / 2
	y := VIEWPORT_H / 2
	cw.SetFgColorRGB(128, 255, 128)
	cw.PutChar('\\', x-1, y-1)
	cw.PutChar('/', x+1, y-1)
	cw.PutChar('/', x-1, y+1)
	cw.PutChar('\\', x+1, y+1)
	cw.Flush_console()
}

func renderLog(flush bool) {
	cw.SetFgColor(cw.WHITE)
	for i := 0; i < LOG_HEIGHT; i++ {
		cw.PutString(log.last_msgs[i].getText(), 0, VIEWPORT_H+i)
	}
	if flush {
		cw.Flush_console()
	}
}
