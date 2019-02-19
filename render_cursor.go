package main

import cw "TCellConsoleWrapper"

func r_renderCursor(f *faction) {
	c := f.cursor
	switch c.currentCursorMode {
	case CURSOR_SELECT:
		renderSelectCursor(f)
	case CURSOR_MOVE:
		renderMoveCursor(f)
	case CURSOR_AMOVE:
		renderAttackMoveCursor(f)
	case CURSOR_BUILD:
		renderBuildCursor(c)
	}
}

func renderSelectCursor(f *faction) {
	c := f.cursor
	x := VIEWPORT_W / 2
	y := VIEWPORT_H / 2
	snap := c.snappedPawn
	// cw.SetFgColorRGB(128, 128, 128)
	if snap == nil {
		cw.SetFgColor(cw.WHITE)
	} else if snap.faction == f {
		cw.SetFgColor(cw.GREEN)
	} else {
		cw.SetFgColor(cw.RED)
	}

	if snap == nil || snap.isUnit() {
		cw.PutChar('[', x-1, y)
		cw.PutChar(']', x+1, y)
	} else {
		w, h := snap.buildingInfo.w, snap.buildingInfo.h
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

func renderMoveCursor(f *faction) {
	c := f.cursor
	x := VIEWPORT_W / 2
	y := VIEWPORT_H / 2
	// cw.SetFgColorRGB(128, 255, 128)
	cw.SetFgColor(cw.GREEN)
	if c.snappedPawn != nil && c.snappedPawn.faction != f {
		cw.SetFgColor(cw.DARK_RED)
		cw.PutChar('}', x-1, y)
		cw.PutChar('{', x+1, y)
		cw.PutChar('-', x-2, y)
		cw.PutChar('-', x+2, y)
	} else {
		cw.PutChar('>', x-1, y)
		cw.PutChar('<', x+1, y)
	}

	//cw.PutChar('\\', x-1, y-1)
	//cw.PutChar('/', x+1, y-1)
	//cw.PutChar('/', x-1, y+1)
	//cw.PutChar('\\', x+1, y+1)

	flushView()
}

func renderAttackMoveCursor(f *faction) {
	x := VIEWPORT_W / 2
	y := VIEWPORT_H / 2
	// cw.SetFgColorRGB(128, 255, 128)
	cw.SetFgColor(cw.DARK_RED)
	cw.PutChar('}', x-1, y)
	cw.PutChar('{', x+1, y)
	cw.SetFgColor(cw.GREEN)
	cw.PutChar('v', x, y-1)
	cw.PutChar('^', x, y+1)

	//cw.PutChar('\\', x-1, y-1)
	//cw.PutChar('/', x+1, y-1)
	//cw.PutChar('/', x-1, y+1)
	//cw.PutChar('\\', x+1, y+1)

	flushView()
}

func renderBuildCursor(c *cursor) {
	x := VIEWPORT_W / 2
	y := VIEWPORT_H / 2

	// TODO: optimize it with getPawnsInRect()
	for i := 0; i < c.w; i++ {
		for j := 0; j < c.h; j++ {
			if (c.buildOnMetalOnly && CURRENT_MAP.getNumberOfMetalDepositsInRect(c.x-c.w/2, c.y-c.h/2, c.w, c.h) == 0) ||
				(c.buildOnThermalOnly && CURRENT_MAP.getNumberOfThermalDepositsInRect(c.x-c.w/2, c.y-c.h/2, c.w, c.h) == 0) {
				cw.SetBgColor(cw.RED)
			} else {
				if CURRENT_MAP.getPawnAtCoordinates(c.x+i-c.w/2, c.y+j-c.h/2) == nil {
					cw.SetBgColor(cw.GREEN)
				} else {
					cw.SetBgColor(cw.RED)
				}
			}
			cw.PutChar(' ', x+i-c.w/2, y+j-c.h/2)
		}
	}
	cw.SetBgColor(cw.BLACK)
}
