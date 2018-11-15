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
	switch fn {
	case 0:
		return 12
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
	cx, cy int // cursor position
	currentMetal, currentEnergy, metalIncome, energyIncome, maxMetal, maxEnergy int
	factionNumber               int
	name string
}

func createFaction(name string, n int) *faction{ // temporary
	return &faction{name: name, factionNumber:n, maxMetal:10, currentMetal:10, maxEnergy:10, currentEnergy:10}
}
