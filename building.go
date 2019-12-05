package main

type building struct {
	code           string
	hasBeenPlaced  bool
	nanolatherInfo *nanolatherInformation
}

func (b *building) getBuildingStaticInfo() *buildingStaticData {
	return getBuildingStaticInfo(b.code)
}

func (b *building) getName() string {
	return b.getBuildingStaticInfo().name
}

func (b *building) getAppearance() *buildingAppearance {
	return b.getBuildingStaticInfo().appearance
}
