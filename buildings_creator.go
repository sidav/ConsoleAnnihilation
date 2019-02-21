package main

func createBuilding(codename string, x, y int, f *faction) *pawn {
	var b *pawn
	switch codename {

	case "metalmaker":
		colors := []int{
			7, 7,
			7, -1}
		app := &buildingAppearance{chars: "" +
			"xx" +
			"xx", colors: colors}
		b = &pawn{name: "Metal Synthesizer",
			buildingInfo:              &building{w: 2, h: 2, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 35, costM: 10, costE: 500},
			res:                       &pawnResourceInformation{metalIncome: 1, energyDrain: 60},
		}

	case "solar":
		colors := []int{
			-1, 7,
			7, -1}
		app := &buildingAppearance{chars: "" +
			"==" +
			"==", colors: colors}
		b = &pawn{name: "Solar Collector",
			buildingInfo:              &building{w: 2, h: 2, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 60, costM: 100, costE: 500},
			res:                       &pawnResourceInformation{energyIncome: 20},
		}

	case "mextractor":
		colors := []int{
			-1, 7, -1,
			7, -1, 7,
			-1, 7, -1,
		}
		app := &buildingAppearance{chars:
			"#|#" +
			"-%-" +
			"#|#", colors: colors}
		b = &pawn{name: "Metal Extractor",
			buildingInfo:              &building{w: 3, h: 3, appearance: app, canBeBuiltOnMetalOnly: true},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 60, costM: 100, costE: 500},
			res:                       &pawnResourceInformation{energyDrain: 9, isMetalExtractor: true},
		}

	case "geo":
		colors := []int{
			7, 13, 13, 7,
			-1, 8, 8, -1,
			-1, 8, 8, -1,
			7, 13, 13, 7}
		app := &buildingAppearance{chars: "" +
			"=^^=" +
			"{00}" +
			"{00}" +
			"=VV=", colors: colors}
		b = &pawn{name: "Geothermal Powerplant",
			buildingInfo:              &building{w: 4, h: 4, appearance: app, canBeBuiltOnThermalOnly: true},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 60, costM: 350, costE: 1500},
			res:                       &pawnResourceInformation{energyIncome: 250},
		}

	case "mstorage":
		colors := []int{
			-1, 7, -1,
			7, -1, 7,
			-1, 7, -1,
		}
		app := &buildingAppearance{chars: "" +
			"=0=" +
			"0$0" +
			"=0=", colors: colors}
		b = &pawn{name: "Metal Storage", maxHitpoints: 100, isHeavy: true, regenPeriod: 10,
			buildingInfo:              &building{w: 3, h: 3, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 60, costM: 100, costE: 800},
			res:                       &pawnResourceInformation{metalStorage: 250},
		}

	case "estorage":
		colors := []int{
			-1, 7, -1,
			7, -1, 7,
			-1, 7, -1,
		}
		app := &buildingAppearance{chars: "" +
			"=|=" +
			"-0-" +
			"=|=", colors: colors}
		b = &pawn{name: "Energy Storage",
			buildingInfo:              &building{w: 3, h: 3, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 60, costM: 200, costE: 400},
			res:                       &pawnResourceInformation{energyStorage: 250},
		}

	case "quark": // cheating building, useful for debugging
		colors := []int{
			7, -1, 7,
			-1, 13, -1,
			7, -1, 7}
		app := &buildingAppearance{chars: "" +
			"=^=" +
			"<O>" +
			"=V=", colors: colors}
		b = &pawn{name: "Quark Antimatter Generator",
			buildingInfo:              &building{w: 3, h: 3, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 40, costM: 100, costE: 1200},
			res:                       &pawnResourceInformation{metalIncome: 5, energyIncome: 100},
		}

	case "armhq": // cheating building for stub AI
		colors := []int{
			7, 13, 13, 7,
			13, -1, -1, 13,
			7, 13, 13, 7}
		app := &buildingAppearance{chars: "" +
			"=^^=" +
			"<HQ>" +
			"=VV=", colors: colors}
		b = &pawn{name: "Arm proxy HQ", maxHitpoints: 300, isHeavy: true, isCommander: true, regenPeriod: 10, radarRadius: 40,
			buildingInfo:              &building{w: 4, h: 3, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 40, costM: 100, costE: 1200},
			res:                       &pawnResourceInformation{metalIncome: 100, energyIncome: 1000},
			nanolatherInfo:            &nanolatherInformation{builderCoeff: 1, allowedUnits: []string{"armpeewee", "armhammer"}},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 12, attackEnergyCost: 15, attackRadius: 6, attacksLand: true,
					hitscan: &WeaponHitscan{baseDamage: 6},
				},
			},
		}

	case "armkbotlab":
		colors := []int{
			7, 7, 7,
			7, -1, 7,
			7, 7, 7}
		app := &buildingAppearance{chars: "" +
			"/=\\" +
			"=x=" +
			"\\=/", colors: colors}
		b = &pawn{name: "Tech 1 KBot Lab",
			buildingInfo:              &building{w: 3, h: 3, appearance: app},
			nanolatherInfo:            &nanolatherInformation{builderCoeff: 1, allowedUnits: []string{"armpeewee", "armhammer"}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500},
		}

	case "corekbotlab":
		colors := []int{
			7, 7, 7,
			7, -1, 7,
			7, 7, 7}
		app := &buildingAppearance{chars: "" +
			"\\=/" +
			"=0=" +
			"/=\\", colors: colors}
		b = &pawn{name: "Tech 1 Core KBot Lab",
			buildingInfo:              &building{w: 3, h: 3, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500},
			nanolatherInfo:            &nanolatherInformation{builderCoeff: 1, allowedUnits: []string{"coreck", "ak", "thud"}},
		}

	case "coret2kbotlab":
		colors := []int{
			7, 7, 7,
			7, -1, 7,
			7, 7, 7}
		app := &buildingAppearance{chars: "" +
			"\\=/" +
			">2<" +
			"/=\\", colors: colors}
		b = &pawn{name: "Tech 2 Core KBot Lab",
			buildingInfo:              &building{w: 3, h: 3, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 150, costM: 450, costE: 1200},
			nanolatherInfo:            &nanolatherInformation{builderCoeff: 1, allowedUnits: []string{"thecan"}},
		}

	case "armvehfactory":
		colors := []int{
			7, 7, 7, 7,
			7, -1, -1, 7,
			7, 7, 7, 7}
		app := &buildingAppearance{chars: "" +
			"====" +
			"|--|" +
			"\\==/", colors: colors}
		b = &pawn{name: "Tech 1 Vehicle Factory",
			buildingInfo:              &building{w: 4, h: 3, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500},
		}

	case "corevehfactory":
		colors := []int{
			7, 7, 7, 7,
			7, -1, -1, 7,
			7, -1, -1, 7}
		app := &buildingAppearance{chars: "" +
			"=--=" +
			"|/\\|" +
			"\\\\//", colors: colors}
		b = &pawn{name: "Tech 1 Vehicle Factory", buildingInfo: &building{w: 4, h: 3, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500},
			nanolatherInfo:            &nanolatherInformation{builderCoeff: 1, allowedUnits: []string{"weasel", "flash"}},
		}

	case "radar":
		colors := []int{
			7, -1, 7,
			-1, 5, 5,
			7, -1, 7}
		app := &buildingAppearance{chars: "" +
			"=-=" +
			"|(-" +
			"=-=", colors: colors}
		b = &pawn{name: "Radar station", maxHitpoints: 25, isLight: true, radarRadius: 30,
			buildingInfo:              &building{w: 3, h: 3, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 250, costE: 450},
			res: &pawnResourceInformation{energyDrain: 15},
		}

	case "lturret":
		colors := []int{-1}
		app := &buildingAppearance{chars: "T", colors: colors}
		b = &pawn{name: "Light Laser Turret", maxHitpoints: 90, isHeavy: true, regenPeriod: 70,
			buildingInfo:              &building{w: 1, h: 1, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 250, costE: 900},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 16, attackEnergyCost: 15, attackRadius: 6, attacksLand: true,
					hitscan: &WeaponHitscan{baseDamage: 6},
				},
			},
		}

	case "guardian":
		colors := []int{
			7, -1, 7,
			-1, 5, -1,
			7, -1, 7}
		app := &buildingAppearance{chars: "" +
			"=-=" +
			"|&|" +
			"=-=", colors: colors}
		b = &pawn{name: "Guardian", maxHitpoints: 65, isHeavy: true,
			buildingInfo:              &building{w: 3, h: 3, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 250, costE: 900},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 35, attackEnergyCost: 250, attackRadius: 12, attacksLand: true,
					hitscan: &WeaponHitscan{baseDamage: 3, heavyMod: 5},
				},
			},
		}

	case "railgunturret":
		colors := []int{
			7, -1, 5,
			-1, 5, -1,
			7, -1, 7}
		app := &buildingAppearance{chars: "" +
			"=-/" +
			"|^|" +
			"=-=", colors: colors}
		b = &pawn{name: "Railgun Turret", maxHitpoints: 40, isHeavy: true,
			buildingInfo:              &building{w: 3, h: 3, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 450, costE: 1200},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 55, attackEnergyCost: 150, attackRadius: 10, attacksLand: true,
					hitscan: &WeaponHitscan{baseDamage: 5, heavyMod: 15},
				},
			},
		}

	default:
		colors := []int{
			-1, -1,
			-1, -1}
		app := &buildingAppearance{chars: "" +
			"??" +
			"??", colors: colors}
		b = &pawn{name: "UNKNOWN BUILDING",
			buildingInfo:              &building{w: 2, h: 2, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500},
		}
	}
	if b.maxHitpoints == 0 {
		b.maxHitpoints = 25
	}
	b.hitpoints = b.maxHitpoints
	b.x = x
	b.y = y
	b.faction = f
	b.codename = codename
	if b.sightRadius == 0 {
		b.sightRadius = 1
	}
	if b.nanolatherInfo != nil && b.res == nil {
		b.res = &pawnResourceInformation{} // adds zero-value resource info struct for spendings usage.
	}
	return b
}

func getBuildingNameAndDescription(code string) (string, string) {
	bld := createBuilding(code, 0, 0, nil)
	name := bld.name
	constr := bld.currentConstructionStatus
	description := constr.getDescriptionString() + " \\n "
	if len(bld.weapons) > 0 {
		for _, wpn := range bld.weapons {
			description += wpn.getDescriptionString() + " \\n "
		}
	}
	switch code {
	case "armkbotlab", "corekbotlab":
		description += "A basic nanolathing facility which is designed to construct the Kinetic Bio-Organic Technology " +
			"mechs, or KBots. "
	case "solar":
		description += "A classic solar battery array. The heavy use of superconductors and wireless energy transfer technologies " +
			"made this energy acqurement devices much more efficient than ever."
	case "lturret":
		description += "A basic yet quite universal base defense structure. Its only weapon uses EM-waves amplified by stimulated emission of radiation."
	default:
		description += "No description."
	}
	return name, description
}
