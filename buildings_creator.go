package main

func createBuilding(name string, x, y int, f *faction) *building {
	var b *building
	switch name {

	case "metalmaker":
		colors := []int{
			7, 7,
			7, -1}
		app := &buildingAppearance{chars: "" +
			"xx" +
			"xx" , colors: colors}
		b = &building{
			name: "Metal Synthesizer", w: 2, h: 2, appearance: app,
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500},
			res: &pawnResourceInformation{metalIncome: 1, energyReqForConditionalMetalIncome: 60},
		}

	case "armkbotlab":
		colors := []int{
			7, 7, 7,
			7, -1, 7,
			7, 7, 7}
		app := &buildingAppearance{chars: "" +
			"/=\\" +
			"=x=" +
			"\\=/", colors: colors}
		b = &building{name: "Tech 1 KBot Lab", w: 3, h: 3, appearance: app,
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500}}

	case "corekbotlab":
		colors := []int{
			7, 7, 7,
			7, -1, 7,
			7, 7, 7}
		app := &buildingAppearance{chars: "" +
			"\\=/" +
			"=0=" +
			"/=\\", colors: colors}
		b = &building{name: "Tech 1 KBot Lab", w: 3, h: 3, appearance: app,
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500}}

	case "armvehfactory":
		colors := []int{
			7, 7, 7, 7,
			7, -1, -1, 7,
			7, 7, 7, 7}
		app := &buildingAppearance{chars: "" +
			"====" +
			"|--|" +
			"\\==/", colors: colors}
		b = &building{name: "Tech 1 Vehicle Factory", w: 4, h: 3, appearance: app,
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500}}

	case "corevehfactory":
		colors := []int{
			7, 7, 7, 7,
			7, -1, -1, 7,
			7, -1, -1, 7}
		app := &buildingAppearance{chars: "" +
			"=--=" +
			"|/\\|" +
			"\\\\//", colors: colors}
		b = &building{name: "Tech 1 Vehicle Factory", w: 4, h: 3, appearance: app,
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500}}

	default:
		colors := []int{
			-1, -1,
			-1, -1}
		app := &buildingAppearance{chars: "" +
			"??" +
			"??", colors: colors}
		b = &building{name: "UNKNOWN BUILDING", w: 2, h: 2, appearance: app,
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500}}
	}
	b.x = x
	b.y = y
	b.faction = f
	return b
}
