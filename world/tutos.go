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

// Positions of the players in the second tuto
var (
	BlueXTuto2  float64 = 9
	WhiteXTuto2 float64 = 3
	PinkXTuto2  float64 = 15
	BlueYTuto2  float64 = 9
	WhiteYTuto2 float64 = 9
	PinkYTuto2  float64 = 10
)

// Tuto2Field is the field for the first tuto
var Tuto2Field [][]FieldTile = [][]FieldTile{
	[]FieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, leftIslandTile, islandLaderTile, islandTile, rightIslandTile, nothingTile, nothingTile, leftIslandTile, islandLaderTile, islandTile, rightIslandTile, nothingTile, nothingTile, leftIslandTile, islandLaderTile, islandTile, rightIslandTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, floorLevelfloorTile, floorLevelBackgroundWallTile, floorLevelBackgroundWallTile, floorLevelfloorTile, nothingTile, nothingTile, floorLevelfloorTile, floorLevelBackgroundWallTile, floorLevelBackgroundWallTile, floorLevelfloorTile, nothingTile, nothingTile, floorLevelfloorTile, floorLevelBackgroundWallTile, floorLevelBackgroundWallTile, floorLevelfloorTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, floorTile, floorTile, backgroundWallTile, floorTile, nothingTile, nothingTile, floorTile, floorTile, backgroundWallTile, floorTile, nothingTile, nothingTile, floorTile, floorTile, backgroundWallTile, floorTile, nothingTile, nothingTile},
}

// Positions of the players in the third tuto
var (
	BlueXTuto3  float64 = 9
	WhiteXTuto3 float64 = 3
	PinkXTuto3  float64 = 15
	BlueYTuto3  float64 = 9
	WhiteYTuto3 float64 = 9
	PinkYTuto3  float64 = 10
)

// Tuto3Field is the field for the first tuto
var Tuto3Field [][]FieldTile = [][]FieldTile{
	[]FieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, leftIslandTile, islandLaderTile, islandTile, rightIslandTile, nothingTile, nothingTile, leftIslandTile, islandLaderTile, islandTile, rightIslandTile, nothingTile, nothingTile, leftIslandTile, islandLaderTile, islandTile, rightIslandTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, floorLevelfloorTile, floorLevelBackgroundWallTile, floorLevelBackgroundWallTile, floorLevelfloorTile, nothingTile, nothingTile, floorLevelfloorTile, floorLevelBackgroundWallTile, floorLevelBackgroundWallTile, floorLevelfloorTile, nothingTile, nothingTile, floorLevelfloorTile, floorLevelBackgroundWallTile, floorLevelBackgroundWallTile, floorLevelfloorTile, nothingTile, nothingTile},
	[]FieldTile{nothingTile, nothingTile, floorTile, floorTile, backgroundWallTile, floorTile, nothingTile, nothingTile, floorTile, floorTile, backgroundWallTile, floorTile, nothingTile, nothingTile, floorTile, floorTile, backgroundWallTile, floorTile, nothingTile, nothingTile},
}
