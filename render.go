package main

import cw "GoSdlConsole/GoSdlConsole"

func r_setFgColorByCcell(c *ccell){
	cw.SetFgColorRGB(c.r, c.g, c.b)
}

func r_renderMap(g *gameMap) {
	cw.Clear_console()
	for x:=0; x < mapW; x++ {
		for y :=0; y < mapH; y++ {
			tileApp := g.tileMap[x][y].appearance
			r_setFgColorByCcell(tileApp)
			cw.PutChar(tileApp.char, x, y)
		}
	}

	for _, u := range g.units {
		tileApp := u.appearance
		r, g, b := getFactionRGB(u.faction)
		cw.SetFgColorRGB(r, g, b)
		cw.PutChar(tileApp.char, u.x, u.y)
	}

	cw.Flush_console()
}
