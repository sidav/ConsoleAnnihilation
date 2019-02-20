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

type faction struct {
	cursor                  *cursor // cursor position
	economy                 *factionEconomy
	factionNumber           int
	name                    string
	playerControlled        bool // used as a stub for now
	seenTiles, tilesInSight [mapW][mapH] bool
}

func createFaction(name string, n int, playerControlled bool) *faction { // temporary
	return &faction{playerControlled: playerControlled, name: name, factionNumber: n, economy: &factionEconomy{currentMetal: 99999, currentEnergy: 99999}, cursor: &cursor{}}
}

func (f *faction) getFactionColor() int {
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
	switch f.factionNumber {
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

/////////////////////////////////////////////////////////////////
// My attempt to make Total Annihilation-like economy system.

func (f *faction) recalculateFactionEconomy(g *gameMap) { // move somewhere?
	eco := f.economy
	eco.resetFlow()
	var metalConditionalInc, metalUnconditionalInc, energyInc int
	var metalDec, energyConditionalDec, energyUnconditionalDec int

	for _, u := range g.pawns {
		if u.faction == f && u.res != nil && u.currentConstructionStatus == nil {
			eco.maxMetal += u.res.metalStorage
			eco.maxEnergy += u.res.energyStorage
			energyInc += u.res.energyIncome // always unconditional

			if u.res.isMetalExtractor && u.res.metalIncome == 0 { // metal income from the extractor needs to be recalculated.
				u.res.metalIncome = CURRENT_MAP.getNumberOfMetalDepositsUnderBuilding(u)
			}

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
	// If energy spending is allowed, then spend/gain ONLY conditional
	if f.isEnergySpendingAllowed(energyInc, energyUnconditionalDec+energyConditionalDec) {
		eco.metalIncome = metalConditionalInc + metalUnconditionalInc
		eco.metalSpending = metalDec
		eco.currentMetal += eco.metalIncome

		eco.energyIncome = energyInc
		eco.energySpending = energyConditionalDec + energyUnconditionalDec
		eco.currentEnergy += eco.energyIncome - energyConditionalDec

	} else { // energy spending is disallowed, so we just gain ONLY unconditional resources
		eco.metalIncome = metalUnconditionalInc
		eco.metalSpending = metalDec
		eco.currentMetal += metalUnconditionalInc

		eco.energyIncome = energyInc
		eco.energySpending = energyConditionalDec + energyUnconditionalDec
		eco.currentEnergy += eco.energyIncome
	}

	eco.nanolatheAllowed = f.isSpendingAllowedWithBalance(0, metalDec, 0, energyUnconditionalDec) // metalInc, energyInc and energyConditionalDec are already taken into account, so 0
	if eco.nanolatheAllowed { // if nanolathe allowed, spend UNconditional energy and metal spendings
		eco.currentMetal -= eco.metalSpending
		eco.currentEnergy -= energyUnconditionalDec
	}

	eco.ensureCorrectStorages()
}

func (f *faction) isEnergySpendingAllowed(energyInc, energyDec int) bool {
	eco := f.economy
	es := eco.currentEnergy
	return es+energyInc >= energyDec
}

func (f *faction) isSpendingAllowedWithBalance(metalInc, metalDec, energyInc, energyDec int) bool {
	eco := f.economy
	ms := eco.currentMetal
	es := eco.currentEnergy
	return (ms+metalInc >= metalDec) && (es+energyInc >= energyDec)
}
