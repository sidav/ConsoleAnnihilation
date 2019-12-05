package main

type buildingStaticData struct {
	name                                           string
	w, h                                           int
	allowsTightPlacement                           bool
	defaultIncomeData                              *pawnIncomeInformation
	canBeBuiltOnMetalOnly, canBeBuiltOnThermalOnly bool
	appearance                                     *buildingAppearance
	defaultConstructionInfo                        *constructionInformation
	maxHitpoints                                   int
	sightRadius, radarRadius                       int
}
