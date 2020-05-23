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

type coordinates struct {
	x int
	y int
}

func reachableByPink(field [][]FieldTile, startingPoint coordinates) [][]bool {

	height := len(field)
	width := len(field[0])

	reachable := make([][]bool, height)
	for i := 0; i < height; i++ {
		reachable[i] = make([]bool, width)
	}
	nexts := make([]coordinates, 1)
	nexts[0] = startingPoint
	reachable[startingPoint.y][startingPoint.x] = true

	for len(nexts) > 0 {
		toLook := nexts[0]
		nexts = nexts[1:len(nexts)]
		// try right
		if toLook.x+1 < width && toLook.y+1 < height {
			if IsBackgroundField(field[toLook.y][toLook.x+1]) &&
				IsFloorField(field[toLook.y+1][toLook.x+1]) {
				if !reachable[toLook.y][toLook.x+1] {
					reachable[toLook.y][toLook.x+1] = true
					nexts = append(nexts, coordinates{toLook.x + 1, toLook.y})
				}
			}
		}
		// try left
		if toLook.x-1 >= 0 && toLook.y+1 < height {
			if IsBackgroundField(field[toLook.y][toLook.x-1]) &&
				IsFloorField(field[toLook.y+1][toLook.x-1]) {
				if !reachable[toLook.y][toLook.x-1] {
					reachable[toLook.y][toLook.x-1] = true
					nexts = append(nexts, coordinates{toLook.x - 1, toLook.y})
				}
			}
		}
		// try up right
		if toLook.x+1 < width && toLook.y-1 >= 0 {
			if IsBackgroundField(field[toLook.y-1][toLook.x]) &&
				IsBackgroundField(field[toLook.y-1][toLook.x+1]) &&
				IsFloorField(field[toLook.y][toLook.x+1]) {
				if !reachable[toLook.y-1][toLook.x+1] {
					reachable[toLook.y-1][toLook.x+1] = true
					nexts = append(nexts, coordinates{toLook.x + 1, toLook.y - 1})
				}
			}
		}
		// try up left
		if toLook.x-1 >= 0 && toLook.y-1 >= 0 {
			if IsBackgroundField(field[toLook.y-1][toLook.x]) &&
				IsBackgroundField(field[toLook.y-1][toLook.x-1]) &&
				IsFloorField(field[toLook.y][toLook.x-1]) {
				if !reachable[toLook.y-1][toLook.x-1] {
					reachable[toLook.y-1][toLook.x-1] = true
					nexts = append(nexts, coordinates{toLook.x - 1, toLook.y - 1})
				}
			}
		}
		// try down
		if toLook.y+2 < height {
			if IsBackgroundField(field[toLook.y+1][toLook.x]) &&
				IsFloorField(field[toLook.y+2][toLook.x]) {
				if !reachable[toLook.y+1][toLook.x] {
					reachable[toLook.y+1][toLook.x] = true
					nexts = append(nexts, coordinates{toLook.x, toLook.y + 1})
				}
			}
		}
		// try down right
		if toLook.x+1 < width && toLook.y+2 < height {
			if IsBackgroundField(field[toLook.y][toLook.x+1]) &&
				IsBackgroundField(field[toLook.y+1][toLook.x+1]) &&
				IsFloorField(field[toLook.y+2][toLook.x+1]) {
				if !reachable[toLook.y+1][toLook.x+1] {
					reachable[toLook.y+1][toLook.x+1] = true
					nexts = append(nexts, coordinates{toLook.x + 1, toLook.y + 1})
				}
			}
		}
		// try down left
		if toLook.x-1 >= 0 && toLook.y+2 < height {
			if IsBackgroundField(field[toLook.y][toLook.x-1]) &&
				IsBackgroundField(field[toLook.y+1][toLook.x-1]) &&
				IsFloorField(field[toLook.y+2][toLook.x-1]) {
				if !reachable[toLook.y+1][toLook.x-1] {
					reachable[toLook.y+1][toLook.x-1] = true
					nexts = append(nexts, coordinates{toLook.x - 1, toLook.y + 1})
				}
			}
		}
	}

	return reachable

}

func reachableByWhite(field [][]FieldTile, startingPoint coordinates) [][]bool {

	height := len(field)
	width := len(field[0])

	reachable := make([][]bool, height)
	for i := 0; i < height; i++ {
		reachable[i] = make([]bool, width)
	}
	nexts := make([]coordinates, 1)
	nexts[0] = startingPoint
	reachable[startingPoint.y][startingPoint.x] = true

	for len(nexts) > 0 {
		toLook := nexts[0]
		nexts = nexts[1:len(nexts)]
		// try right
		if toLook.x+1 < width && toLook.y+1 < height {
			if IsBackgroundField(field[toLook.y][toLook.x+1]) &&
				IsFloorField(field[toLook.y+1][toLook.x+1]) &&
				IsFloorField(field[toLook.y+1][toLook.x]) {
				if !reachable[toLook.y][toLook.x+1] {
					reachable[toLook.y][toLook.x+1] = true
					nexts = append(nexts, coordinates{toLook.x + 1, toLook.y})
				}
			}
		}
		// try left
		if toLook.x-1 >= 0 && toLook.y+1 < height {
			if IsBackgroundField(field[toLook.y][toLook.x-1]) &&
				IsFloorField(field[toLook.y+1][toLook.x-1]) &&
				IsFloorField(field[toLook.y+1][toLook.x]) {
				if !reachable[toLook.y][toLook.x-1] {
					reachable[toLook.y][toLook.x-1] = true
					nexts = append(nexts, coordinates{toLook.x - 1, toLook.y})
				}
			}
		}
		// try up
		if toLook.y-1 >= 0 {
			continueLader := IsLaderField(field[toLook.y-1][toLook.x])
			isFloorAbove := IsFloorField(field[toLook.y][toLook.x]) &&
				IsBackgroundField(field[toLook.y-1][toLook.x])
			if IsLaderField(field[toLook.y][toLook.x]) &&
				(continueLader || isFloorAbove) {
				if !reachable[toLook.y-1][toLook.x] {
					reachable[toLook.y-1][toLook.x] = true
					nexts = append(nexts, coordinates{toLook.x, toLook.y - 1})
				}
			}
		}
		// try down
		if toLook.y+1 < height {
			if IsLaderField(field[toLook.y+1][toLook.x]) {
				if !reachable[toLook.y+1][toLook.x] {
					reachable[toLook.y+1][toLook.x] = true
					nexts = append(nexts, coordinates{toLook.x, toLook.y + 1})
				}
			}
		}
	}

	return reachable

}

func displayPaths(field [][]FieldTile, paths [][]bool) {
	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths[0]); j++ {
			if paths[i][j] {
				field[i][j].Info = reachableType
			}
		}
	}
}
