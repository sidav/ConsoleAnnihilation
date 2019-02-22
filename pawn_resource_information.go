package main

type pawnResourceInformation struct {
	metalIncome, energyIncome, metalStorage, energyStorage int
	energyDrain                                            int  //
	metalSpending, energySpending                          int  // both are unconditional only
	isMetalExtractor                                       bool // gives metal only when placed at metal deposit
	isGeothermalPowerplant                                 bool // gives energy only when placed at thermal vents
}

func (pre *pawnResourceInformation) resetSpendings() {
	pre.metalSpending = 0
	pre.energySpending = 0
}
