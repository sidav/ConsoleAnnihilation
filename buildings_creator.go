package main

func createBuilding(codename string, x, y int, f *faction) *pawn {
	var b *pawn
	bStatic := getBuildingStaticInfo(codename)
	b = &pawn{
		buildingInfo: &building{code: codename},
		x: x,
		y: y,
		currentConstructionStatus: bStatic.defaultConstructionInfo.clone(),
		hitpoints: bStatic.maxHitpoints,
	}
	if b.sightRadius == 0 {
		b.sightRadius = bStatic.w + 2
	}
	if b.nanolatherInfo != nil && b.res == nil {
		b.res = &pawnResourceInformation{} // adds zero-value resource info struct for spendings usage.
	}
	return b
}

func getBuildingNameAndDescription(code string) (string, string) {
	bld := createBuilding(code, 0, 0, nil)
	name := bld.getName()
	var description string
	if bld.currentConstructionStatus != nil {
		constr := bld.currentConstructionStatus
		description += constr.getDescriptionString() + " \\n "
	}
	if len(bld.weapons) > 0 {
		for _, wpn := range bld.weapons {
			description += wpn.getDescriptionString() + " \\n "
		}
	}
	switch code {
	case "metalmaker":
		description += "A very complicated device which uses quantum fluctuations to convert large energy amounts to some metal."
	case "mextractor":
		description += "A basic ore extraction and purification device. Should be placed on metal deposits."
	case "geo":
		description += "Classic heat-to-electricity conversion device. Should be placed on thermal vents."
	case "mstorage":
		description += "Allows to store more metal."
	case "estorage":
		description += "Allows to store more energy."
	case "quark": // cheating building, useful for debugging
		description += "You should not see that text."
	case "armhq": // cheating building for stub AI
		description += "Enemy HQ. Can gather metal, generate energy and produce Tech 1 land forces."
	case "armvehfactory", "corevehfactory":
		description += "A basic nanolathing facility which is designed to construct wheeled or tracked vehicles. "
	case "armkbotlab", "corekbotlab":
		description += "A basic nanolathing facility which is designed to construct the Kinetic Bio-Organic Technology " +
			"mechs, or KBots. "
	case "coret2kbotlab":
		description += "A more advanced nanolathing facility which is designed to construct Tech 2 KBots."
	case "solar":
		description += "A classic solar battery array. The heavy use of superconductors and wireless energy transfer technologies " +
			"made this energy acqurement devices much more efficient than ever."
	case "lturret":
		description += "A basic yet quite universal base defense structure. Its only weapon uses EM-waves amplified by stimulated emission of radiation."
	case "guardian":
		description += "A stationary plasma artillery with great range and damage, but slow rate of fire."
	case "railgunturret":
		description += "A stationary defense structure which fires projectiles accelerated with Lorenz' force to hypersound velocities. " +
			"Has great range and damage, but slow rate of fire."
	case "radar":
		description += "A radar facility. Reveals enemy units' locations in an area arount itself. Drains energy."
	case "wall":
		description += "A hard metal block designed to block enemy movement."
	default:
		description += "No description."
	}
	return name, description
}
