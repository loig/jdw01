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

// How to fill outside the play field
const (
	DummyX = 3
	DummyY = 0
)

var (
	nothingType         = FieldType{nothing, 10, 1}
	destroyableWallType = FieldType{destroyableWall, 4, 2}
	traversableWallType = FieldType{traversableWall, 7, 3}
	backgroundWallType  = FieldType{backgroundWall, 4, 0}
	floorType           = FieldType{floor, 3, 0}
	laderType           = FieldType{lader, 8, 4}
	grassType           = FieldType{floor, 1, 1}
	leftGrassType       = FieldType{floor, 0, 1}
	rightGrassType      = FieldType{floor, 3, 1}
	reachableType       = FieldType{nothing, 5, 3}
)

var (
	nothingTile                   = FieldTile{nothingType, nothingType, nothingType, nothingType}
	destroyableWallTile           = FieldTile{destroyableWallType, nothingType, nothingType, nothingType}
	underworldDestroyableWallTile = FieldTile{destroyableWallType, nothingType, backgroundWallType, nothingType}
	onladerDestroyableWallTile    = FieldTile{destroyableWallType, nothingType, laderType, nothingType}
	traversableWallTile           = FieldTile{traversableWallType, nothingType, nothingType, nothingType}
	backgroundWallTile            = FieldTile{backgroundWallType, nothingType, nothingType, nothingType}
	floorLevelBackgroundWallTile  = FieldTile{backgroundWallType, grassType, nothingType, nothingType}
	floorTile                     = FieldTile{floorType, nothingType, nothingType, nothingType}
	floorLevelfloorTile           = FieldTile{floorType, grassType, nothingType, nothingType}
	laderTile                     = FieldTile{laderType, nothingType, nothingType, nothingType}
	islandTile                    = FieldTile{floorType, grassType, nothingType, nothingType}
	islandLaderTile               = FieldTile{floorType, laderType, nothingType, nothingType}
	leftIslandTile                = FieldTile{floorType, leftGrassType, nothingType, nothingType}
	rightIslandTile               = FieldTile{floorType, rightGrassType, nothingType, nothingType}
)
