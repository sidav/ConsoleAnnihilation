package main

type pawnResourceInformation struct {
	metalIncome, energyIncome, metalStorage, energyStorage int
	energyReqForConditionalMetalIncome                     int
	metalSpending, energySpending                          int // both are unconditional only
	isMetalExtractor                                       bool // gives metal only when placed at metal deposit
}

func (pre *pawnResourceInformation) resetSpendings() {
	pre.metalSpending = 0
	pre.energySpending = 0
}
