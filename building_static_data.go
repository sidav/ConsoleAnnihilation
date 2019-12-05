package main

type buildingStaticData struct {
	name                                           string
	w, h                                           int
	allowsTightPlacement                           bool
	defaultResourceInfo                            *pawnResourceInformation
	canBeBuiltOnMetalOnly, canBeBuiltOnThermalOnly bool
	appearance                                     *buildingAppearance
	defaultConstructionInfo                        *constructionInformation
	maxHitpoints                                   int
}
