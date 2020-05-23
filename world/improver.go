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

// make ceilings higher
func improveUnderworld(field [][]FieldTile, tmpFloorLevel, width, height int) {
	for y := height - 1; y > tmpFloorLevel+1; y-- {
		for x := 0; x < width; x++ {
			if field[y][x] == backgroundWallTile {
				if x-1 >= 0 &&
					(field[y-2][x-1] == backgroundWallTile ||
						field[y-2][x-1] == floorLevelBackgroundWallTile) {
					continue
				}
				if x+1 < width &&
					(field[y-2][x+1] == backgroundWallTile ||
						field[y-2][x+1] == floorLevelBackgroundWallTile) {
					continue
				}
				if field[y-2][x] == backgroundWallTile ||
					field[y-2][x] == floorLevelBackgroundWallTile {
					continue
				}
				if rand.Intn(chanceToDig) != 0 {
					field[y-1][x] = backgroundWallTile
				}
			}
		}
	}
}

// make ceilings higher
func improveFlyworld(field [][]FieldTile, tmpFloorLevel, width, height int) {
	for y := tmpFloorLevel - 2; y >= 0; y-- {
		for x := 0; x < width; x++ {
			if IsFloorField(field[y][x]) {
				fmt.Println("floorFound:", x, y)
				minX := x
				maxX := x
				for maxX < width && IsFloorField(field[y][maxX]) {
					maxX++
				}
				maxX--
				maxY := tmpFloorLevel
				for currentX := minX; currentX <= maxX; currentX++ {
					currentY := y + 1
					for currentY < tmpFloorLevel-1 && !IsFloorField(field[currentY][currentX]) {
						currentY++
					}
					currentY -= 3
					if currentY < maxY {
						maxY = currentY
					}
				}
				middle := minX + (maxX-minX+1)/2
				fmt.Println(maxX, maxY, middle)
				for currentY := y + 1; currentY <= maxY; currentY++ {
					proba := chanceToGrow - currentY + y
					if proba <= 0 {
						proba = 1
					}
					if rand.Intn(proba) != 0 {
						addIslandTile(field, middle, currentY)
					} else {
						break
					}
				}
				for rightX := middle + 1; rightX <= maxX; rightX++ {
					for currentY := y + 1; currentY <= maxY; currentY++ {
						if IsFloorField(field[currentY][rightX-1]) {
							proba := chanceToGrow - currentY + y - rightX + middle
							if proba <= 0 {
								proba = 1
							}
							if rand.Intn(proba) != 0 {
								addIslandTile(field, rightX, currentY)
							} else {
								break
							}
						}
					}
				}
				for leftX := middle - 1; leftX >= minX; leftX-- {
					for currentY := y + 1; currentY <= maxY; currentY++ {
						if IsFloorField(field[currentY][leftX+1]) {
							proba := chanceToGrow - currentY + y - middle + leftX
							if proba <= 0 {
								proba = 1
							}
							if rand.Intn(proba) != 0 {
								addIslandTile(field, leftX, currentY)
							} else {
								break
							}
						}
					}
				}
				x = maxX + 1
			}
		}
	}
}

func addIslandTile(field [][]FieldTile, x, y int) {
	if IsLaderField(field[y][x]) {
		field[y][x] = islandLaderTile
	} else {
		field[y][x] = floorTile
	}
}
