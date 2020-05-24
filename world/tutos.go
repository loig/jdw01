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

// Positions of the players in the first tuto
var (
	BlueXTuto1  float64 = 1
	WhiteXTuto1 float64 = 1
	PinkXTuto1  float64 = 1
	BlueYTuto1  float64 = 6
	WhiteYTuto1 float64 = 2
	PinkYTuto1  float64 = 10
)

// Tuto1Field is the field for the first tuto
var Tuto1Field [][]FieldTile = [][]FieldTile{
	[]FieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, traversableWallTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, nothingTile, traversableWallTile, nothingTile, nothingTile, nothingTile, nothingTile, traversableWallTile, destroyableWallTile, nothingTile, nothingTile, nothingTile, nothingTile, traversableWallTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, traversableWallTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, nothingTile, traversableWallTile, nothingTile, nothingTile, nothingTile, nothingTile, traversableWallTile, destroyableWallTile, nothingTile, nothingTile, nothingTile, nothingTile, traversableWallTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, traversableWallTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, nothingTile, traversableWallTile, nothingTile, nothingTile, nothingTile, nothingTile, traversableWallTile, destroyableWallTile, nothingTile, nothingTile, nothingTile, nothingTile, traversableWallTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile, floorLevelfloorTile},
}
