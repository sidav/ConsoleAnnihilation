package main

type building struct {
	w, h                                           int
	appearance                                     *buildingAppearance
	hasBeenPlaced                                  bool
	allowsTightPlacement                           bool
	canBeBuiltOnMetalOnly, canBeBuiltOnThermalOnly bool
}
