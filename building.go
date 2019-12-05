package main

type building struct {
	code                                           string
	hasBeenPlaced                                  bool
}

func (b *building) getBuildingStaticInfo() *buildingStaticData {
	return getBuildingStaticInfo(b.code)
}

func (b *building) getAppearance() *buildingAppearance {
	return b.getBuildingStaticInfo().appearance
}
