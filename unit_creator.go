package main

func createUnit(name string, x, y, faction int) *unit {
	switch name {
	case "commander":
		return &unit{name: name, faction: faction, x: x, y: y, appearance:ccell{char:'@'}}
	}
	return &unit{name: "UNKNOWN UNIT", faction: faction, x: x, y: y, appearance:ccell{char:'?'}}
}
