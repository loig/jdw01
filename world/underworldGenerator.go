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

	"github.com/loig/jdw01/util"
)

func generateUnderworld(field [][]FieldTile, floorLevel, width, height int) {
	// 1. generate entry points to the underworld
	numEntryPoints := rand.Intn(maxEntryPoints+1-minEntryPoints) + minEntryPoints
	tilesPerEntry := width / numEntryPoints
	entryPointsRegister := make([][]int, numEntryPoints)
	// position of entry points
	for i := 0; i < numEntryPoints; i++ {
		posEntryPoint := tilesPerEntry*i + 1 + rand.Intn(tilesPerEntry-2)
		entryPointsRegister[i] = []int{posEntryPoint, posEntryPoint + 1}
	}
	// merge entry points that overlap, sort the register
	// this is probably no longuer useful
	tmpRegister := make([][]int, 0)
	for len(entryPointsRegister) != 0 {
		nextEntryPoint := entryPointsRegister[0]
		entryPointsRegister = entryPointsRegister[1:len(entryPointsRegister)]
		overlaping := make([]int, 0)
		for i, entryPoint := range entryPointsRegister {
			if util.Overlap(nextEntryPoint, entryPoint) {
				overlaping = append(overlaping, i)
				nextEntryPoint = util.Merge(nextEntryPoint, entryPoint)
			}
		}
		for i := len(overlaping) - 1; i >= 0; i-- {
			entryPointsRegister = append(
				entryPointsRegister[:overlaping[i]],
				entryPointsRegister[overlaping[i]+1:]...,
			)
		}
		tmpRegister = append(tmpRegister, nextEntryPoint)
	}
	entryPointsRegister = util.RegisterSort(tmpRegister)
	// draw entry points
	for _, entryPoint := range entryPointsRegister {
		for _, posx := range entryPoint {
			field[floorLevel][posx] = floorLevelBackgroundWallTile
		}
	}

	// 2. Associate entry points
	maxPossibleAssociatedIndice := make([]int, len(entryPointsRegister))
	associatedIndice := make([]int, len(entryPointsRegister))
	hasAssociatedIndice := make([]bool, len(entryPointsRegister))
	for i := 0; i < len(entryPointsRegister); i++ {
		maxPossibleAssociatedIndice[i] = len(entryPointsRegister)
	}
	for i := 0; i < len(entryPointsRegister); i++ {
		if !hasAssociatedIndice[i] {
			if maxPossibleAssociatedIndice[i] == i+1 {
				associatedIndice[i] = i
				hasAssociatedIndice[i] = true
			} else {
				j := rand.Intn(maxPossibleAssociatedIndice[i]-i-1) + i + 1
				associatedIndice[i] = j
				associatedIndice[j] = i
				hasAssociatedIndice[i] = true
				hasAssociatedIndice[j] = true
				for k := i + 1; k < j; k++ {
					maxPossibleAssociatedIndice[k] = j
				}
			}
		}
	}

	// 3. Build a cave between each pair of entry points
	caveBuilt := make([]bool, len(entryPointsRegister))
	maxDepth := make([]int, width)
	for i := 0; i < width; i++ {
		maxDepth[i] = height - 2
	}
	for i := 0; i < len(entryPointsRegister); i++ {
		//for i := 0; i < 1; i++ {
		if !caveBuilt[i] {
			j := associatedIndice[i]
			caveBuilt[i] = true
			caveBuilt[j] = true
			numInBetween := ((j - i) - (j-i)%2) / 2
			minDepth := numInBetween*2 + floorLevel + 1
			firstStart := -1
			lastEnd := -1
			if i+1 < j {
				firstStart = entryPointsRegister[i+1][0]
				lastEnd = entryPointsRegister[j-1][len(entryPointsRegister[j-1])-1]
			}
			start := entryPointsRegister[i][0]
			end := entryPointsRegister[j][len(entryPointsRegister[j])-1]
			buildCave(field, maxDepth, floorLevel, width, height, minDepth, start, end, firstStart, lastEnd)
		}
	}

}

