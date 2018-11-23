package main

func getFactionRGB(fn int) (uint8, uint8, uint8) {
	switch fn {
	case 0:
		return 0, 0, 255
	case 1:
		return 255, 0, 0
	case 2:
		return 0, 255, 0
	case 3:
		return 255, 255, 0
	}
	return 32, 32, 32
}

func getFactionColor(fn int) int {
	//BLACK        = 0
	//DARK_RED     = 1
	//DARK_GREEN   = 2
	//DARK_YELLOW  = 3
	//DARK_BLUE    = 4
	//DARK_MAGENTA = 5
	//DARK_CYAN    = 6
	//BEIGE        = 7
	//DARK_GRAY    = 8
	//RED          = 9
	//GREEN        = 10
	//YELLOW       = 11
	//BLUE         = 12
	//MAGENTA      = 13
	//CYAN         = 14
	//WHITE        = 15
	switch fn {
	case 0:
		return 14
	case 1:
		return 9
	case 2:
		return 10
	case 3:
		return 11
	}
	return 7
}

type faction struct {
	cursor                                                                      *cursor // cursor position
	economy *factionEconomy
	factionNumber                                                               int
	name                                                                        string
	playerControlled                                                            bool // used as a stub for now
}

func createFaction(name string, n int, playerControlled bool) *faction{ // temporary
	return &faction{playerControlled: playerControlled, name: name, factionNumber:n, economy: &factionEconomy{maxMetal:10, currentMetal:99999, maxEnergy:10, currentEnergy:99999}, cursor: &cursor{}}
}
