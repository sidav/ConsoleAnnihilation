package main

import (
	"SomeTBSGame/routines"
	"strconv"
)

func ai_makeBuildOrderForBuilding(builder *pawn, buildingCode string) bool { // decides building placement, NOT a metal extractor / geothermal powerplant

	BUILD_SEARCH_RANGE := 10

	success := false

	bx, by := builder.getCenter()
	building := createBuilding(buildingCode, bx, by, builder.faction)
	b_w, b_h := building.buildingInfo.w, building.buildingInfo.h

	var placex, placey int
	for try := 0; try < 10; try++ {
		placex = routines.RandInRange(bx-BUILD_SEARCH_RANGE, bx+BUILD_SEARCH_RANGE)
		placey = routines.RandInRange(by-BUILD_SEARCH_RANGE, by+BUILD_SEARCH_RANGE)
		if CURRENT_MAP.canBuildingBeBuiltAt(building, placex, placey) &&
			CURRENT_MAP.getNumberOfMetalDepositsInRect(placex-b_w/2, placey-b_h/2, b_w, b_h) == 0 { // restrict building on metal
			building.x, building.y = placex-b_w/2, placey-b_h/2
			builder.setOrder(&order{orderType: order_build, buildingToConstruct: building})
			success = true
			break
		}
	}
	return success
	//for x:=bx-BUILD_SEARCH_RANGE; x < bx + BUILD_SEARCH_RANGE; x++ {
	//	for y:=by-BUILD_SEARCH_RANGE; y<by+BUILD_SEARCH_RANGE;y++ {
	//
	//	}
	//}
}

func ai_tryBuildMetalExtractor(builder *pawn, buildingCode string) bool {

	BUILD_SEARCH_RANGE := 25

	success := false

	bx, by := builder.getCenter()
	building := createBuilding(buildingCode, bx, by, builder.faction)
	b_w, b_h := building.buildingInfo.w, building.buildingInfo.h

	var placex, placey int
	goodplacex, goodplacey := -1, -1 // center, not upper-left
	metalForGoodPlace := 0
	for placex = bx - BUILD_SEARCH_RANGE; placex < bx+BUILD_SEARCH_RANGE; placex++ {
		for placey = by - BUILD_SEARCH_RANGE; placey < by+BUILD_SEARCH_RANGE; placey++ {
			metalHere := CURRENT_MAP.getNumberOfMetalDepositsInRect(placex-b_w/2, placey-b_h/2, b_w, b_h)
			// ai_write(fmt.Sprintf("%d METAL AT %d, %d (%dx%d)", metalHere, placex, placey, b_w, b_h))
			if metalHere > metalForGoodPlace &&
				CURRENT_MAP.canBuildingBeBuiltAt(building, placex, placey) {
				goodplacex, goodplacey = placex, placey
				metalForGoodPlace = metalHere
			}
		}
	}
	if metalForGoodPlace > 0 {
		ai_write("building metal extractor for " + strconv.Itoa(metalForGoodPlace) + " deposits.")
		building.x, building.y = goodplacex-b_w/2, goodplacey-b_h/2
		builder.setOrder(&order{orderType: order_build, buildingToConstruct: building})
		success = true
	} else {
		ai_write("A good place for metal extractor is not found.")
	}
	return success
}
