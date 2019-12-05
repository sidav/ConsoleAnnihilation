package main

var buildingsStaticDataMap map[string]*buildingStaticData

func initBuildingsStaticDataMap() {
	buildingsStaticDataMap = make(map[string]*buildingStaticData)

	colors := []int{
		7, 7,
		7, -1}
	app := &buildingAppearance{chars: "" +
		"xx" +
		"xx", colors: colors}
	buildingsStaticDataMap["metalmaker"] = &buildingStaticData{name: "Metal Synthesizer",
		w: 2, h: 2, appearance: app, allowsTightPlacement: true,
		defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 35, costM: 10, costE: 500},
		defaultIncomeData:       &pawnIncomeInformation{metalIncome: 1, energyDrain: 60},
	}

	colors = []int{
		-1, 7,
		7, -1}
	app = &buildingAppearance{chars: "" +
		"==" +
		"==", colors: colors}
	buildingsStaticDataMap["solar"] = &buildingStaticData{name: "Solar Collector",
		w: 2, h: 2, appearance: app, allowsTightPlacement: true,
		defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 60, costM: 100, costE: 500},
		defaultIncomeData:       &pawnIncomeInformation{energyIncome: 20},
	}

		colors = []int{
			-1, 7, -1,
			7, -1, 7,
			-1, 7, -1,
		}
		app = &buildingAppearance{chars:
			"#|#" +
			"-%-" +
			"#|#", colors: colors}
			buildingsStaticDataMap["mextractor"] = &buildingStaticData{name: "Metal Extractor",
			w: 3, h: 3, appearance: app, canBeBuiltOnMetalOnly: true, allowsTightPlacement: true,
			defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 60, costM: 100, costE: 500},
			defaultIncomeData:       &pawnIncomeInformation{energyDrain: 9, isMetalExtractor: true},
		}
 
		colors = []int{
			7, 13, 13, 7,
			-1, 8, 8, -1,
			-1, 8, 8, -1,
			7, 13, 13, 7}
		app = &buildingAppearance{chars: "" +
			"=^^=" +
			"{00}" +
			"{00}" +
			"=VV=", colors: colors}
			buildingsStaticDataMap["geo"] = &buildingStaticData{name: "Geothermal Powerplant",
			w: 4, h: 4, appearance: app, canBeBuiltOnThermalOnly: true, allowsTightPlacement: true,
			defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 60, costM: 350, costE: 1500},
			defaultIncomeData:       &pawnIncomeInformation{isGeothermalPowerplant: true},
		}

	
		colors = []int{
			-1, 7, -1,
			7, -1, 7,
			-1, 7, -1,
		}
		app = &buildingAppearance{chars: "" +
			"=0=" +
			"0$0" +
			"=0=", colors: colors}
			buildingsStaticDataMap["mstorage"] =  &buildingStaticData{name: "Metal Storage", maxHitpoints: 100, isHeavy: true, regenPeriod: 10,
			w: 3, h: 3, appearance: app, allowsTightPlacement: true,
			defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 60, costM: 100, costE: 800},
			defaultIncomeData:       &pawnIncomeInformation{metalStorage: 250},
		}

	
		colors = []int{
			-1, 7, -1,
			7, -1, 7,
			-1, 7, -1,
		}
		app = &buildingAppearance{chars: "" +
			"=|=" +
			"-0-" +
			"=|=", colors: colors}
			buildingsStaticDataMap["estorage"] = &buildingStaticData{name: "Energy Storage",
			w: 3, h: 3, appearance: app, allowsTightPlacement: true,
			defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 60, costM: 200, costE: 400},
			defaultIncomeData:       &pawnIncomeInformation{energyStorage: 250},
		}

	 // cheating building, useful for debugging
		colors = []int{
			7, -1, 7,
			-1, 13, -1,
			7, -1, 7}
		app = &buildingAppearance{chars: "" +
			"=^=" +
			"<O>" +
			"=V=", colors: colors}
			buildingsStaticDataMap["quark"] = &buildingStaticData{name: "Quark Antimatter Generator",
			w: 3, h: 3, appearance: app,
			defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 40, costM: 100, costE: 1200},
			defaultIncomeData:       &pawnIncomeInformation{metalIncome: 5, energyIncome: 100},
		}

	// buildingsStaticDataMap["armhq"] =  // cheating building for stub AI
	// 	colors = []int{
	// 		7, 13, 13, 7,
	// 		13, -1, -1, 13,
	// 		7, 13, 13, 7}
	// 	app = &buildingAppearance{chars: "" +
	// 		"=^^=" +
	// 		"<HQ>" +
	// 		"=VV=", colors: colors}
	// 	b = &buildingStaticData{name: "Arm proxy HQ", maxHitpoints: 300, isHeavy: true, isCommander: true, regenPeriod: 10, radarRadius: 40,
	// 		w: 4, h: 3, appearance: app},
	// 		defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 40, costM: 100, costE: 1200},
	// 		defaultIncomeData:       &pawnIncomeInformation{metalIncome: 100, metalStorage: 1000, energyIncome: 1000, energyStorage: 1000},
	// 		nanolatherInfo:            &nanolatherInformation{builderCoeff: 1, allowedUnits: []string{"armpeewee", "armhammer"}},
	// 		weapons: []*pawnWeaponInformation{
	// 			{attackDelay: 12, attackEnergyCost: 15, attackRadius: 6, attacksLand: true,
	// 				hitscan: &WeaponHitscan{baseDamage: 6},
	// 			},
	// 		},
	// 	}

	
		colors = []int{
			7, 7, 7,
			7, -1, 7,
			7, 7, 7}
		app = &buildingAppearance{chars: "" +
			"/=\\" +
			"=x=" +
			"\\=/", colors: colors}
			buildingsStaticDataMap["armkbotlab"] =  &buildingStaticData{name: "Tech 1 KBot Lab",
			w: 3, h: 3, appearance: app,
			nanolatherInfo:            &nanolatherInformation{builderCoeff: 1, allowedUnits: []string{"armck", "armpeewee", "armhammer"}},
			defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500},
		}

	
		colors = []int{
			7, 7, 7,
			7, -1, 7,
			7, 7, 7}
		app = &buildingAppearance{chars: "" +
			"\\=/" +
			"=0=" +
			"/=\\", colors: colors}
			buildingsStaticDataMap["corekbotlab"] =  &buildingStaticData{name: "Tech 1 Core KBot Lab",
			w: 3, h: 3, appearance: app,
			defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500},
			nanolatherInfo:            &nanolatherInformation{builderCoeff: 1, allowedUnits: []string{"coreck", "coreak", "corethud"}},
		}
	 
		colors = []int{
			7, 7, 7,
			7, -1, 7,
			7, 7, 7}
		app = &buildingAppearance{chars: "" +
			"|=|" +
			"=2=" +
			"|=|", colors: colors}
			buildingsStaticDataMap["armt2kbotlab"] = &buildingStaticData{name: "Tech 2 Arm KBot Lab",
			w: 3, h: 3, appearance: app,
			defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 150, costM: 450, costE: 1200},
			nanolatherInfo:            &nanolatherInformation{builderCoeff: 1, allowedUnits: []string{"armfido"}},
		}
	
		colors = []int{
			7, 7, 7,
			7, -1, 7,
			7, 7, 7}
		app = &buildingAppearance{chars: "" +
			"\\=/" +
			">2<" +
			"/=\\", colors: colors}
			buildingsStaticDataMap["coret2kbotlab"] =  &buildingStaticData{name: "Tech 2 Core KBot Lab",
			w: 3, h: 3, appearance: app,
			defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 150, costM: 450, costE: 1200},
			nanolatherInfo:            &nanolatherInformation{builderCoeff: 1, allowedUnits: []string{"corethecan"}},
		}

	
		colors = []int{
			7, 7, 7, 7,
			7, -1, -1, 7,
			7, 7, 7, 7}
		app = &buildingAppearance{chars: "" +
			"====" +
			"|--|" +
			"\\==/", colors: colors}
		buildingsStaticDataMap["armvehfactory"] = &buildingStaticData{name: "Tech 1 Vehicle Factory", w: 4, h: 3, appearance: app,
			defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 120, costM: 175, costE: 700},
			nanolatherInfo:            &nanolatherInformation{builderCoeff: 1, allowedUnits: []string{"armcv", "armjeffy", "armflash"},
		}}

		colors = []int{
			7, 7, 7, 7,
			7, -1, -1, 7,
			7, -1, -1, 7}
		app = &buildingAppearance{chars: "" +
			"=--=" +
			"|/\\|" +
			"\\\\//", colors: colors}

			buildingsStaticDataMap["corevehfactory"] = &buildingStaticData{name: "Tech 1 Vehicle Factory", w: 4, h: 3, appearance: app,
			defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 125, costM: 180, costE: 700},
			nanolatherInfo:            &nanolatherInformation{builderCoeff: 1, allowedUnits: []string{"corecv", "coreweasel", "coreraider"}},
		}

	
		colors = []int{
			7, -1, 7,
			-1, 5, 5,
			7, -1, 7}
		app = &buildingAppearance{chars: "" +
			"=-=" +
			"|(-" +
			"=-=", colors: colors}
			buildingsStaticDataMap["radar"] =  &buildingStaticData{name: "Radar station", maxHitpoints: 25, isLight: true, radarRadius: 30,
			w: 3, h: 3, appearance: app,
			defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 100, costM: 250, costE: 450},
			res: &pawnResourceInformation{energyDrain: 15},
		}

	// buildingsStaticDataMap["lturret"] = 
	// 	colors = []int{-1}
	// 	app = &buildingAppearance{chars: "T", colors: colors}
	// 	b = &buildingStaticData{name: "Light Laser Turret", maxHitpoints: 90, isHeavy: true, regenPeriod: 20, sightRadius: 6,
	// 		w: 1, h: 1, allowsTightPlacement: true, appearance: app,
	// 		defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 100, costM: 250, costE: 900},
	// 		weapons: []*pawnWeaponInformation{
	// 			{attackDelay: 16, attackEnergyCost: 15, attackRadius: 6, attacksLand: true,
	// 				hitscan: &WeaponHitscan{baseDamage: 6},
	// 			},
	// 		},
	// 	}

	// buildingsStaticDataMap["guardian"] = 
	// 	colors = []int{
	// 		7, -1, 7,
	// 		-1, 5, -1,
	// 		7, -1, 7}
	// 	app = &buildingAppearance{chars: "" +
	// 		"=-=" +
	// 		"|&|" +
	// 		"=-=", colors: colors}
	// 	b = &buildingStaticData{name: "Guardian", maxHitpoints: 65, isHeavy: true,
	// 		w: 3, h: 3, appearance: app,
	// 		defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 100, costM: 250, costE: 900},
	// 		weapons: []*pawnWeaponInformation{
	// 			{attackDelay: 35, attackEnergyCost: 250, attackRadius: 12, attacksLand: true,
	// 				hitscan: &WeaponHitscan{baseDamage: 3, heavyMod: 5},
	// 			},
	// 		},
	// 	}

	// buildingsStaticDataMap["railgunturret"] = 
	// 	colors = []int{
	// 		7, -1, 5,
	// 		-1, 5, -1,
	// 		7, -1, 7}
	// 	app = &buildingAppearance{chars: "" +
	// 		"=-/" +
	// 		"|^|" +
	// 		"=-=", colors: colors}
	// 	b = &buildingStaticData{name: "Railgun Turret", maxHitpoints: 40, isHeavy: true,
	// 		w: 3, h: 3, appearance: app,
	// 		defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 100, costM: 450, costE: 1200},
	// 		weapons: []*pawnWeaponInformation{
	// 			{attackDelay: 55, attackEnergyCost: 150, attackRadius: 10, attacksLand: true,
	// 				hitscan: &WeaponHitscan{baseDamage: 5, heavyMod: 15},
	// 			},
	// 		},
	// 	}

	// buildingsStaticDataMap["wall"] = 
	// 	colors = []int{
	// 		-1}
	// 	app = &buildingAppearance{chars: "#", colors: colors}
	// 	b = &buildingStaticData{name: "Wall section", maxHitpoints: 140, isHeavy: true, regenPeriod: 9,
	// 		w: 1, h: 1, allowsTightPlacement: true, appearance: app,
	// 		defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 75, costM: 100, costE: 150},
	// 	}

	// colors = []int{
	// 	-1, -1,
	// 	-1, -1}
	// app = &buildingAppearance{chars: "" +
	// 	"??" +
	// 	"??", colors: colors}
	// buildingsStaticDataMap["DEFAULT"] = &buildingStaticData{name: "UNKNOWN BUILDING",
	// 	w: 2, h: 2, appearance: app,
	// 	defaultConstructionInfo: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500},
	// }

	// sanity check all the data
	for _, v := range buildingsStaticDataMap {
		if v.maxHitpoints == 0 {
			v.maxHitpoints = 25
			log.appendMessage("No hitpoints set for " + v.name)
		}
		if v.sightRadius == 0 {
			v.sightRadius = v.w + 2
			log.appendMessage("No visionRange for " + v.name)
		}
	}
}

func getBuildingStaticInfo(codename string) *buildingStaticData {
	newBuilding := buildingsStaticDataMap[codename]
	if newBuilding == nil {
		newBuilding = buildingsStaticDataMap["DEFAULT"]
	}
	return newBuilding
}
