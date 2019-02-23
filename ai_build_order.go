package main

type ai_buildOrderStep struct {
	desiredMIncome, desiredEIncome int
	desiredEngineers               int

	buildCode, buildCodeAlt string // alt is for when a faction has no buildCode
}

func (ai *aiData) getCurrentOrderStep() *ai_buildOrderStep {
	return (*ai.buildOrder)[ai.currentStepNumber]
}

func (ai *aiData) orderStepSatisfied() {
	if ai.currentStepNumber + 1 < len(*ai.buildOrder) {
		ai.currentStepNumber++
	}
}

var ai_allBuildOrders = [][]*ai_buildOrderStep{
	// order 1: 2 KBot Labs
	{
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 1, buildCode: "armkbotlab", buildCodeAlt: "corekbotlab"},
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 1, buildCode: "solar"},
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 1, buildCode: "armkbotlab", buildCodeAlt: "corekbotlab"},
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 2, buildCode: "lturret"},
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 2, buildCode: ""},
	},

	// order 2: 2 vehicle factories
	{
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 1, buildCode: "armvehfactory", buildCodeAlt: "corevehfactory"},
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 1, buildCode: "solar"},
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 1, buildCode: "armvehfactory", buildCodeAlt: "corevehfactory"},
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 2, buildCode: "lturret"},
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 2, buildCode: ""},
	},

	// order 3: single vehicle and single kbot
	{
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 1, buildCode: "armkbotlab", buildCodeAlt: "corekbotlab"},
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 1, buildCode: "solar"},
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 1, buildCode: "armvehfactory", buildCodeAlt: "corevehfactory"},
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 1, buildCode: "solar"},
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 1, buildCode: "solar"},
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 2, buildCode: "lturret"},
		{desiredMIncome: 0, desiredEIncome: 0, desiredEngineers: 2, buildCode: ""},
	},
}