func buildCave(field [][]FieldTile, maxDepth []int, floorLevel, width, height, minDepth, start, end, firstStart, lastEnd int) {
	// record for the new max depths
	updatedMaxDepth := make([]int, width)
	for i := 0; i < width; i++ {
		updatedMaxDepth[i] = maxDepth[i]
	}
	// generation variables
	currentDepth := floorLevel + 1
	currentMinDepth := floorLevel + 1
	currentMaxDepth := 0
	// normal case, two entry points
	for currentx := start + 1; currentx+1 < end; currentx++ {
		// place current underground
		field[currentDepth][currentx] = backgroundWallTile
		field[currentDepth][currentx+1] = backgroundWallTile

		// update currentMinDepth
		if firstStart >= 0 {
			mustGoDown := currentx - firstStart + floorLevel + 5
			canGoUp := lastEnd - currentx + floorLevel + 2
			if canGoUp > mustGoDown {
				canGoUp = mustGoDown
			}
			if canGoUp > minDepth {
				canGoUp = minDepth
			}
			if canGoUp < floorLevel+1 {
				canGoUp = floorLevel + 1
			}
			currentMinDepth = canGoUp
		}

		// update currentMaxDepth
		currentMaxDepth = end - currentx + floorLevel - 2
		if currentMaxDepth > maxDepth[currentx] {
			currentMaxDepth = maxDepth[currentx]
		}
		if currentMaxDepth > maxDepth[currentx+1] {
			currentMaxDepth = maxDepth[currentx+1]
		}
		if currentx+2 < width && currentMaxDepth > maxDepth[currentx+2] {
			currentMaxDepth = maxDepth[currentx+2]
		}
		if currentx+3 < width && currentMaxDepth > maxDepth[currentx+3] {
			currentMaxDepth = maxDepth[currentx+3]
		}

		// update maxDepth
		if updatedMaxDepth[currentx] > currentDepth-2 {
			updatedMaxDepth[currentx] = currentDepth - 2
		}
		if updatedMaxDepth[currentx+1] > currentDepth-2 {
			updatedMaxDepth[currentx+1] = currentDepth - 2
		}

		// if no choice go up
		if currentDepth > currentMaxDepth {
			currentDepth--
			continue
		}

		// if no choice, go down
		if currentDepth < currentMinDepth {
			currentDepth++
			continue
		}

		// else choose
		switch rand.Intn(6) {
		case 0, 1:
			// goUp
			if currentDepth-1 >= currentMinDepth {
				currentDepth--
			}
		case 2:
			// don't change depth
			continue
		case 3, 4, 5:
			// goDown
			if currentDepth+1 <= currentMaxDepth {
				currentDepth++
			}
		}

	}

	// Case with only one entry point
	if start+1 == end {
		fmt.Println("plop")
		// go left or right
		direction := rand.Intn(2)
		if direction == 0 {
			direction = -1
		}
		// set position and depth
		currentx := end
		if direction < 0 {
			currentx = start
		}
		// currentMinDepth and currentDepth are already set to floorLevel + 1
		// set currentMaxDepth
		currentMaxDepth = maxDepth[currentx]
		if currentx+direction >= 0 &&
			currentx+direction < width &&
			currentMaxDepth > maxDepth[currentx+direction] {
			currentMaxDepth = maxDepth[currentx+direction]
		}
		// generate cave
		for currentMaxDepth > floorLevel && currentx > 0 && currentx+1 < width && currentx+direction > 0 && currentx+direction+1 < width {
			// place current underground
			field[currentDepth][currentx] = backgroundWallTile
			field[currentDepth][currentx+direction] = backgroundWallTile

			// update currentMaxDepth
			currentMaxDepth = maxDepth[currentx+direction]
			if currentx+2*direction >= 0 &&
				currentx+2*direction < width &&
				currentMaxDepth > maxDepth[currentx+2*direction] {
				currentMaxDepth = maxDepth[currentx+2*direction]
			}
			if currentx+3*direction >= 0 &&
				currentx+3*direction < width &&
				currentMaxDepth > maxDepth[currentx+3*direction] {
				currentMaxDepth = maxDepth[currentx+3*direction]
			}
			if currentx-direction >= 0 &&
				currentx-direction < width &&
				currentMaxDepth > maxDepth[currentx-direction] {
				currentMaxDepth = maxDepth[currentx-direction]
			}
			fmt.Println(currentMaxDepth, direction, currentx)

			// update maxDepth
			if updatedMaxDepth[currentx] > currentDepth-2 {
				updatedMaxDepth[currentx] = currentDepth - 2
			}
			if updatedMaxDepth[currentx+direction] > currentDepth-2 {
				updatedMaxDepth[currentx+direction] = currentDepth - 2
			}

			// update currentx
			currentx += direction

			// if no choice go up
			if currentDepth > currentMaxDepth {
				currentDepth--
				continue
			}

			// else choose
			switch rand.Intn(9) {
			case 0, 1, 2:
				// goUp
				if currentDepth-1 >= currentMinDepth {
					currentDepth--
				}
			case 3, 4, 5, 6, 7, 8:
				// goDown
				if currentDepth+1 <= currentMaxDepth {
					currentDepth++
				}
			}

		}

		// if the cave was not possible to build, erase it
		if currentx == start || currentx == end {
			field[floorLevel][start] = floorLevelBackgroundWallTile
			field[floorLevel][end] = floorLevelBackgroundWallTile
		}
	}

	// update maxDepth at start and end
	updatedMaxDepth[start] = floorLevel - 2
	updatedMaxDepth[end] = floorLevel - 2

	//update maxDepth
	for i := 0; i < width; i++ {
		maxDepth[i] = updatedMaxDepth[i]
	}
}
