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
			if minX < 0 {
				minX = 1
			}
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
	}
	fmt.Println("Islands:", islands)

	// 3. Generate the altitude of each island
	// sort the islands by increasing xstart
	for i := numIsland - 1; i >= 1; i-- {
		for j := 0; j < i; j++ {
			if islands[j+1].xstart < islands[j].xstart {
				tmpIsland := islands[j]
				islands[j] = islands[j+1]
				islands[j+1] = tmpIsland
			}
		}
	}
	fmt.Println("Islands:", islands)
	// list islands that overlap (or could be just close to)
	// with each island and start before this island
	numOverlaping := make([]int, numIsland)
	overlaping := make([][]int, numIsland)
	for i := 0; i < numIsland; i++ {
		overlaping[i] = make([]int, 0)
	}
	for i := 0; i < numIsland; i++ {
		for j := i + 1; j < numIsland; j++ {
			if islands[i].xend+1 >= islands[j].xstart {
				//overlaping[i] = append(overlaping[i], j)
				numOverlaping[i]++
				numOverlaping[j]++
				overlaping[j] = append(overlaping[j], i)
			}
		}
	}
	fmt.Println(overlaping)
	// determine the number of possible island altitudes
	maxOverlaping := 0
	for i := 0; i < numIsland; i++ {
		if numOverlaping[i] > maxOverlaping {
			maxOverlaping = numOverlaping[i]
		}
	}
	numAltitude := (floorLevel - 4) / 5
	if maxOverlaping+1 < numAltitude {
		numAltitude = maxOverlaping + 1
	}
	// give an altitude to each island, different from the islands
	// starting before
	for i := 0; i < numIsland; i++ {
		altitudePosition := rand.Intn(numAltitude - len(overlaping[i]))
		positioned := false
		count := 0
		var altitudeNum int
		for altitudeNum = 0; !positioned; altitudeNum++ {
			taken := false
			for _, j := range overlaping[i] {
				taken = taken || islands[j].altitude == altitudeNum
			}
			if !taken {
				if count == altitudePosition {
					positioned = true
				}
				count++
			}
		}
		islands[i].altitude = altitudeNum - 1
	}
	fmt.Println(islands)
	for i := 0; i < numIsland; i++ {
		islands[i].altitude = floorLevel - (islands[i].altitude+1)*5
	}
	fmt.Println(islands)

	// 4. Draw the islands
	for i := 0; i < numIsland; i++ {
		drawIsland(field, islands[i])
	}

	// 5. Add the laders
	// sort by altitude
	for i := numIsland - 1; i >= 1; i-- {
		for j := 0; j < i; j++ {
			if islands[j+1].altitude < islands[j].altitude {
				tmpIsland := islands[j]
				islands[j] = islands[j+1]
				islands[j+1] = tmpIsland
			}
		}
	}
	fmt.Println(islands)
	// find an island just below the left side of each island
	// and an island just below the right side of each island
	justBelowLeft := make([]int, numIsland)
	justBelowRight := make([]int, numIsland)
	for i := 0; i < numIsland; i++ {
		foundLeft := false
		foundRight := false
		length := islands[i].xend - islands[i].xstart + 1
		for j := i + 1; j < numIsland && (!foundLeft || !foundRight); j++ {
			if !foundLeft &&
				islands[j].xend >= islands[i].xstart &&
				islands[j].xstart <= islands[i].xend-length/2 {
				foundLeft = true
				justBelowLeft[i] = j
			}
			if !foundRight &&
				islands[j].xstart <= islands[i].xend &&
				islands[j].xend >= islands[i].xend-length/2+1 {
				foundRight = true
				justBelowRight[i] = j
			}
		}
		if !foundLeft {
			justBelowLeft[i] = -1
		}
		if !foundRight {
			justBelowRight[i] = -1
		}
	}
	fmt.Println(justBelowLeft)
	fmt.Println(justBelowRight)
	// add laders belowLeft and belowRight of each island (if possible)
	laders := make([][]bool, numIsland)
	for i := 0; i < numIsland; i++ {
		laders[i] = make([]bool, width)
	}
	for i := numIsland - 1; i >= 0; i-- {
		doLeft := true
		doRight := true
		if justBelowLeft[i] == justBelowRight[i] {
			// only one lader when two are possible
			choice := rand.Intn(2)
			if choice == 0 {
				doLeft = false
			} else {
				doRight = false
			}
		}
		// leftLader
		if doLeft {
			j := justBelowLeft[i]
			var islandDown island
			var ladersDown []bool
			if j == -1 {
				islandDown.xstart = 0
				islandDown.xend = width - 1
				islandDown.altitude = floorLevel
				ladersDown = make([]bool, width)
			} else {
				islandDown = islands[j]
				ladersDown = laders[j]
			}
			islandUp := islands[i]
			islandWidth := islandUp.xend - islandUp.xstart + 1
			islandUp.xend = islandUp.xend - islandWidth/2
			position, okPosition := setLaderPosition(laders[i], ladersDown, islandUp, islandDown)
			if okPosition {
				drawLader(field, position, islands[i].altitude, islandDown.altitude)
			}
		}
		// rightLader
		if doRight {
			j := justBelowRight[i]
			var islandDown island
			var ladersDown []bool
			if j == -1 {
				islandDown.xstart = 0
				islandDown.xend = width - 1
				islandDown.altitude = floorLevel
				ladersDown = make([]bool, width)
			} else {
				islandDown = islands[j]
				ladersDown = laders[j]
			}
			islandUp := islands[i]
			islandWidth := islandUp.xend - islandUp.xstart + 1
			islandUp.xstart = islandUp.xstart + islandWidth/2
			position, okPosition := setLaderPosition(laders[i], ladersDown, islandUp, islandDown)
			if okPosition {
				drawLader(field, position, islands[i].altitude, islandDown.altitude)
			}
		}
	}
}

