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
	nothingType         = FieldType{nothing, 10, 1}
	destroyableWallType = FieldType{destroyableWall, 4, 2}
	traversableWallType = FieldType{traversableWall, 7, 3}
	backgroundWallType  = FieldType{backgroundWall, 4, 0}
	floorType           = FieldType{floor, 3, 0}
	laderType           = FieldType{lader, 8, 4}
	grassType           = FieldType{floor, 1, 1}
)

var (
	nothingTile                   = FieldTile{nothingType, nothingType, nothingType}
	detstroyableWallTile          = FieldTile{destroyableWallType, nothingType, nothingType}
	underworldDestroyableWallTile = FieldTile{destroyableWallType, nothingType, backgroundWallType}
	traversableWallTile           = FieldTile{traversableWallType, nothingType, nothingType}
	backgroundWallTile            = FieldTile{backgroundWallType, nothingType, nothingType}
	floorLevelBackgroundWallTile  = FieldTile{backgroundWallType, grassType, nothingType}
	floorTile                     = FieldTile{floorType, nothingType, nothingType}
	floorLevelfloorTile           = FieldTile{floorType, grassType, nothingType}
	laderTile                     = FieldTile{laderType, nothingType, nothingType}
	floorladerTile                = FieldTile{floorType, laderType, nothingType}
)
