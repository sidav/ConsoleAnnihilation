package main

import cw "GoSdlConsole/GoSdlConsole"

func main() {
	cw.Init_console()
	defer cw.Close_console()
}
