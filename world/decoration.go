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

var (
	hole        []FieldTile = []FieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile}
	houseBottom []FieldTile = []FieldTile{tree11Tile, tree21Tile, tree31Tile, house1Tile, house2Tile, house3Tile, house7Tile, tree31Tile, tree41Tile, nothingTile}
	houseTop    []FieldTile = []FieldTile{tree12Tile, tree22Tile, tree32Tile, house4Tile, house5Tile, house6Tile, house8Tile, tree32Tile, tree42Tile, nothingTile}
	treeMid     []FieldTile = []FieldTile{tree13Tile, tree23Tile, tree33Tile, tree43Tile, nothingTile, tree13Tile, tree23Tile, tree33Tile, tree43Tile, nothingTile}
	treeMid2    []FieldTile = []FieldTile{tree13Tile, tree23Tile, tree33Tile, tree43Tile, nothingTile, tree14Tile, tree24Tile, tree34Tile, tree44Tile, nothingTile}
	treeMid3    []FieldTile = []FieldTile{tree13Tile, tree23Tile, tree33Tile, tree43Tile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile}
	treeTop     []FieldTile = []FieldTile{tree14Tile, tree24Tile, tree34Tile, tree44Tile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile}
	floorLine   []FieldTile = []FieldTile{floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile}
	underground []FieldTile = []FieldTile{floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile}
)

func addHouse(field [][]FieldTile, floorLevel int) [][]FieldTile {
	for y := 0; y < len(field); y++ {
		switch {
		case y == floorLevel-1:
			field[y] = append(field[y], houseBottom...)
		case y == floorLevel-2:
			field[y] = append(field[y], houseTop...)
		case y == floorLevel-3:
			field[y] = append(field[y], treeMid...)
		case y == floorLevel-4:
			field[y] = append(field[y], treeMid2...)
		case y == floorLevel-5:
			field[y] = append(field[y], treeMid3...)
		case y == floorLevel-6:
			field[y] = append(field[y], treeTop...)
		case y == floorLevel:
			field[y] = append(field[y], floorLine...)
		case y < floorLevel-6:
			field[y] = append(field[y], hole...)
		case y > floorLevel:
			field[y] = append(field[y], underground...)
		}
	}
	return field
}
