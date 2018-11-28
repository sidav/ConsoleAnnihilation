package main

type constructionInformation struct { // for buildings which are under construction right now
	currentConstructionAmount, maxConstructionAmount int
	costM, costE int
}

func (ci *constructionInformation) isCompleted() bool {
	return ci.currentConstructionAmount >= ci.maxConstructionAmount
}

func (ci *constructionInformation) getCompletionPercent() int {
	return ci.currentConstructionAmount * 100 / ci.maxConstructionAmount
}
