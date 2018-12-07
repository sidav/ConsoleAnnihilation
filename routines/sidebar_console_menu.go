package routines

import (
	cw "TCellConsoleWrapper/tcell_wrapper"
	"strings"
)

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

func ShowSidebarSingleSelectMenu(title string, titleColor, mx, my, mw int, mh int, items []string) int { // returns an index of selected element or -1 if none selected.
	drawSidebarMenuTitle(title, titleColor, mx, my, mw)
	cursorIndex := 0
	for {

		for y := 1; y < mh; y++ {
			cw.PutString(strings.Repeat(" ", mw), mx, y+my) // clear menu screen space
		}
		for i := 0; i < len(items); i++ {
			str := items[i]
			if i == cursorIndex {
				cw.SetBgColor(cw.BEIGE)
				cw.SetFgColor(cw.BLACK)
				// str = "->"+str
			} else {
				cw.SetBgColor(cw.BLACK)
				cw.SetFgColor(cw.BEIGE)
			}
			// str += strings.Repeat(" ", mw - len(str)) // fill the whole menu width
			cw.PutString(str, mx, my+i+1)
		}
		cw.SetBgColor(cw.BLACK)
		cw.Flush_console()

		key := cw.ReadKey()
		switch key {
		case "DOWN", "2":
			cursorIndex = (cursorIndex + 1) % len(items)
		case "UP", "8":
			cursorIndex -= 1
			if cursorIndex < 0 {
				cursorIndex = len(items) - 1
			}
		case "ENTER":
			return cursorIndex
		case "ESCAPE":
			return -1
		}
	}
}
