package main

import "SomeTBSGame/routines"

func (f *faction) recalculateSeenTiles() {
	for i := 0; i < mapW; i++ {
		for j := 0; j < mapH; j++ {
			f.tilesInSight[i][j] = false
			f.radarCoverage[i][j] = false
		}
	}
	for _, p := range CURRENT_MAP.pawns {
		if p.faction == f {
			px, py := p.getCenter()
			radiusToIterate := p.sightRadius
			if p.radarRadius > p.sightRadius {
				radiusToIterate = p.radarRadius
			}
			for x := px - radiusToIterate; x <= px+radiusToIterate; x++ {
				for y := py - radiusToIterate; y <= py+radiusToIterate; y++ {
					if areCoordsValid(x, y) {
						if routines.AreCoordsInRange(px, py, x, y, p.sightRadius) {
							f.seenTiles[x][y] = true
							f.tilesInSight[x][y] = true
						}
						if p.radarRadius > 0 && routines.AreCoordsInRange(px, py, x, y, p.radarRadius) {
							f.radarCoverage[x][y] = true
						}
					}
				}
			}
		}
	}
}
