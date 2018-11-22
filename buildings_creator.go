package main

func createBuilding(name string, x, y int, f *faction) *building {
	var b *building
	switch name {
	case "kbotlab":
		colors := []int {
			7, 7, 7,
			7, -1, 7,
			7, 7, 7}
		app := &buildingAppearance{chars: "" +
			"\\=/" +
			"=0=" +
			"/|\\", colors: colors}
		b = &building{name: "Tech 1 KBot Lab", w: 3, h: 3, appearance: app}
	case "vehfactory":
		colors := []int {
			7, 7, 7, 7,
			7, -1, -1, 7,
			7, 7, 7, 7}
		app := &buildingAppearance{chars: "" +
			"====" +
			"|--|" +
			"=\\/=", colors: colors}
		b = &building{name: "Tech 1 Vehicle Factory", w: 4, h: 3, appearance: app}
	}
	b.x = x
	b.y = y
	b.faction = f
	return b
}
