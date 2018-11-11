package main

func createUnit(name string, x, y int, f *faction) *unit {
	switch name {
	case "commander":
		return &unit{name: name, faction: f, x: x, y: y, appearance:ccell{char:'@'}}
	}
	return &unit{name: "UNKNOWN UNIT", faction: f, x: x, y: y, appearance:ccell{char:'?'}}
}
