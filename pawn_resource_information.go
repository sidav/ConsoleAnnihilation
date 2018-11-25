package main

type pawnResourceInformation struct {
	metalIncome, energyIncome, metalStorage, energyStorage int
	energyReqForConditionalMetalIncome                     int
	metalSpending, energySpending                          int // both are unconditional only

}

func (pre *pawnResourceInformation) resetSpendings() {
	pre.metalSpending = 0
	pre.energySpending = 0
}
