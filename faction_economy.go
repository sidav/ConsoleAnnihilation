package main

type factionEconomy struct {
	currentMetal, currentEnergy, metalIncome, energyIncome, metalSpending int
	energySpending, maxMetal, maxEnergy int
	nanolatheAllowed bool
}

func (fe *factionEconomy) resetFlow() {
	fe.metalIncome = 0
	fe.metalSpending = 0
	fe.energyIncome = 0
	fe.energySpending = 0
	fe.maxMetal = 0
	fe.maxEnergy = 0
}

func (eco *factionEconomy) ensureCorrectStorages() {
	// ensure that storages does not exceed max value
	if eco.currentMetal > eco.maxMetal {
		eco.currentMetal = eco.maxMetal
	}
	if eco.currentEnergy > eco.maxEnergy {
		eco.currentEnergy = eco.maxEnergy
	}
}

//func (fe *factionEconomy) getCurrentGlobalBuildingCoeff() float64 {
//
//}

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

			if u.res.isGeothermalPowerplant && u.res.energyIncome == 0 { // energy income from the geothermal needs to be recalculated.
				u.res.energyIncome = CURRENT_MAP.getNumberOfThermalDepositsUnderBuilding(u) * 40 // +40 energy per thermal vent
			}

			energyInc += u.res.energyIncome // always unconditional

			if u.res.isMetalExtractor && u.res.metalIncome == 0 { // metal income from the extractor needs to be recalculated.
				u.res.metalIncome = CURRENT_MAP.getNumberOfMetalDepositsUnderBuilding(u)
			}

			// calculate conditional metal income and matching energy spendings
			if u.res.energyDrain > 0 {
				metalConditionalInc += u.res.metalIncome
				energyConditionalDec += u.res.energyDrain
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

