package main

// TODO: move unit stats to JSON
func createUnit(name string, x, y int, f *faction) *pawn {
	var newUnit *pawn
	switch name {
	case "commander":
		newUnit = &pawn{ name: "Commander",
			unitInfo: &unit{
			 ticksForMoveOneCell: 10, appearance: ccell{char: '@'}},
			res:            &pawnResourceInformation{metalIncome: 1, energyIncome: 10, metalStorage: 100, energyStorage: 1000},
			nanolatherInfo: &nanolatherInformation{builderCoeff: 10, allowedBuildings: []string{"corekbotlab", "solar", "metalmaker"} },
		}
	//case "weasel":
	//	newUnit = &unit{name: "Weasel", ticksForMoveOneCell: 6, appearance: ccell{char: 'w'}}
	//case "thecan":
	//	newUnit = &unit{name: "The Can", ticksForMoveOneCell: 17, appearance: ccell{char: 'c'}}
	//case "ak":
	//	newUnit = &unit{name: "A.K.", ticksForMoveOneCell: 9, appearance: ccell{char: 'a'}}
	default:
		newUnit = &pawn {name: "UNKNOWN UNIT",
			unitInfo: &unit{appearance: ccell{char: '?'}},
		}
	}
	newUnit.x = x
	newUnit.y = y
	newUnit.faction = f
	return newUnit
}
