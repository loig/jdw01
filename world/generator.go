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
	generateUnderworld(field, tmpFloorLevel, width)
	return field, float64(tmpFloorLevel - 1)
}

func generateUnderworld(field [][]FieldTile, floorLevel, width int) {
	// 1. generate entry points to the underworld
	numEntryPoints := rand.Intn((width-2*minTilesPerUnderworldEntryPoint)/minTilesPerUnderworldEntryPoint) + 2
	if numEntryPoints < minEntryPoints {
		numEntryPoints = minEntryPoints
	}
	entryPointsRegister := make([][]int, numEntryPoints)
	// first entry point, necessarily in the first tiles
	posEntryPoint := rand.Intn(minTilesPerUnderworldEntryPoint-2) + 1
	field[floorLevel][posEntryPoint] = floorLevelBackgroundWallTile
	field[floorLevel][posEntryPoint+1] = floorLevelBackgroundWallTile
	entryPointsRegister[numEntryPoints-1] = []int{posEntryPoint, posEntryPoint + 1}
	numEntryPoints--
	// last entry point, necessarily in the last tiles
	posEntryPoint = width - rand.Intn(minTilesPerUnderworldEntryPoint-2) - 2
	field[floorLevel][posEntryPoint] = floorLevelBackgroundWallTile
	field[floorLevel][posEntryPoint-1] = floorLevelBackgroundWallTile
	entryPointsRegister[numEntryPoints-1] = []int{posEntryPoint - 1, posEntryPoint}
	numEntryPoints--
	// other entry points
	for numEntryPoints > 0 {
		posEntryPoint = rand.Intn(width-minTilesPerUnderworldEntryPoint*2) + minTilesPerUnderworldEntryPoint
		field[floorLevel][posEntryPoint] = floorLevelBackgroundWallTile
		field[floorLevel][posEntryPoint+1] = floorLevelBackgroundWallTile
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

}
