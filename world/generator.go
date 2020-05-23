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
)

const (
	minEntryPoints           = 10
	maxEntryPoints           = 15
	minIslands               = 10
	maxIslands               = 20
	minSizeIsland            = 8
	maxSizeIsland            = 25
	minBlueWalls             = 10
	maxBlueWalls             = 20
	minSpaceBetweenBlueWalls = 15
	maxSpaceBetweenBlueWalls = 30
)

// GenerateField generates a field and returns it
func GenerateField(width, height int) (field [][]FieldTile, floorLevel float64) {
	// seeding the random generator
	rand.Seed(time.Now().UTC().UnixNano())

	tmpFloorLevel := height * 3 / 4
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
	/*
		field[tmpFloorLevel-1][25] = traversableWallTile
		field[tmpFloorLevel-2][25] = traversableWallTile
		field[tmpFloorLevel-1][29] = traversableWallTile
		field[tmpFloorLevel-1][30] = destroyableWallTile
		field[tmpFloorLevel-1][1] = traversableWallTile
		field[tmpFloorLevel-1][0] = destroyableWallTile
	*/
	pinkStart := coordinates{0, tmpFloorLevel - 1}
	whiteStart := coordinates{0, tmpFloorLevel - 1}
	goal := coordinates{width - 1, tmpFloorLevel - 1}

	generateUnderworld(field, tmpFloorLevel, width, height)
	generateSkyworld(field, tmpFloorLevel, width, height)
	generateBlueWalls(field, tmpFloorLevel, width, height, pinkStart, whiteStart, goal)

	paths := reachableByPink(field, pinkStart)
	displayPaths(field, paths)

	// If things are added outside of the playing field, it must be
	// done after this point (i.e. things added to the left/right of the field)

	return field, float64(tmpFloorLevel - 1)
}
