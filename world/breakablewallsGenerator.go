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

import "math/rand"

func generateBreakableWalls(field [][]FieldTile, tmpFloorLevel, width, height int, blueStart, pinkStart, whiteStart, goal coordinates) {

	// Go through the field from the end and check where
	// walls would block at least one character, but no
	// more than two
	for x := width - 1; x >= 0; x-- {
		if x == blueStart.x || x == pinkStart.x || x == whiteStart.x {
			continue
		}
		if stopsPink(field, x, tmpFloorLevel-1) {
			doPutWall := rand.Intn(chanceOfBreakWall) != 0
			for x >= 0 && IsBackgroundField(field[tmpFloorLevel][x]) {
				if doPutWall {
					putWall(field, x, tmpFloorLevel-1)
				}
				x--
			}
		} else if stopsWhite(field, x, tmpFloorLevel-1) {
			if rand.Intn(chanceOfBreakWall) != 0 {
				putWall(field, x, tmpFloorLevel-1)
			}
		} else if x < width-1 && stopsBlue(field, x, tmpFloorLevel-1) {
			if breakableByOther(field, x, tmpFloorLevel-1, pinkStart, whiteStart) {
				putWall(field, x, tmpFloorLevel-1)
			}
		}
	}

}

// checks that it blocks pink, without blocking blue
func stopsPink(field [][]FieldTile, x, y int) bool {
	ok := IsBackgroundField(field[y+1][x])
	xtmp := x
	for ok && xtmp >= 0 && IsBackgroundField(field[y+1][xtmp]) {
		xtmp--
		if xtmp >= 0 {
			ok = !IsTraversableField(field[y][xtmp])
		}
	}
	return ok
}

// checks that it blocks white, without blocking blue
func stopsWhite(field [][]FieldTile, x, y int) bool {
	return IsLaderField(field[y][x]) &&
		((x-1 >= 0 && !IsTraversableField(field[y][x-1])) || x == 0)
}

// checks that it blocks blue
func stopsBlue(field [][]FieldTile, x, y int) bool {
	return x-1 >= 0 && IsTraversableField(field[y][x-1])
}

// checks that the block can be destroyed by
// pink or by white
func breakableByOther(field [][]FieldTile, x, y int, pinkStart, whiteStart coordinates) bool {
	accessible := false
	currentTile := field[y][x]
	field[y][x] = destroyableWallTile
	paths := reachableByWhite(field, whiteStart)
	accessible = paths[y][x+1]
	if !accessible {
		paths = reachableByPink(field, pinkStart)
		accessible = paths[y][x+1]
	}
	field[y][x] = currentTile
	return accessible
}

// put a wall
func putWall(field [][]FieldTile, x, y int) {
	if IsLaderField(field[y][x]) {
		field[y][x] = onladerDestroyableWallTile
	} else {
		field[y][x] = destroyableWallTile
	}
}
