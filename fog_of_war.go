package main

import geometry "github.com/sidav/golibrl/geometry"

func (f *faction) recalculateSeenTiles() {
	if CHEAT_IGNORE_FOW {
		return
	}
	for i := 0; i < mapW; i++ {
		for j := 0; j < mapH; j++ {
			f.tilesInSight[i][j] = false
			f.radarCoverage[i][j] = false
		}
	}
	for _, p := range CURRENT_MAP.pawns {
		if p.faction == f {
			if p.currentConstructionStatus != nil {
				continue
			}
			px, py := p.getCenter()
			radiusToIterate := p.sightRadius
			if p.radarRadius > p.sightRadius {
				radiusToIterate = p.radarRadius
			}
			for x := px - radiusToIterate; x <= px+radiusToIterate; x++ {
				for y := py - radiusToIterate; y <= py+radiusToIterate; y++ {
					if areCoordsValid(x, y) {
						if geometry.AreCoordsInRange(px, py, x, y, p.sightRadius) {
							f.seenTiles[x][y] = true
							f.tilesInSight[x][y] = true
						}
						if p.radarRadius > 0 && geometry.AreCoordsInRange(px, py, x, y, p.radarRadius) {
							f.radarCoverage[x][y] = true
						}
					}
				}
			}
		}
	}
}

func (f *faction) areCoordsInSight(x, y int) bool {
	return f.tilesInSight[x][y] || CHEAT_IGNORE_FOW || f.aiControlled
}

func (f *faction) wereCoordsSeen(x, y int) bool {
	return f.seenTiles[x][y] || CHEAT_IGNORE_FOW || f.aiControlled
}

func (f *faction) areCoordsInRadarRadius(x, y int) bool {
	return f.radarCoverage[x][y] || f.aiControlled
}
