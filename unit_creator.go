package main

// TODO: move unit stats to JSON
func createUnit(name string, x, y int, f *faction, alreadyConstructed bool) *pawn {
	var newUnit *pawn
	switch name {
	case "commander":
		newUnit = &pawn{ name: "Commander",
			unitInfo: &unit{
			 ticksForMoveOneCell: 10, appearance: ccell{char: '@'}},
			res:            &pawnResourceInformation{metalIncome: 1, energyIncome: 10, metalStorage: 1000, energyStorage: 1000},
			nanolatherInfo: &nanolatherInformation{builderCoeff: 10, allowedBuildings: []string{"corekbotlab", "solar", "metalmaker", "corevehfactory", "quark"} },
		}
	case "weasel":
		newUnit = &pawn{name: "Weasel",
		unitInfo: &unit{ticksForMoveOneCell: 6, appearance: ccell{char: 'w'}},
		}
	case "thecan":
		newUnit = &pawn {name: "The Can",
		unitInfo: &unit{ticksForMoveOneCell: 17, appearance: ccell{char: 'c'}},
		}
	case "ak":
		newUnit = &pawn {name: "A.K.",
		unitInfo: &unit{ticksForMoveOneCell: 9, appearance: ccell{char: 'a'}},
		currentConstructionStatus: &constructionInformation{maxConstructionAmount: 10, costM: 250, costE: 500},
		}
	default:
		newUnit = &pawn {name: "UNKNOWN UNIT",
			unitInfo: &unit{appearance: ccell{char: '?'}},
		}
	}
	newUnit.x = x
	newUnit.y = y
	newUnit.faction = f
	if alreadyConstructed {
		newUnit.currentConstructionStatus = nil
	}
	return newUnit
}
