package main

type pawnResourceInformation struct {
	metalIncome, energyIncome, metalStorage, energyStorage int
	energyReqForConditionalMetalIncome                     int
	metalSpending, energySpending                          int // both are unconditional only

}
