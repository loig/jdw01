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
	"time"

	"github.com/loig/jdw01/util"
)

const (
	minTilesPerUnderworldEntryPoint = 12
	minEntryPoints                  = 5
)

// GenerateField generates a field and returns it
func GenerateField(width, height int) (field [][]FieldTile, floorLevel float64) {
	// seeding the random generator
	rand.Seed(time.Now().UTC().UnixNano())

	tmpFloorLevel := height * 2 / 3
	field = make([][]FieldTile, height)
	for y := 0; y < height; y++ {
		field[y] = make([]FieldTile, width)
		switch {
		case y == tmpFloorLevel:
			for x := 0; x < width; x++ {
				field[y][x] = floorLevelfloorTile
			}
		case y < tmpFloorLevel:
			for x := 0; x < width; x++ {
				field[y][x] = nothingTile
			}
		case y > tmpFloorLevel:
			for x := 0; x < width; x++ {
				field[y][x] = floorTile
			}
		}
	}
	generateUnderworld(field, tmpFloorLevel, width, height)
	return field, float64(tmpFloorLevel - 1)
}

func generateUnderworld(field [][]FieldTile, floorLevel, width, height int) {
	// 1. generate entry points to the underworld
	numEntryPoints := rand.Intn((width-2*minTilesPerUnderworldEntryPoint)/minTilesPerUnderworldEntryPoint) + 2
	if numEntryPoints < minEntryPoints {
		numEntryPoints = minEntryPoints
	}
	entryPointsRegister := make([][]int, numEntryPoints)
	// first entry point, necessarily in the first tiles
	posEntryPoint := rand.Intn(minTilesPerUnderworldEntryPoint-2) + 1
	entryPointsRegister[numEntryPoints-1] = []int{posEntryPoint, posEntryPoint + 1}
	numEntryPoints--
	// last entry point, necessarily in the last tiles
	posEntryPoint = width - rand.Intn(minTilesPerUnderworldEntryPoint-2) - 2
	entryPointsRegister[numEntryPoints-1] = []int{posEntryPoint - 1, posEntryPoint}
	numEntryPoints--
	// other entry points
	for numEntryPoints > 0 {
		posEntryPoint = rand.Intn(width-minTilesPerUnderworldEntryPoint*2) + minTilesPerUnderworldEntryPoint
		entryPointsRegister[numEntryPoints-1] = []int{posEntryPoint, posEntryPoint + 1}
		numEntryPoints--
	}
	fmt.Println(entryPointsRegister)
	// merge entry points that overlap, sort the register
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
	fmt.Println(entryPointsRegister)
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
			fmt.Print(i, " will be associated with… ")
			if maxPossibleAssociatedIndice[i] == i+1 {
				fmt.Println(i)
				associatedIndice[i] = i
				hasAssociatedIndice[i] = true
			} else {
				j := rand.Intn(maxPossibleAssociatedIndice[i]-i-1) + i + 1
				fmt.Println(j)
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
	fmt.Println(associatedIndice)

	// 3. Build a cave between each pair of entry points
	caveBuilt := make([]bool, len(entryPointsRegister))
	maxDepth := make([]int, len(entryPointsRegister))
	for i := 0; i < len(entryPointsRegister); i++ {
		maxDepth[i] = height
	}
	//for i := 0; i < len(entryPointsRegister); i++ {
	for i := 0; i < 1; i++ {
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
			higherPoint := buildCave(field, floorLevel, height, minDepth, maxDepth[i], start, end, firstStart, lastEnd)
			for k := i + 1; k < j; k++ {
				maxDepth[k] = higherPoint
			}
		}
	}

}

func buildCave(field [][]FieldTile, floorLevel, height, minDepth, maxDepth, start, end, firstStart, lastEnd int) (higherPoint int) {
	fmt.Println("floorLevel:", floorLevel, "minDepth:", minDepth, "maxDepth:", maxDepth, "start:", start, "end:", end, "firstStart:", firstStart, "lastEnd:", lastEnd)
	currentDepth := floorLevel + 1
	currentMinDepth := floorLevel + 1
	currentMaxDepth := maxDepth // to update
	for currentx := start + 1; currentx+1 < end; currentx++ {
		//fmt.Println(currentx, currentDepth, currentMinDepth, currentMaxDepth)
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
		if currentMaxDepth > maxDepth {
			currentMaxDepth = maxDepth
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
		switch rand.Intn(3) {
		case 0:
			// goUp
			if currentDepth-1 >= currentMinDepth {
				currentDepth--
			}
		case 1:
			// don't change depth
			continue
		case 2:
			// goDown
			if currentDepth+1 <= currentMaxDepth {
				currentDepth++
			}
		}

	}

	/*
		x := start + 1
		if firstStart != -1 {
			// Go down only to reach minDepth before firstStart + minDepth
			for x < firstStart+minDepth {
				if currentDepth == floorLevel {
					currentDepth++
					field[currentDepth][x] = backgroundWallTile
					field[currentDepth][x+1] = backgroundWallTile
				} else if x+1+(minDepth-currentDepth) < firstStart+minDepth-floorLevel-2 {
					// Do anything, if it does not prevent to reach minDepth at a good moment
					field[currentDepth][x] = backgroundWallTile
					field[currentDepth][x+1] = backgroundWallTile
				} else if x+1+(minDepth-currentDepth) == firstStart+minDepth-floorLevel-2 {
					// must go down at any cost
					currentDepth++
					field[currentDepth][x] = backgroundWallTile
					field[currentDepth][x+1] = backgroundWallTile
				} else {
					fmt.Println("oups")
				}
				x++
			}
		}
	*/
	return 0
}
