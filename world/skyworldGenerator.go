// Package world implements world representation and generation
/*
Copyright (C) 2020  Loïg Jezequel

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package world

import (
	"fmt"
	"math/rand"
)

type island struct {
	xstart   int
	xend     int
	altitude int
}

func generateSkyworld(field [][]FieldTile, floorLevel, width, height int) {
	// 1. Generate the number of islands
	numIsland := rand.Intn(maxIslands+1-minIslands) + minIslands
	fmt.Println("Generating", numIsland, "islands")
	islands := make([]island, numIsland)

	// 2. Generate the x range of each Island
	lastMaxX := 0
	for i := 0; i < numIsland; i++ {
		// start of current island
		var minX int
		if lastMaxX == 0 {
			// first island always starts at the same point
			minX = 1
		} else if lastMaxX > width-minSizeIsland {
			// if last generated island was almost at the end of the field
			// restart from start of the field
			startRange := width / (numIsland - i)
			minX = rand.Intn(startRange) + 1
			lastMaxX = minX
		} else {
			// else the next island should start somewhere before the last
			// island end
			minX = lastMaxX - 4 - rand.Intn(minSizeIsland)
		}
		// end of current island
		var maxX int
		maxX = minX + minSizeIsland + rand.Intn(maxSizeIsland-minSizeIsland)
		if lastMaxX+minSizeIsland > maxX {
			maxX = lastMaxX + minSizeIsland
		}
		if maxX > width-2 {
			maxX = width - 2
		}
		// update lastMaxX
		lastMaxX = maxX
		// build island
		islands[i].xstart = minX
		islands[i].xend = maxX
		islands[i].altitude = floorLevel - 2*i - 2
	}
	fmt.Println("Islands:", islands)

	// 3. Generate the altitude of each island

	// Draw the islands
	for i := 0; i < numIsland; i++ {
		for x := islands[i].xstart; x <= islands[i].xend; x++ {
			field[islands[i].altitude][x] = floorLevelfloorTile
		}
	}
}
