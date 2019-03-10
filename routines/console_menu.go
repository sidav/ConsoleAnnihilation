package routines

import (
	cw "github.com/sidav/goLibRL/console"
	"strings"
)

const (
	TITLE_COLOR = cw.DARK_BLUE
	TEXT_COLOR  = cw.BEIGE
)

func drawTitle(title string) {
	cw.SetColor(cw.BLACK, TITLE_COLOR)
	consoleWidth, _ := cw.GetConsoleSize()
	for x := 0; x < consoleWidth; x++ {
		cw.PutChar(' ', x, 0)
	}
	cw.SetColor(TITLE_COLOR, cw.BLACK)
	titleXCoord := consoleWidth/2 - len(title)/2
	cw.PutString(" "+title+" ", titleXCoord, 0)
}

func DrawWrappedTextInRect(text string, x, y, w, h int) {
	currentLine := 0
	currentLineLength := 0
	for yy := 1; yy < h; yy++ {
		cw.PutString(strings.Repeat(" ", w), x, yy+y) // clear menu screen space
	}
	words := strings.Split(text, " ")
	for _, word := range words {
		if word == "\\n" {
			currentLine += 1
			currentLineLength = 0
			continue
		}
		if currentLineLength+len(word) >= w {
			currentLine += 1
			currentLineLength = 0
		}
		if currentLine == h {
			return
		}
		cw.PutString(word+" ", x+currentLineLength, y+currentLine)
		currentLineLength += len(word) + 1
	}
}

func ShowSimpleInfoWindow(title, text string, w, h, outlineColor int) {
	c_w, c_h := cw.GetConsoleSize()
	for {
		x, y := (c_w-w)/2, (c_h-h)/2
		// draw background
		cw.SetBgColor(outlineColor)
		for i := x; i < x+w; i++ {
			cw.PutChar(' ', i, y)
			cw.PutChar(' ', i, y+h)
		}
		for j := y; j <= y+h; j++ {
			cw.PutChar(' ', x, j)
			cw.PutChar(' ', x+w, j)
		}
		cw.SetBgColor(cw.BLACK)
		for i := x+1; i < x+w-1; i++ {
			for j := y+1; j < y+h-1; j++ {
				cw.PutChar(' ', i, j)
			}
		}
		cw.SetBgColor(cw.BEIGE)
		cw.SetFgColor(cw.BLACK)
		cw.PutString(title, x+(w-len(title))/2, y+1)
		cw.SetBgColor(cw.BLACK)
		cw.SetFgColor(cw.BEIGE)
		DrawWrappedTextInRect(text, x+1, y+2, w-1, h-2)
		cw.Flush_console()
		key := cw.ReadKey()
		switch key {
		case "SPACE", "ENTER", "ESCAPE":
			return
		}
	}
}

func ShowSimpleYNChoiceModalWindow(text string) bool {
	var w, h int
	var wrap bool
	c_w, c_h := cw.GetConsoleSize()
	for {
		if len(text) <= c_w/3 {
			w = len(text) + 2
			h = 5
		} else {
			wrap = true
			w = c_w/3 + 2
			h = len(text)/w + 4
		}
		x, y := (c_w-w)/2, (c_h-h)/2
		// draw background
		cw.SetBgColor(cw.DARK_GRAY)
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				cw.PutChar(' ', i, j)
			}
		}
		if wrap {
			DrawWrappedTextInRect(text, x+1, y+1, w-1, h-3)
		} else {
			cw.PutString(text, x+1, y+1)
		}
		cw.PutString("PRESS Y OR N", x+w/2-6, y+h-2)
		cw.Flush_console()
		key := cw.ReadKey()
		cw.SetBgColor(cw.BLACK)
		switch key {
		case "y", "ENTER":
			return true
		case "n", "ESCAPE":
			return false
		}
	}
}

func ShowSingleChoiceMenu(title, subheading string, lines []string) int { //returns the index of selected line or -1 if nothing was selected.
	val := lines
	cursor := 0
	for {
		cw.Clear_console()
		drawTitle(title)
		cw.SetFgColor(cw.BEIGE)
		cw.PutString(subheading, 0, 1)
		for i, _ := range val {
			if cursor == i {
				cw.SetColor(cw.BLACK, TEXT_COLOR)
			} else {
				cw.SetColor(TEXT_COLOR, cw.BLACK)
			}
			cw.PutString(" "+val[i]+" ", 1, 2+i)
			cw.SetBgColor(cw.BLACK)
		}
		cw.Flush_console()
		key := cw.ReadKey()
		switch key {
		case "2":
			cursor++
			if cursor == len(val) {
				cursor = 0
			}
		case "8":
			cursor--
			if cursor < 0 {
				cursor = len(val) - 1
			}
		case "ENTER":
			return cursor
		case "ESCAPE":
			return -1
		}
	}
}
