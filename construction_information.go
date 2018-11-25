package main

type constructionInformation struct { // for buildings which are under construction right now
	currentConstructionAmount, maxConstructionAmount int
	costM, costE int
}

func (ci *constructionInformation) isCompleted() bool {
	return ci.currentConstructionAmount >= ci.maxConstructionAmount
}
