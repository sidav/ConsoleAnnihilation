package main

import "fmt"

func createBuilding(name string, x, y int, f *faction) *pawn {
	var b *pawn
	switch name {

	case "metalmaker":
		colors := []int{
			7, 7,
			7, -1}
		app := &buildingAppearance{chars: "" +
			"xx" +
			"xx", colors: colors}
		b = &pawn{name: "Metal Synthesizer",
			buildingInfo:              &building{w: 2, h: 2, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 35, costM: 10, costE: 500},
			res:                       &pawnResourceInformation{metalIncome: 1, energyReqForConditionalMetalIncome: 60},
		}

	case "solar":
		colors := []int{
			-1, 7,
			7, -1}
		app := &buildingAppearance{chars: "" +
			"==" +
			"==", colors: colors}
		b = &pawn{name: "Solar Collector",
			buildingInfo:              &building{w: 2, h: 2, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 60, costM: 100, costE: 500},
			res:                       &pawnResourceInformation{energyIncome: 20},
		}

	case "quark": // cheating building, useful for debugging
		colors := []int{
			7, -1, 7,
			-1, 13, -1,
			7, -1, 7}
		app := &buildingAppearance{chars: "" +
			"=^=" +
			"<O>" +
			"=V=", colors: colors}
		b = &pawn{name: "Quark Antimatter Generator",
			buildingInfo:              &building{w: 3, h: 3, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 40, costM: 100, costE: 1200},
			res:                       &pawnResourceInformation{metalIncome: 5, energyIncome: 100},
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
		b = &pawn{name: "Tech 1 KBot Lab",
			buildingInfo:              &building{w: 3, h: 3, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500},
		}

	case "corekbotlab":
		colors := []int{
			7, 7, 7,
			7, -1, 7,
			7, 7, 7}
		app := &buildingAppearance{chars: "" +
			"\\=/" +
			"=0=" +
			"/=\\", colors: colors}
		b = &pawn{name: "Tech 1 KBot Lab",
			buildingInfo:              &building{w: 3, h: 3, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500},
			nanolatherInfo:            &nanolatherInformation{builderCoeff: 1, allowedUnits: []string{"coreck", "ak", "thud"}},
		}

	case "armvehfactory":
		colors := []int{
			7, 7, 7, 7,
			7, -1, -1, 7,
			7, 7, 7, 7}
		app := &buildingAppearance{chars: "" +
			"====" +
			"|--|" +
			"\\==/", colors: colors}
		b = &pawn{name: "Tech 1 Vehicle Factory",
			buildingInfo:              &building{w: 4, h: 3, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500},
		}

	case "corevehfactory":
		colors := []int{
			7, 7, 7, 7,
			7, -1, -1, 7,
			7, -1, -1, 7}
		app := &buildingAppearance{chars: "" +
			"=--=" +
			"|/\\|" +
			"\\\\//", colors: colors}
		b = &pawn{name: "Tech 1 Vehicle Factory", buildingInfo: &building{w: 4, h: 3, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500},
		}

	case "lturret":
		colors := []int{-1}
		app := &buildingAppearance{chars: "T", colors: colors}
		b = &pawn{name: "Light Laser Turret", maxHitpoints: 90, isHeavy: true,
			buildingInfo: &building{w: 1, h: 1, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 250, costE: 900},
			weapons: []*pawnWeaponInformation{
				{attackDelay: 12, attackEnergyCost: 15, attackRadius: 6, attacksLand: true,
					hitscan: &WeaponHitscan{baseDamage:6},
				},
			},
		}

	default:
		colors := []int{
			-1, -1,
			-1, -1}
		app := &buildingAppearance{chars: "" +
			"??" +
			"??", colors: colors}
		b = &pawn{name: "UNKNOWN BUILDING",
			buildingInfo:              &building{w: 2, h: 2, appearance: app},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 100, costM: 100, costE: 500},
		}
	}
	if b.maxHitpoints == 0 {
		b.maxHitpoints = 1
	}
	b.hitpoints = b.maxHitpoints
	b.x = x
	b.y = y
	b.faction = f
	if b.nanolatherInfo != nil && b.res == nil {
		b.res = &pawnResourceInformation{} // adds zero-value resource info struct for spendings usage.
	}

	return b
}

func getBuildingNameAndDescription(code string) (string, string) {
	bld := createBuilding(code, 0, 0, nil)
	name := bld.name
	constr := bld.currentConstructionStatus
	description := fmt.Sprintf("Metal: %d ENERGY: %d Base build time: %d \\n ", constr.costM, constr.costE, constr.maxConstructionAmount)
	switch code {
	case "armkbotlab", "corekbotlab":
		description += "A basic nanolathing facility which is designed to construct the Kinetic Bio-Organic Technology " +
			"mechs, or KBots. "
	case "solar":
		description += "A classic solar battery array. The heavy use of superconductors and wireless energy transfer technologies " +
			"made this energy acqurement devices much more efficient than ever."
	default:
		description += "No description."
	}
	return name, description
}
