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
)

func generateBlueWalls(field [][]FieldTile, floorLevel, width, height int, blueStart, pinkStart, whiteStart, goal coordinates) {

	numBlueWalls := minBlueWalls + rand.Intn(maxBlueWalls-minBlueWalls+1)
	lastPos := 0
	for i := 0; i < numBlueWalls; i++ {
		inBetween := minSpaceBetweenBlueWalls + rand.Intn(maxSpaceBetweenBlueWalls-minSpaceBetweenBlueWalls+1)
		currentPos := lastPos + inBetween
		if currentPos >= width {
			currentPos = 0
		}
		lastPos = currentPos
		placed := false
		for !placed {
			high := floorLevel - 1
			if isValidForBlueWalls(field, currentPos, high, blueStart, pinkStart, whiteStart, goal) {
				field[high][currentPos] = traversableWallTile
				lastPos = currentPos
				placed = true
				high--
				if rand.Intn(3) == 0 {
					if isValidForBlueWalls(field, currentPos, high, blueStart, pinkStart, whiteStart, goal) {
						field[high][currentPos] = traversableWallTile
					}
				}
			}
			currentPos++
			if currentPos >= width {
				currentPos = 0
			}
			if currentPos == lastPos {
				break
			}
		}
		if placed == false { // no longer possible to place walls
			break
		}
	}

}

func isValidForBlueWalls(field [][]FieldTile, x, y int, blueStart, pinkStart, whiteStart, goal coordinates) bool {
	if !IsBackgroundField(field[y][x]) {
		return false
	}
	if x == blueStart.x && y == blueStart.y {
		return false
	}
	if x == pinkStart.x && y == pinkStart.y {
		return false
	}
	if x == whiteStart.x && y == whiteStart.y {
		return false
	}
	if y+1 < len(field) && IsBackgroundField(field[y+1][x]) {
		return false
	}
	if IsLaderField(field[y][x]) {
		return false
	}
	if x+1 < len(field[0]) && IsTraversableField(field[y][x+1]) {
		return false
	}
	if x-1 >= 0 && IsTraversableField(field[y][x-1]) {
		return false
	}
	currentTile := field[y][x]
	field[y][x] = traversableWallTile
	pinkPaths := reachableByPink(field, pinkStart)
	if !pinkPaths[goal.y][goal.x] {
		field[y][x] = currentTile
		return false
	}
	whitePaths := reachableByWhite(field, whiteStart)
	if !whitePaths[goal.y][goal.x] {
		field[y][x] = currentTile
		return false
	}
	return true
}
