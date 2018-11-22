package routines

import cw "TCellConsoleWrapper/tcell_wrapper"

func drawSidebarMenuTitle(title string, titleColor, mx, my, mw int) {
	cw.SetColor(cw.BLACK, titleColor)
	for x:=mx; x<mx+mw; x++{
		cw.PutChar(' ', x, my)
	}
	cw.SetColor(titleColor, cw.BLACK)
	titleXCoord := mx + (mw / 2 - (len(title)+2) / 2)
	cw.PutString(" " +title+" ", titleXCoord, my)
}

func DrawSidebarInfoMenu(title string, titleColor, mx, my, mw int, items []string) { // no cursor, no selection, etc., just plain menu
	drawSidebarMenuTitle(title, titleColor, mx, my, mw)
	cw.SetFgColor(cw.BEIGE)
	for i:=0; i<len(items); i++ {
		cw.PutString(items[i], mx, my+i+1)
	}
	// no flush?
}

func DrawSidebarLineMenu(title string, mx, my, mw int, items []string) {
	// drawSidebarMenuTitle(title, mx, my, mw)
}
