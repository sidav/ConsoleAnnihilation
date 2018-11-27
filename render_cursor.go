package main

import cw "TCellConsoleWrapper/tcell_wrapper"

func r_renderCursor(f *faction) {
	c := f.cursor
	switch c.currentCursorMode {
	case CURSOR_SELECT:
		renderSelectCursor(c)
	case CURSOR_MOVE:
		renderMoveCursor(c)
	}
}

func renderSelectCursor(c *cursor) {
	x := VIEWPORT_W / 2
	y := VIEWPORT_H / 2
	snap := c.snappedBuilding
	// cw.SetFgColorRGB(128, 128, 128)
	cw.SetFgColor(cw.WHITE)

	if snap == nil {
		cw.PutChar('[', x-1, y)
		cw.PutChar(']', x+1, y)
	} else {
		w, h := snap.w, snap.h
		offset := w % 2
		for cy := 0; cy < h; cy++ {
			cw.PutChar('[', x-w/2-1, cy-h/2+y)
			cw.PutChar(']', x+w/2+offset, cy-h/2+y)
		}
	}

	// outcommented for non-SDL console
	//cw.PutChar(16*13+10, x-1, y-1)
	//cw.PutChar(16*11+15, x+1, y-1)
	//cw.PutChar(16*12, x-1, y+1)
	//cw.PutChar(16*13+9, x+1, y+1)
	flushView()
}

func renderMoveCursor(c *cursor) {
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
