package main

// TODO: move unit stats to JSON
func createUnit(codename string, x, y int, f *faction, alreadyConstructed bool) *pawn {
	var newUnit *pawn
	switch codename {
	case "armcommander":
		newUnit = &pawn{name: "Arm Commander", maxHitpoints: 200, isHeavy: true, isCommander: true,
			unitInfo:       &unit{appearance: ccell{char: '@'}},
			moveInfo:       &pawnMovementInformation{ticksForMoveSingleCell: 10, movesOnLand: true, movesOnSea: true}, regenPeriod: 7, radarRadius: 15,
			res:            &pawnResourceInformation{metalIncome: 1, energyIncome: 20, metalStorage: 250, energyStorage: 1000},
			nanolatherInfo: &nanolatherInformation{builderCoeff: 10, allowedBuildings: []string{}},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 10, attackEnergyCost: 1, attackRadius: 6, attacksLand: true,
					hitscan: &WeaponHitscan{baseDamage:4},
				},
			},
		}
	case "corecommander":
		newUnit = &pawn{name: "Core Commander", maxHitpoints: 200, isHeavy: true, isCommander: true,
			unitInfo:       &unit{appearance: ccell{char: '@'}},
			moveInfo:       &pawnMovementInformation{ticksForMoveSingleCell: 10, movesOnLand: true, movesOnSea: true}, regenPeriod: 7, radarRadius: 15,
			res:            &pawnResourceInformation{metalIncome: 1, energyIncome: 20, metalStorage: 250, energyStorage: 1000},
			nanolatherInfo: &nanolatherInformation{builderCoeff: 10, allowedBuildings: []string{"corekbotlab", "solar", "metalmaker", "corevehfactory", "lturret"}},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 10, attackEnergyCost: 1, attackRadius: 5, attacksLand: true,
					hitscan: &WeaponHitscan{baseDamage:5},
				},
			},
		}
	case "protocommander":
		newUnit = &pawn{name: "Prototype Commander Unit", maxHitpoints: 200, isHeavy: true, isCommander: true, regenPeriod: 7, radarRadius: 15,
			unitInfo:       &unit{appearance: ccell{char: '@'}},
			moveInfo:       &pawnMovementInformation{ticksForMoveSingleCell: 10, movesOnLand: true, movesOnSea: true},
			res:            &pawnResourceInformation{metalIncome: 1, energyIncome: 20, metalStorage: 250, energyStorage: 1000},
			nanolatherInfo: &nanolatherInformation{builderCoeff: 10, allowedBuildings: []string{
				"corekbotlab", "corevehfactory", "solar", "mextractor", "metalmaker", "lturret", "wall", "bunkerwithmarines"},
			},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 10, attackEnergyCost: 1, attackRadius: 5, attacksLand: true,
					hitscan: &WeaponHitscan{baseDamage:5},
				},
			},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 15000, costM: 1567650, costE: 59120150},
		}
	case "coreck":
		newUnit = &pawn{name: "Tech 1 Construction KBot", maxHitpoints: 25, isLight: true,
			unitInfo:                  &unit{appearance: ccell{char: 'k'}},
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 15, movesOnLand: true},
			res:                       &pawnResourceInformation{metalStorage: 25, energyStorage: 50},
			nanolatherInfo:            &nanolatherInformation{builderCoeff: 5, allowedBuildings: []string{
				"corekbotlab", "coret2kbotlab", "mstorage", "estorage", "mextractor", "solar", "metalmaker", "lturret", "railgunturret", "geo", "radar",},
			},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 25, costM: 650, costE: 1200},
		}
	case "corecv":
		newUnit = &pawn{name: "Tech 1 Construction Vehicle", maxHitpoints: 25, isLight: true,
			unitInfo:                  &unit{appearance: ccell{char: 'v'}},
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 15, movesOnLand: true},
			res:                       &pawnResourceInformation{metalStorage: 25, energyStorage: 50},
			nanolatherInfo:            &nanolatherInformation{builderCoeff: 5, allowedBuildings: []string{
				"corevehfactory", "mstorage", "estorage", "mextractor", "solar", "metalmaker", "lturret", "railgunturret", "geo", "radar",},
			},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 25, costM: 650, costE: 1200},
		}
	case "coreweasel":
		newUnit = &pawn{name: "Weasel", maxHitpoints: 20, isLight: true, sightRadius: 12, regenPeriod: 30,
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 6, movesOnLand: true},
			unitInfo:                  &unit{appearance: ccell{char: 'w'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 12, costM: 150, costE: 650},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 7, attackEnergyCost: 1, attackRadius: 4, attacksLand: true, canBeFiredOnMove: true,
					hitscan: &WeaponHitscan{baseDamage:2},
				},
			},
		}
	case "armjeffy":
		newUnit = &pawn{name: "Jeffy", maxHitpoints: 20, isLight: true, sightRadius: 12, regenPeriod: 25,
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 5, movesOnLand: true},
			unitInfo:                  &unit{appearance: ccell{char: 'j'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 12, costM: 150, costE: 650},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 7, attackEnergyCost: 1, attackRadius: 4, attacksLand: true, canBeFiredOnMove: true,
					hitscan: &WeaponHitscan{baseDamage:2, lightMod:2},
				},
			},
		}
	case "coreraider":
		newUnit = &pawn{name: "Raider", maxHitpoints: 55, isLight: true, regenPeriod: 6,
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 7, movesOnLand: true},
			unitInfo:                  &unit{appearance: ccell{char: 'r'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 18, costM: 350, costE: 600},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 13, attackEnergyCost: 1, attackRadius: 4, attacksLand: true, canBeFiredOnMove: true,
					hitscan: &WeaponHitscan{baseDamage:4},
				},
			},
		}
	case "armflash":
		newUnit = &pawn{name: "Flash", maxHitpoints: 55, isLight: true, regenPeriod: 6,
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 7, movesOnLand: true},
			unitInfo:                  &unit{appearance: ccell{char: 'f'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 18, costM: 350, costE: 600},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 13, attackEnergyCost: 1, attackRadius: 4, attacksLand: true, canBeFiredOnMove: true,
					hitscan: &WeaponHitscan{baseDamage:4},
				},
			},
		}
	case "corethecan":
		newUnit = &pawn{name: "The Can", maxHitpoints: 125, isHeavy: true,
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 10, movesOnLand: true},
			unitInfo:                  &unit{appearance: ccell{char: 'c'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 25, costM: 250, costE: 500},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 13, attackEnergyCost: 1, attackRadius: 4, attacksLand: true,
					hitscan: &WeaponHitscan{baseDamage:4},
				},
			},
		}
	case "coreak":
		newUnit = &pawn{name: "A.K.", maxHitpoints: 30,
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 10, movesOnLand: true},
			unitInfo:                  &unit{appearance: ccell{char: 'a'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 14, costM: 75, costE: 500},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 7, attackEnergyCost: 1, attackRadius: 5, attacksLand: true, canBeFiredOnMove: true,
					hitscan: &WeaponHitscan{baseDamage:3},
				},
			},
		}
	case "armpeewee":
		newUnit = &pawn{name: "P.I.V.-1", maxHitpoints: 35, isLight: true,
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 10, movesOnLand: true},
			unitInfo:                  &unit{appearance: ccell{char: 'p'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 14, costM: 75, costE: 500},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 6, attackEnergyCost: 1, attackRadius: 5, attacksLand: true, canBeFiredOnMove: true,
					hitscan: &WeaponHitscan{baseDamage:3, lightMod:3},
				},
			},
		}
	case "corethud":
		newUnit = &pawn{name: "Thud", maxHitpoints: 35, isHeavy: true,
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 16, movesOnLand: true},
			unitInfo:                  &unit{appearance: ccell{char: 't'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 22, costM: 350, costE: 650},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 13, attackEnergyCost: 1, attackRadius: 7, attacksLand: true,
					hitscan: &WeaponHitscan{baseDamage:3, heavyMod:4},
				},
			},
		}
	case "armhammer":
		newUnit = &pawn{name: "Hammer", maxHitpoints: 25, isHeavy: true,
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 13, movesOnLand: true},
			unitInfo:                  &unit{appearance: ccell{char: 'h'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 22, costM: 250, costE: 500},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 14, attackEnergyCost: 1, attackRadius: 6, attacksLand: true, canBeFiredOnMove: false,
					hitscan: &WeaponHitscan{baseDamage:3, heavyMod:3},
				},
			},
		}
	default:
		newUnit = &pawn{name: "UNKNOWN UNIT " + codename,
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 10, movesOnLand: true},
			unitInfo:                  &unit{appearance: ccell{char: '?'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 10, costM: 250, costE: 500},
		}
	}
	if newUnit.maxHitpoints == 0 {
		newUnit.maxHitpoints = 1
		log.appendMessage("No hitpoints set for "+newUnit.name)
	}
	newUnit.hitpoints = newUnit.maxHitpoints
	newUnit.x = x
	newUnit.y = y
	newUnit.faction = f
	newUnit.codename = codename
	if newUnit.sightRadius == 0 {
		newUnit.sightRadius = 5
	}
	if alreadyConstructed {
		newUnit.currentConstructionStatus = nil
	}
	return newUnit
}

func getUnitNameAndDescription(code string) (string, string) {
	unit := createUnit(code, 0, 0, nil, false)
	name := unit.name
	var description string
	if unit.currentConstructionStatus != nil {
		constr := unit.currentConstructionStatus
		description += constr.getDescriptionString() + " \\n "
	}
	if len(unit.weapons) > 0 {
		for _, wpn := range unit.weapons {
			description += wpn.getDescriptionString() + " \\n "
		}
	}
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
