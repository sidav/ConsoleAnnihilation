package main

// TODO: move unit stats to JSON
func createUnit(codename string, x, y int, f *faction, alreadyConstructed bool) *pawn {
	var newUnit *pawn
	switch codename {
	case "armcommander":
		newUnit = &pawn{name: "Arm Commander", maxHitpoints: 100, isHeavy: true, isCommander: true,
			unitInfo:       &unit{appearance: ccell{char: '@'}},
			moveInfo:       &pawnMovementInformation{ticksForMoveSingleCell: 10, movesOnLand: true, movesOnSea: true}, eachTickToRegen: 7,
			res:            &pawnResourceInformation{metalIncome: 1, energyIncome: 20, metalStorage: 250, energyStorage: 1000},
			nanolatherInfo: &nanolatherInformation{builderCoeff: 10, allowedBuildings: []string{}},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 10, attackEnergyCost: 1, attackRadius: 6, attacksLand: true,
					hitscan: &WeaponHitscan{baseDamage:4},
				},
			},
		}
	case "corecommander":
		newUnit = &pawn{name: "Core Commander", maxHitpoints: 100, isHeavy: true, isCommander: true,
			unitInfo:       &unit{appearance: ccell{char: '@'}},
			moveInfo:       &pawnMovementInformation{ticksForMoveSingleCell: 10, movesOnLand: true, movesOnSea: true}, eachTickToRegen: 7,
			res:            &pawnResourceInformation{metalIncome: 1, energyIncome: 20, metalStorage: 250, energyStorage: 1000},
			nanolatherInfo: &nanolatherInformation{builderCoeff: 10, allowedBuildings: []string{"corekbotlab", "solar", "metalmaker", "corevehfactory", "lturret"}},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 10, attackEnergyCost: 1, attackRadius: 5, attacksLand: true,
					hitscan: &WeaponHitscan{baseDamage:5},
				},
			},
		}
	case "protocommander":
		newUnit = &pawn{name: "Prototype Commander Unit", maxHitpoints: 100, isHeavy: true, isCommander: true, eachTickToRegen: 7,
			unitInfo:       &unit{appearance: ccell{char: '@'}},
			moveInfo:       &pawnMovementInformation{ticksForMoveSingleCell: 10, movesOnLand: true, movesOnSea: true},
			res:            &pawnResourceInformation{metalIncome: 1, energyIncome: 20, metalStorage: 250, energyStorage: 1000},
			nanolatherInfo: &nanolatherInformation{builderCoeff: 10, allowedBuildings: []string{"corekbotlab", "solar", "mextractor", "metalmaker", "corevehfactory", "lturret", "geo", "radar"}},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 10, attackEnergyCost: 1, attackRadius: 5, attacksLand: true,
					hitscan: &WeaponHitscan{baseDamage:5},
				},
			},
		}
	case "coreck":
		newUnit = &pawn{name: "Tech 1 Construction KBot",
			unitInfo:                  &unit{appearance: ccell{char: 'k'}},
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 15, movesOnLand: true},
			res:                       &pawnResourceInformation{metalStorage: 25, energyStorage: 50},
			nanolatherInfo:            &nanolatherInformation{builderCoeff: 5, allowedBuildings: []string{"corekbotlab", "coret2kbotlab", "mstorage", "estorage", "solar", "metalmaker", "quark"}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 25, costM: 650, costE: 1200},
		}
	case "weasel":
		newUnit = &pawn{name: "Weasel", maxHitpoints: 20, isLight: true, sightRadius: 12, eachTickToRegen: 30,
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 6, movesOnLand: true},
			unitInfo:                  &unit{appearance: ccell{char: 'w'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 10, costM: 150, costE: 650},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 7, attackEnergyCost: 1, attackRadius: 4, attacksLand: true, canBeFiredOnMove: true,
					hitscan: &WeaponHitscan{baseDamage:2},
				},
			},
		}
	case "flash":
		newUnit = &pawn{name: "Flash", maxHitpoints: 55, isLight: true, eachTickToRegen: 6,
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 7, movesOnLand: true},
			unitInfo:                  &unit{appearance: ccell{char: 'f'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 12, costM: 350, costE: 600},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 13, attackEnergyCost: 1, attackRadius: 4, attacksLand: true, canBeFiredOnMove: true,
					hitscan: &WeaponHitscan{baseDamage:4},
				},
			},
		}
	case "thecan":
		newUnit = &pawn{name: "The Can", maxHitpoints: 125, isHeavy: true,
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 10, movesOnLand: true},
			unitInfo:                  &unit{appearance: ccell{char: 'c'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 10, costM: 250, costE: 500},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 13, attackEnergyCost: 1, attackRadius: 4, attacksLand: true,
					hitscan: &WeaponHitscan{baseDamage:4},
				},
			},
		}
	case "ak":
		newUnit = &pawn{name: "A.K.", maxHitpoints: 30,
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 10, movesOnLand: true},
			unitInfo:                  &unit{appearance: ccell{char: 'a'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 10, costM: 250, costE: 500},
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
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 10, costM: 250, costE: 500},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 6, attackEnergyCost: 1, attackRadius: 5, attacksLand: true, canBeFiredOnMove: true,
					hitscan: &WeaponHitscan{baseDamage:3, lightMod:3},
				},
			},
		}
	case "thud":
		newUnit = &pawn{name: "Thud", maxHitpoints: 35, isHeavy: true,
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 16, movesOnLand: true},
			unitInfo:                  &unit{appearance: ccell{char: 't'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 12, costM: 350, costE: 650},
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
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 12, costM: 250, costE: 500},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 14, attackEnergyCost: 1, attackRadius: 6, attacksLand: true, canBeFiredOnMove: false,
					hitscan: &WeaponHitscan{baseDamage:3, heavyMod:3},
				},
			},
		}
	default:
		newUnit = &pawn{name: "UNKNOWN UNIT",
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
	constr := unit.currentConstructionStatus
	description := constr.getDescriptionString() + " \\n "
	if len(unit.weapons) > 0 {
		for _, wpn := range unit.weapons {
			description += wpn.getDescriptionString() + " \\n "
		}
	}
	switch code {
	case "ak":
		description += "A basic assault KBot effective against light armor."
	case "thud":
		description += "A basic artillery KBot. Effective against heavy armor. Designed to take out buildings. "
	case "thecan":
		description += "Slow and clunky, The Can is designed to take part in front-line assault. Although its " +
			"armor can sustain significant amount of punishment, this KBot should be supported due to its short range."
	default:
		description += "No description."
	}
	return name, description
}
