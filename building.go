package main

type building struct {
	w, h                  int
	appearance            *buildingAppearance
	hasBeenPlaced         bool
	canBeBuiltOnMetalOnly, canBeBuiltOnThermalOnly bool
}
