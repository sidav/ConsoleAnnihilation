package main

type factionEconomy struct {
	currentMetal, currentEnergy, metalIncome, energyIncome, metalSpending, energySpending, maxMetal, maxEnergy int
}

func (fe *factionEconomy) resetFlow() {
	fe.metalIncome = 0
	fe.metalSpending = 0
	fe.energyIncome = 0
	fe.energySpending = 0
	fe.maxMetal = 0
	fe.maxEnergy = 0
}

//func (fe *factionEconomy) getCurrentGlobalBuildingCoeff() float64 {
//
//}
