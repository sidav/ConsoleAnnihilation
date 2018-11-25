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
