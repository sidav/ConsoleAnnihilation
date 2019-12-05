package main 

var squadMembersStaticDataTable map[string]*squadMemberStaticData 

func initSquadMembersStaticDataMap() {
	squadMembersStaticDataTable = make(map[string]*squadMemberStaticData)
	squadMembersStaticDataTable["armcommander"] = &squadMemberStaticData {
		name: "Arm Commander",
		takesWholeSquad: true,
		movementInfo: &pawnMovementInformation {ticksForMoveSingleCell: 10, movesOnLand: true},
		maxHp: 100, sightRadius: 7, radarRadius: 15,
		weaponInfo: &pawnWeaponInformation {attackDelay: 10, attackEnergyCost: 1, attackRadius: 5, attacksLand: true,
			hitscan: &WeaponHitscan{baseDamage:5},
		},
		income:            &pawnIncomeInformation{metalIncome: 2, energyIncome: 30, metalStorage: 250, energyStorage: 1000},
		nanolatherInfo: &nanolatherInformation{builderCoeff: 10, allowedBuildings: []string{
			"armkbotlab", "armvehfactory", "solar", "mextractor", "metalmaker", "lturret", "wall"},
		},
		currentConstructionStatus: &constructionInformation{maxConstructionAmount: 10, costM: 250, costE: 500},
	}
	squadMembersStaticDataTable["DEFAULT"] = &squadMemberStaticData {
		name: "WTF",
		movementInfo: &pawnMovementInformation {ticksForMoveSingleCell: 10, movesOnLand: true},
		maxHp: 100,
		weaponInfo: &pawnWeaponInformation {attackDelay: 10, attackEnergyCost: 1, attackRadius: 5, attacksLand: true,
			hitscan: &WeaponHitscan{baseDamage:5},
		},
		income:            &pawnIncomeInformation{metalIncome: 2, energyIncome: 30, metalStorage: 250, energyStorage: 1000},
		nanolatherInfo: &nanolatherInformation{builderCoeff: 10, allowedBuildings: []string{
			"armkbotlab", "armvehfactory", "solar", "mextractor", "metalmaker", "lturret", "wall"},
		},
		currentConstructionStatus: &constructionInformation{maxConstructionAmount: 10, costM: 250, costE: 500},
	}
}

func getSquadMemberStaticInfo(codename string) *squadMemberStaticData {
	var newUnit *squadMemberStaticData
	newUnit = squadMembersStaticDataTable[codename]
	if newUnit == nil {
		newUnit = squadMembersStaticDataTable["DEFAULT"]
	}
	if newUnit.maxHp == 0 {
		newUnit.maxHp = 1
		log.appendMessage("No hitpoints set for "+newUnit.name)
	}
	// newUnit.hitpoints = newUnit.maxHitpoints
	// newUnit.x = x
	// newUnit.y = y
	// newUnit.faction = f
	// newUnit.codename = codename
	// if newUnit.sightRadius == 0 {
	// 	newUnit.sightRadius = 8
	// }
	// if alreadyConstructed {
	// 	newUnit.currentConstructionStatus = nil
	// }
	return newUnit
}

func getSquadmemberNameAndDescription(code string) (string, string) {
	unit := getSquadMemberStaticInfo(code)
	name := unit.name
	var description string
	// if unit.currentConstructionStatus != nil {
	// 	constr := unit.currentConstructionStatus
	// 	description += constr.getDescriptionString() + " \\n "
	// }
	// if len(unit.weapons) > 0 {
	// 	for _, wpn := range unit.weapons {
	// 		description += wpn.getDescriptionString() + " \\n "
	// 	}
	// }
	switch code {
	case "armcommander":
		description += "A Commander Unit of the Arm Rebellion."
	case "corecommander":
		description += "A Commander Unit of the Core Corporation."
	case "protocommander":
		description += "The brain and heart of any modern military operation, Command Unit is a massive bipedal amphibious hulk " +
			"equipped with anything needed for forward operations base establishment and support - a radar, nanolather, " +
			"quantum generator, metal synthesizer, and light weaponry. " +
			"Although the Command Unit is extremely heavily armored and is capable of self-repair, it have to be well protected, " +
			"because its loss means inevitable defeat. \\n " +
			"This old prototype lacks some more modern equipment, such as the Disintegrator Gun."
	case "coreck":
		description += "An engineering KBot equipped with nanolather. Can build more advanced buildings than Commander do."
	case "corecv":
		description += "An engineering vehicle equipped with nanolather. Can build more advanced buildings than Commander do."
	case "coreweasel":
		description += "Fast recon vehicle. It has very weak attack, but is equipped with advanced visual sensors array " +
			"which is providing quite huge vision range."
	case "armjeffy":
		description += "Fast recon vehicle. It has very weak attack, but is equipped with advanced visual sensors array " +
			"which is providing quite huge vision range."
	case "coreraider":
		description += "Fast light tank. Its speed and armor regeneration ability makes it useful for hit-and-run tactics."
	case "armflash":
		description +=  "Fast light tank. Its speed and armor regeneration ability makes it useful for hit-and-run tactics."
	case "coreak":
		description += "A basic assault KBot effective against light armor."
	case "armpeewee":
		description += "A cheap and relatively fast basic assault KBot effective against light armor."
	case "armhammer":
		description += "A basic artillery KBot. Effective against heavy armor. Designed to take out buildings. "
	case "corethud":
		description += "A basic artillery KBot. Effective against heavy armor. Designed to take out buildings. "
	case "corethecan":
		description += "Slow and clunky, The Can is designed to take part in front-line assault. Although its " +
			"armor can sustain significant amount of punishment, this KBot should be supported due to its short range."
	default:
		description += "No description."
	}
	return name, description
}
