package main

import cw "TCellConsoleWrapper"

var (
	coreWon = []string {
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
	coreLost = []string { 
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

func r_showTitleScreen() {
	r_drawWinLogo()

}

func r_drawWinLogo() {
	logo := &coreWon
	cw.Clear_console()
	cy := (cw.CONSOLE_HEIGHT - len(*logo))/2
	cx := (cw.CONSOLE_WIDTH - len((*logo)[0]))/2
	for y :=0; y <len(*logo); y++{
		for x :=0; x <len((*logo)[y]); x++ {
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
	cy := (cw.CONSOLE_HEIGHT - len(*logo))/2
	cx := (cw.CONSOLE_WIDTH - len((*logo)[0]))/2
	for y :=0; y <len(*logo); y++{
		for x :=0; x <len((*logo)[y]); x++ {
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
