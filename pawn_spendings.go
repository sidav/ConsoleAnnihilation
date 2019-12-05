package main 

type pawnSpendings struct {
	metalSpending, energySpending int // both are unconditional only
}

func (p *pawnSpendings) resetSpendings() {
	p.metalSpending = 0
	p.energySpending = 0
}
