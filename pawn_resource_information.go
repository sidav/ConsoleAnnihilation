package main

type pawnIncomeInformation struct {
	metalIncome, energyIncome, metalStorage, energyStorage int
	energyDrain                                            int 
	isMetalExtractor                                       bool // gives metal only when placed at metal deposit
	isGeothermalPowerplant                                 bool // gives energy only when placed at thermal vents
}
