package main

import "SomeTBSGame/routines"

func ai_makeBuildOrderForBuilding(builder *pawn, buildingCode string) {
	BUILD_SEARCH_RANGE := 20
	bx, by := builder.getCenter()
	building := createBuilding(buildingCode, bx, by, builder.faction)
	b_w, b_h := building.buildingInfo.w, building.buildingInfo.h

	var placex, placey int
	for try := 0; try < 10; try++ {
		placex = routines.RandInRange(bx-BUILD_SEARCH_RANGE, bx+BUILD_SEARCH_RANGE)
		placey = routines.RandInRange(by-BUILD_SEARCH_RANGE, by+BUILD_SEARCH_RANGE)
		if CURRENT_MAP.canBuildingBeBuiltAt(building, placex, placey) {
			building.x, building.y = placex - b_w / 2, placey - b_h / 2
			builder.setOrder(&order{orderType: order_build, buildingToConstruct: building})
			break
		}
	}
	//for x:=bx-BUILD_SEARCH_RANGE; x < bx + BUILD_SEARCH_RANGE; x++ {
	//	for y:=by-BUILD_SEARCH_RANGE; y<by+BUILD_SEARCH_RANGE;y++ {
	//
	//	}
	//}
}
