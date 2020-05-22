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
	// for testing
	field[tmpFloorLevel-1][25] = traversableWallTile
	field[tmpFloorLevel-2][25] = traversableWallTile
	field[tmpFloorLevel-1][29] = traversableWallTile
	field[tmpFloorLevel-1][30] = destroyableWallTile
	field[tmpFloorLevel-1][1] = traversableWallTile
	field[tmpFloorLevel-1][0] = destroyableWallTile
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
	maxDepth := make([]int, len(entryPointsRegister))
	for i := 0; i < len(entryPointsRegister); i++ {
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
			higherPoint := buildCave(field, floorLevel, height, minDepth, maxDepth[i], start, end, firstStart, lastEnd)
			for k := i + 1; k < j; k++ {
				maxDepth[k] = higherPoint - 2
			}
		}
	}

}

func buildCave(field [][]FieldTile, floorLevel, height, minDepth, maxDepth, start, end, firstStart, lastEnd int) (higherPoint int) {
	currentDepth := floorLevel + 1
	currentMinDepth := floorLevel + 1
	currentMaxDepth := maxDepth
	higherPoint = maxDepth
	for currentx := start + 1; currentx+1 < end; currentx++ {
		// place current underground
		field[currentDepth][currentx] = backgroundWallTile
		field[currentDepth][currentx+1] = backgroundWallTile

		// update higherPoint
		if firstStart >= 0 {
			if currentx >= firstStart-1 && currentx+1 <= lastEnd {
				if currentDepth < higherPoint {
					higherPoint = currentDepth
				}
			}
		}

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

	return higherPoint
}
