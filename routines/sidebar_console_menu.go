package routines

import (
	cw "TCellConsoleWrapper"
	"strconv"
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

func ShowSidebarSingleChoiceMenu(title string, titleColor, mx, my, mw int, mh int, items []string, descriptions []string) int { // returns an index of selected element or -1 if none selected.
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
			cw.SetBgColor(cw.BLACK)
			cw.SetFgColor(cw.BEIGE)

			if len(descriptions) > 0 {
				drawWrappedTextInRect(descriptions[cursorIndex], 0, my+mh+1, mx+mw, 5)
			}
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

func ShowSidebarPickValuesMenu(title string, titleColor, mx, my, mw int, mh int, items []string) *[]int { // returns a pointer to a slice of indices for items.
	drawSidebarMenuTitle(title, titleColor, mx, my, mw)
	cursorIndex := 0
	values := make([]int, len(items))
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
			valStr := "<" + strconv.Itoa(values[i]) + ">"
			str += strings.Repeat(" ", mw - len(str) - len(valStr)) // fill the whole menu width
			str += valStr
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
		case "LEFT", "4":
			if values[cursorIndex] > 0 {
				values[cursorIndex] -= 1
			}
		case "RIGHT", "6":
			values[cursorIndex]++
		case "ENTER":
			return &values
		case "ESCAPE":
			return nil
		}
	}
}

func ShowSidebarCreateQueueMenu(title string, titleColor, mx, my, mw int, mh int, items []string) *[]int { // returns a slice of ordered indices
	drawSidebarMenuTitle(title, titleColor, mx, my, mw)
	cursorIndex := 0
	values := make([]int, 0)
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
			valStr := ""
			for ind, val := range values {
				if val == i {
					if len(valStr) > 0 {
						valStr += ","
					}
					valStr += strconv.Itoa(ind+1)
				}
			}
			charsForFill := mw - len(str) - len(valStr)
			if charsForFill > 0 {
				str += strings.Repeat(" ", charsForFill) // fill the whole menu width
			}
			str += valStr
			cw.PutString(str, mx, my+i+1)
		}
		cw.SetBgColor(cw.BLACK)

		queue := make([]string, 0)
		multiplicator := 1
		for ind, str := range values {
			if ind > 0 && values[ind-1] == str {
				multiplicator += 1
			} else {
				if multiplicator > 1 {
					queue[len(queue)-1] += "  x"+strconv.Itoa(multiplicator)
					multiplicator = 1
				}
				queue = append(queue, items[str])
			}
		}
		if multiplicator > 1 {
			queue[len(queue)-1] += "  x"+strconv.Itoa(multiplicator)
			multiplicator = 1
		}
		DrawSidebarInfoMenu("QUEUE", cw.BLUE, mx, my+len(items), mw, queue)

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
		case "LEFT", "4":
			for i := len(values) - 1; i >= 0; i-- {
				if values[i] == cursorIndex {
					values = append(values[:i], values[i+1:]...) // removes i-th element
					break
				}
			}
		case "RIGHT", "6":
			values = append(values, cursorIndex)
		case "ENTER":
			return &values
		case "ESCAPE":
			return nil
		}
	}
}
