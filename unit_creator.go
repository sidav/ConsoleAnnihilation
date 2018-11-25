package main

// TODO: move unit stats to JSON
func createUnit(name string, x, y int, f *faction) *unit {
	var newUnit *unit
	switch name {
	case "commander":
		newUnit =  &unit{
			name: "Commander", ticksForMoveOneCell: 10, appearance: ccell{char: '@'},
			res: &resourceInformation{1, 10, 10, 100, 0},
		}
	case "weasel":
		newUnit =  &unit{name: "Weasel", ticksForMoveOneCell: 6, appearance: ccell{char: 'w'}}
	case "thecan":
		newUnit =  &unit{name: "The Can", ticksForMoveOneCell: 17, appearance: ccell{char: 'c'}}
	case "ak":
		newUnit =  &unit{name: "A.K.", ticksForMoveOneCell: 9, appearance: ccell{char: 'a'}}
	default:
		newUnit =  &unit{name: "UNKNOWN UNIT", faction: f, x: x, y: y, appearance: ccell{char: '?'}}
	}
	newUnit.x = x
	newUnit.y = y
	newUnit.faction = f
	return newUnit
}
