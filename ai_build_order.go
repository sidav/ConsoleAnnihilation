package main

type ai_buildOrderStep struct {
	desiredMIncome, desiredEIncome int
	desiredEngineers               int

	buildCode, buildCodeAlt string // alt is for when a faction has no buildCode
}

var ai_allBuildOrders = [][]*ai_buildOrderStep{
	// order 1: KBot Lab
	{
		{0, 0, 1, "armkbotlab", "corekbotlab"},
		{0, 0, 2, "armkbotlab", "corekbotlab"},
	},

	// order 2: vehicle factory
	{
		{0, 0, 1, "armvehfactory", "corevahfactory"},
	},
}


