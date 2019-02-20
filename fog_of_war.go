package main

func (f *faction) recalculateSeenTiles() {
	for i := 0; i < mapW; i++ {
		for j := 0; j < mapH; j++ {
			f.tilesInSight[i][j] = false
		}
	}
	for _, p := range CURRENT_MAP.pawns {
		if p.faction == f {
			px, py := p.getCenter()
			for x := px - p.sightRadius; x <= px+p.sightRadius; x++ {
				for y := py - p.sightRadius; y <= py+p.sightRadius; y++ {
					if areCoordsValid(x, y) {
						if areCoordsInRange(px, py, x, y, p.sightRadius) {
							f.seenTiles[x][y] = true
							f.tilesInSight[x][y] = true
						}
					}
				}
			}
		}
	}
}
