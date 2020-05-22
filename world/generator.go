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

// GenerateField generates a field and returns it
func GenerateField() [][]FieldTile {
	return [][]FieldTile{
		[]FieldTile{nothingTile, nothingTile, nothingTile, detstroyableWallTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]FieldTile{floorTile, floorTile, floorladerTile, floorTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]FieldTile{nothingTile, nothingTile, floorladerTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]FieldTile{nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]FieldTile{nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]FieldTile{nothingTile, floorTile, floorTile, floorTile, floorTile, floorladerTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]FieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]FieldTile{nothingTile, traversableWallTile, nothingTile, nothingTile, nothingTile, laderTile, detstroyableWallTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]FieldTile{floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]FieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]FieldTile{backgroundWallTile, backgroundWallTile, backgroundWallTile, backgroundWallTile, backgroundWallTile, backgroundWallTile, backgroundWallTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]FieldTile{backgroundWallTile, backgroundWallTile, floorTile, floorTile, underworldDestroyableWallTile, backgroundWallTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile},
		[]FieldTile{backgroundWallTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]FieldTile{floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
	}
}
