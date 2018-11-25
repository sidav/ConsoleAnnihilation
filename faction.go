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
	return &faction{playerControlled: playerControlled, name: name, factionNumber:n, economy: &factionEconomy{currentMetal:99999, currentEnergy:99999}, cursor: &cursor{}}
}

// My attempt to make Total Annihilation-like economy system.  

func (f *faction) recalculateFactionEconomy(g *gameMap) { // move somewhere?
	eco := f.economy
	eco.resetFlow()
	var metalConditionalInc, metalUnconditionalInc, energyInc int
	var metalDec, energyConditionalDec, energyUnconditionalDec int

	for _, u := range g.units { // TODO: only units? FUCK!
		if u.faction == f && u.res != nil {
			eco.maxMetal += u.res.metalStorage
			eco.maxEnergy += u.res.energyStorage
			energyInc += u.res.energyIncome // always unconditional

			// calculate conditional metal income and mathing energy spendings
			if u.res.energyReqForConditionalMetalIncome > 0 {
				metalConditionalInc += u.res.metalIncome
				energyConditionalDec += u.res.energyReqForConditionalMetalIncome
			} else {
				metalUnconditionalInc += u.res.metalIncome
			}
			// Calculate unconditional spendings
			metalDec += u.res.metalSpending // always unconditional
			energyUnconditionalDec += u.res.energySpending
		}
	}
	for _, building := range g.buildings { //
		if building.faction == f && building.res != nil {
			eco.maxMetal += building.res.metalStorage
			eco.maxEnergy += building.res.energyStorage
			energyInc += building.res.energyIncome // always unconditional

			// calculate conditional metal income and mathing energy spendings
			if building.res.energyReqForConditionalMetalIncome > 0 {
				metalConditionalInc += building.res.metalIncome
				energyConditionalDec += building.res.energyReqForConditionalMetalIncome
			} else {
				metalUnconditionalInc += building.res.metalIncome
			}
			// Calculate unconditional spendings
			metalDec += building.res.metalSpending // always unconditional
			energyUnconditionalDec += building.res.energySpending
		}
	}
	// If spending is allowed with conditional, then spend/gain everything
	if f.isSpendingAllowedWithBalance(metalConditionalInc+metalUnconditionalInc, metalDec, energyInc, energyUnconditionalDec+energyConditionalDec) {
		eco.nanolatheAllowed = true

		eco.metalIncome = metalConditionalInc + metalUnconditionalInc
		eco.metalSpending = metalDec
		eco.currentMetal += eco.metalIncome - eco.metalSpending

		eco.energyIncome = energyInc
		eco.energySpending = energyConditionalDec + energyUnconditionalDec
		eco.currentEnergy += eco.energyIncome - eco.energySpending

	} else { // spending is disallowed, so we just gain resources and got no spendings
		eco.nanolatheAllowed = false

		eco.metalIncome = metalUnconditionalInc
		eco.metalSpending = metalDec
		eco.currentMetal += metalUnconditionalInc

		eco.energyIncome = energyInc
		eco.energySpending = energyConditionalDec + energyUnconditionalDec
		eco.currentEnergy += eco.energyIncome
	}

	eco.ensureCorrectStorages()
}

func (f *faction) isSpendingAllowedWithBalance(metalInc, metalDec, energyInc, energyDec int) bool {
	eco := f.economy
	ms := eco.currentMetal
	es := eco.currentEnergy
	return (ms + metalInc >= metalDec) && (es+energyInc >= energyDec)
}