func setLaderPosition(ladersUp, ladersDown []bool, islandUp, islandDown island) (position int, okPosition bool) {
	minPosition := islandUp.xstart
	if islandDown.xstart > minPosition {
		minPosition = islandDown.xstart
	}
	maxPosition := islandUp.xend
	if islandDown.xend < maxPosition {
		maxPosition = islandDown.xend
	}
	positionRange := maxPosition - minPosition + 1
	if positionRange <= 0 { // should never occur, but not sufficiently tested
		return 0, false
	}
	position = minPosition + rand.Intn(positionRange)
	tmpPosition := position
	okPosition = !ladersUp[position] && !ladersDown[position]
	for tmpPosition < maxPosition && !okPosition {
		tmpPosition++
		okPosition = !ladersUp[tmpPosition] && !ladersDown[tmpPosition]
	}
	if !okPosition {
		tmpPosition = minPosition
	}
	for tmpPosition < position && !okPosition {
		okPosition = !ladersUp[tmpPosition] && !ladersDown[tmpPosition]
		tmpPosition++
	}
	position = tmpPosition
	if okPosition {
		ladersUp[position] = true
		ladersDown[position] = true
	}
	return position, okPosition
}

func drawIsland(field [][]FieldTile, isl island) {
	field[isl.altitude][isl.xstart] = leftIslandTile
	field[isl.altitude][isl.xend] = rightIslandTile
	for x := isl.xstart + 1; x < isl.xend; x++ {
		field[isl.altitude][x] = islandTile
	}
}

// ystart is the highest altitude
func drawLader(field [][]FieldTile, x, ystart, yend int) {
	for y := ystart; y < yend; y++ {
		if field[y][x] == islandTile {
			field[y][x] = islandLaderTile
		} else if field[y][x] == leftIslandTile {
			field[y][x] = islandLaderTile
		} else if field[y][x] == rightIslandTile {
			field[y][x] = islandLaderTile
		} else {
			field[y][x] = laderTile
		}
	}
}
