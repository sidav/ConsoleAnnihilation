package main

import "fmt"

type constructionInformation struct { // for buildings which are under construction right now
	currentConstructionAmount, maxConstructionAmount int
	costM, costE int
}

func (ci *constructionInformation) clone() *constructionInformation {
	newConstr := constructionInformation{
		costE: ci.costE,
		costM: ci.costM,
		currentConstructionAmount: ci.currentConstructionAmount,
		maxConstructionAmount: ci.maxConstructionAmount,
	}
	if ci == &newConstr {
		panic("omg not cloned")
	}
	return &newConstr
}

func (ci *constructionInformation) isCompleted() bool {
	return ci.currentConstructionAmount >= ci.maxConstructionAmount
}

func (ci *constructionInformation) getCompletionPercent() int {
	return ci.currentConstructionAmount * 100 / ci.maxConstructionAmount
}

func (ci *constructionInformation) getDescriptionString() string {
	return fmt.Sprintf("Metal: %d ENERGY: %d Base build time: %d", ci.costM, ci.costE, ci.maxConstructionAmount)
}
