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

func (ai *aiData) shouldProduceEngineers() bool {
	return ai.currentEngineersCount < ai.getCurrentOrderStep().desiredEngineers
}

var ai_buildOrderNames = []string {
	"CRAZY (NO STRATEGY)",
	"2 KBOT LABS RUSH",
	"2 VEHICLE PLANTS RUSH",
	"KBOT/VEHICLE BALANCED",
	"TECH RUSH",
}

var ai_allBuildOrders = [][]*ai_buildOrderStep{
	// order 0: full random
	{
		{desiredMIncome: 5, desiredEIncome: 40, desiredEngineers: 1},
		{desiredMIncome: 10, desiredEIncome: 60, desiredEngineers: 2},
	},
	// order 1: 2 KBot Labs
	{
		{desiredMIncome: 5, desiredEIncome: 0, desiredEngineers: 1, buildCode: "armkbotlab", buildCodeAlt: "corekbotlab"},
		{desiredMIncome: 7, desiredEIncome: 0, desiredEngineers: 1, buildCode: "armkbotlab", buildCodeAlt: "corekbotlab"},
		{desiredMIncome: 10, desiredEIncome: 100, desiredEngineers: 2, buildCode: "lturret"},
		{desiredMIncome: 15, desiredEIncome: 150, desiredEngineers: 2, buildCode: ""},
	},

	// order 2: 2 vehicle factories
	{
		{desiredMIncome: 5, desiredEIncome: 0, desiredEngineers: 1, buildCode: "armvehfactory", buildCodeAlt: "corevehfactory"},
		{desiredMIncome: 6, desiredEIncome: 100, desiredEngineers: 1, buildCode: "solar"},
		{desiredMIncome: 7, desiredEIncome: 100, desiredEngineers: 1, buildCode: "armvehfactory", buildCodeAlt: "corevehfactory"},
		{desiredMIncome: 8, desiredEIncome: 150, desiredEngineers: 2, buildCode: "lturret"},
		{desiredMIncome: 12, desiredEIncome: 150, desiredEngineers: 2, buildCode: ""},
	},

	// order 3: single vehicle and single kbot
	{
		{desiredMIncome: 5, desiredEIncome: 0, desiredEngineers: 1, buildCode: "armkbotlab", buildCodeAlt: "corekbotlab"},
		{desiredMIncome: 7, desiredEIncome: 50, desiredEngineers: 1, buildCode: "armvehfactory", buildCodeAlt: "corevehfactory"},
		{desiredMIncome: 9, desiredEIncome: 100, desiredEngineers: 2, buildCode: "lturret"},
		{desiredMIncome: 12, desiredEIncome: 200, desiredEngineers: 2, buildCode: ""},
	},
	// order 4: straight to T2
	{
		{desiredMIncome: 5, desiredEIncome: 0, desiredEngineers: 2, buildCode: "armkbotlab", buildCodeAlt: "corekbotlab"},
		{desiredMIncome: 9, desiredEIncome: 100, desiredEngineers: 2, buildCode: "armt2kbotlab", buildCodeAlt: "coret2kbotlab"},
		{desiredMIncome: 12, desiredEIncome: 150, desiredEngineers: 3, buildCode: "lturret"},
		{desiredMIncome: 18, desiredEIncome: 200, desiredEngineers: 3, buildCode: ""},
	},
}
