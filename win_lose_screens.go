package main

import cw "TCellConsoleWrapper"

var (
	coreLogo = []string{
		"   ** **       ",
		"  ** # **      ",
		" ** ### **     ",
		"** ##### **    ",
		"* ####### *    ",
		" #########     ",
		"##### #####    ",
		"#### * ####    ",
		"### *** ###    ",
		"## ***** ##    ",
		" #  ***  #     ",
		" #       #     ",
		" ##     ##     ",
		"  #     #      ",
	}
	coreWon = []string{
		"          *          ",
		"       ##***##       ",
		"     ## ***** ##     ",
		"    #  *******  #    ",
		"   #  *********  #   ",
		"  #  ***********  #  ",
		"  # ************* #  ",
		" # *************** # ",
		" #*****************# ",
		" ******************* ",
		"********** **********",
		"*********   *********",
		" *******     ******* ",
		" ******       ****** ",
		"  ****         ****  ",
		"  ****         ****  ",
		"  ****         ****  ",
		"   ***         ***   ",
		"   ***         ***   ",
		"    **##     ##**    ",
		"    **  #####  **    ",
		"     *         *     ",
		"     *         *     ",
	}
	coreLost = []string{
		"          *          ",
		"       ##***##       ",
		"     ## ***** #      ",
		"    #  ******   #    ",
		"   #  ****** **  #   ",
		"     *****  ****  #  ",
		"  #  **    ****** #  ",
		" # *    * *  ***** # ",
		" #*** ***  ** *****# ",
		" **** **** *** ** ** ",
		"****** *** ***   ****",
		"   *** **   ** ** ***",
		" *  **         ***** ",
		" *** **       ****** ",
		"  ** *         **    ",
		"  ** *         *  *  ",
		"  **             **  ",
		"   ***         ***   ",
		"   ***         ***   ",
		"    **##     ##**    ",
		"    **   ## #  **    ",
		"     *         *     ",
		"     *         *     ",
	}
)

func r_drawWinLogo() {
	logo := &coreWon
	cw.Clear_console()
	cy := (cw.CONSOLE_HEIGHT - len(*logo)) / 2
	cx := (cw.CONSOLE_WIDTH - len((*logo)[0])) / 2
	for y := 0; y < len(*logo); y++ {
		for x := 0; x < len((*logo)[y]); x++ {
			chr := (*logo)[y][x]
			switch chr {
			case ' ':
				continue
			case '#':
				cw.SetBgColor(cw.DARK_GREEN)
				cw.PutChar(' ', x+cx, y+cy)
				cw.SetBgColor(cw.BLACK)
			case '*':
				cw.SetFgColor(cw.GREEN)
				cw.PutChar('*', x+cx, y+cy)
			default:
				cw.PutChar('?', x+cx, y+cy)
			}
		}
	}
	cw.Flush_console()
}

func r_drawLoseLogo() {
	logo := &coreLost
	cw.Clear_console()
	cy := (cw.CONSOLE_HEIGHT - len(*logo)) / 2
	cx := (cw.CONSOLE_WIDTH - len((*logo)[0])) / 2
	for y := 0; y < len(*logo); y++ {
		for x := 0; x < len((*logo)[y]); x++ {
			chr := (*logo)[y][x]
			switch chr {
			case ' ':
				continue
			case '#':
				cw.SetBgColor(cw.DARK_GRAY)
				cw.PutChar(' ', x+cx, y+cy)
				cw.SetBgColor(cw.BLACK)
			case '*':
				cw.SetFgColor(cw.DARK_GREEN)
				cw.PutChar('*', x+cx, y+cy)
			default:
				cw.PutChar('?', x+cx, y+cy)
			}
		}
	}
	cw.Flush_console()
}

func r_showTitleScreen() {
	r_drawWinLogo()
	str := "TOTAL ANNIHILATION: THE PREQUEL"
	cw.PutString(str, (cw.CONSOLE_WIDTH-len(str))/2, 0)
	str = "Press enter to continue"
	cw.PutString(str, (cw.CONSOLE_WIDTH-len(str))/2, cw.CONSOLE_HEIGHT-1)
	cw.Flush_console()
	key := ""
	for key != "ESCAPE" && key != "ENTER" {
		key = cw.ReadKey()
	}
}

func r_gamelostScreen() {
	r_drawLoseLogo()
	cw.SetFgColor(cw.GREEN)
	str := "YOU HAVE LOST"
	cw.PutString(str, (cw.CONSOLE_WIDTH-len(str))/2, 0)
	str = "THE ARM VERMIN WILL SOON OVERRUN THALASSEAN"
	cw.PutString(str, (cw.CONSOLE_WIDTH-len(str))/2, cw.CONSOLE_HEIGHT-1)
	cw.Flush_console()
	key := ""
	for key != "ESCAPE" && key != "ENTER" {
		key = cw.ReadKey()
	}
}

func r_gameWonScreen() {
	r_drawWinLogo()
	cw.SetFgColor(cw.GREEN)
	str := "MISSION ACCOMPLISHED"
	cw.PutString(str, (cw.CONSOLE_WIDTH-len(str))/2, 0)
	str = "THANKS FOR PLAYING!"
	cw.PutString(str, (cw.CONSOLE_WIDTH-len(str))/2, cw.CONSOLE_HEIGHT-1)
	cw.Flush_console()
	key := ""
	for key != "ESCAPE" && key != "ENTER" {
		key = cw.ReadKey()
	}
}
