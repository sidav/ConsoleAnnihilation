package main

func createUnit(name string, x, y, faction int) *unit {
	switch name {
	case "Commander":
		return &unit{faction: faction, x: x, y: y, appearance:ccell{char:'@'}}
	}
	return nil
}
