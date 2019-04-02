package main

import (
	cw "github.com/sidav/golibrl/console"
	"fmt"
)

func r_renderCursor(f *faction) {
	c := f.cursor
	cx, cy := c.getCoords()
	if !areGlobalCoordsOnScreen(cx, cy) {
		return
	}
	switch c.currentCursorMode {
	case CURSOR_SELECT:
		renderSelectCursor(f)
	case CURSOR_MULTISELECT:
		renderBandboxCursor(f)
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
	x, y := c.getOnScreenCoords()
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
	// flushView()
}

func renderBandboxCursor(f *faction) {
	cw.SetFgColor(cw.WHITE)
	fromx, fromy := f.cursor.xorig, f.cursor.yorig
	tox, toy := f.cursor.getCoords()
	if fromx > tox {
		t := fromx
		fromx = tox
		tox = t
	}
	if fromy > toy {
		t := fromy
		fromy = toy
		toy = t
	}
	for i := fromx-1; i <= tox+1; i++ {
		for j := fromy-1; j <= toy+1; j++ {
			if i == fromx-1 || i == tox+1 {
				renderCharByGlobalCoords('|', i, j)
				continue
			}
			if j == fromy-1 || j == toy+1 {
				renderCharByGlobalCoords('-', i, j)
				continue
			}
		}
	}
	flushView()
}

func renderMoveCursor(f *faction) {
	c := f.cursor
	x, y := c.getOnScreenCoords()

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
	c := f.cursor
	x, y := c.getOnScreenCoords()

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
	x, y := c.getOnScreenCoords()

	// TODO: optimize it with getPawnsInRect()
	totalMetalUnderCursor := CURRENT_MAP.getNumberOfMetalDepositsInRect(c.x-c.w/2, c.y-c.h/2, c.w, c.h)
	totalThermalUnderCursor := CURRENT_MAP.getNumberOfThermalDepositsInRect(c.x-c.w/2, c.y-c.h/2, c.w, c.h)

	if c.radius > 0 {
		cw.SetFgColor(cw.RED)
		renderCircle(c.x, c.y, c.radius, '.', false)
	}

	for i := 0; i < c.w; i++ {
		for j := 0; j < c.h; j++ {
			if (c.buildOnMetalOnly && totalMetalUnderCursor == 0) ||
				(c.buildOnThermalOnly && totalThermalUnderCursor == 0) {
				cw.SetBgColor(cw.RED)
			} else {
				if areCoordsValid(c.x+i-c.w/2, c.y+j-c.h/2) && CURRENT_MAP.tileMap[c.x+i-c.w/2][c.y+j-c.h/2].isPassable &&
					CURRENT_MAP.getPawnAtCoordinates(c.x+i-c.w/2, c.y+j-c.h/2) == nil {
					cw.SetBgColor(cw.GREEN)
				} else {
					cw.SetBgColor(cw.RED)
				}
			}
			cw.PutChar(' ', x+i-c.w/2, y+j-c.h/2)
		}
	}
	resInfoString := ""
	if totalMetalUnderCursor > 0 {
		resInfoString += fmt.Sprintf(" %dx METAL ", totalMetalUnderCursor)
	}
	if totalThermalUnderCursor > 0 {
		resInfoString += fmt.Sprintf(" %dx THERMAL ", totalThermalUnderCursor)
	}
	if len(resInfoString) > 0 {
		cw.SetBgColor(cw.DARK_GRAY)
		cw.SetFgColor(cw.WHITE)
		cw.PutString(resInfoString, x-c.w/2+c.w, y-c.h/2+c.h)
	}
	cw.SetBgColor(cw.BLACK)
}
