package main

import cw "TCellConsoleWrapper"

var (
	coreLogoArt = []string {
"            *                ", 
"         ##***##             ",
"       ## ***** ##           ",
"      #  *******  #          ",
"     #  *********  #         ",
"    #  ***********  #        ",
"   #  *************  #       ",
"  #  ***************  #      ",
"  # ******** ******** #      ",
"   ********   ********       ",
"  ********     ********      ",
"   ******       ******       ",
"   ******       ******       ",
"    *****       *****        ",
"    *****       *****        ",
"     ****#     #****         ",
"     **** ##### ****         ",
"      ***       ***          ",
"      ***       ***          ",
"       **       **           ",
"       **       **           ",
"        *       *            ",
	}
)

func r_drawLogo(logo *[]string) {
	for x:=0;x<len(*logo);x++{
		for y:=0;y<len((*logo)[x]);y++ {
			chr := (*logo)[x][y]
			switch chr {
			case ' ':
				continue
			case '#':
				cw.SetBgColor(cw.DARK_GREEN)
				cw.PutChar(' ', y, x)
				cw.SetBgColor(cw.BLACK)
			case '*':
				cw.SetFgColor(cw.DARK_GREEN)
				cw.PutChar('*', y, x)
			default:
				cw.PutChar('?', y, x)
			}
		}
	}
	cw.Flush_console()
}
