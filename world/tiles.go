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
	house1Type          = FieldType{backgroundWall, 6, 6}
	house2Type          = FieldType{backgroundWall, 7, 6}
	house3Type          = FieldType{backgroundWall, 8, 6}
	house4Type          = FieldType{backgroundWall, 6, 5}
	house5Type          = FieldType{backgroundWall, 7, 5}
	house6Type          = FieldType{backgroundWall, 8, 5}
	house7Type          = FieldType{backgroundWall, 9, 6}
	house8Type          = FieldType{backgroundWall, 9, 5}
	tree11Type          = FieldType{backgroundWall, 2, 8}
	tree21Type          = FieldType{backgroundWall, 3, 8}
	tree31Type          = FieldType{backgroundWall, 4, 8}
	tree41Type          = FieldType{backgroundWall, 5, 8}
	tree12Type          = FieldType{backgroundWall, 2, 7}
	tree22Type          = FieldType{backgroundWall, 3, 7}
	tree32Type          = FieldType{backgroundWall, 4, 7}
	tree42Type          = FieldType{backgroundWall, 5, 7}
	tree13Type          = FieldType{backgroundWall, 6, 8}
	tree23Type          = FieldType{backgroundWall, 7, 8}
	tree33Type          = FieldType{backgroundWall, 8, 8}
	tree43Type          = FieldType{backgroundWall, 9, 8}
	tree14Type          = FieldType{backgroundWall, 6, 7}
	tree24Type          = FieldType{backgroundWall, 7, 7}
	tree34Type          = FieldType{backgroundWall, 8, 7}
	tree44Type          = FieldType{backgroundWall, 9, 7}
	flower1type         = FieldType{nothing, 8, 1}
	flower2type         = FieldType{nothing, 9, 1}
	smallTree11type     = FieldType{nothing, 0, 8}
	smallTree12type     = FieldType{nothing, 1, 8}
	smallTree13type     = FieldType{nothing, 0, 7}
	smallTree14type     = FieldType{nothing, 1, 7}
	smallTree21type     = FieldType{nothing, 10, 6}
	smallTree22type     = FieldType{nothing, 11, 6}
	smallTree23type     = FieldType{nothing, 10, 5}
	smallTree24type     = FieldType{nothing, 11, 5}
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
	house1Tile                    = FieldTile{house1Type, nothingType, nothingType, nothingType}
	house2Tile                    = FieldTile{house2Type, nothingType, nothingType, nothingType}
	house3Tile                    = FieldTile{house3Type, nothingType, nothingType, nothingType}
	house4Tile                    = FieldTile{house4Type, nothingType, nothingType, nothingType}
	house5Tile                    = FieldTile{house5Type, nothingType, nothingType, nothingType}
	house6Tile                    = FieldTile{house6Type, nothingType, nothingType, nothingType}
	house7Tile                    = FieldTile{house7Type, nothingType, nothingType, nothingType}
	house8Tile                    = FieldTile{tree22Type, house8Type, nothingType, nothingType}
	tree11Tile                    = FieldTile{tree11Type, nothingType, nothingType, nothingType}
	tree21Tile                    = FieldTile{tree21Type, nothingType, nothingType, nothingType}
	tree31Tile                    = FieldTile{tree31Type, nothingType, nothingType, nothingType}
	tree41Tile                    = FieldTile{tree41Type, nothingType, nothingType, nothingType}
	tree12Tile                    = FieldTile{tree12Type, nothingType, nothingType, nothingType}
	tree22Tile                    = FieldTile{tree22Type, nothingType, nothingType, nothingType}
	tree32Tile                    = FieldTile{tree32Type, nothingType, nothingType, nothingType}
	tree42Tile                    = FieldTile{tree42Type, nothingType, nothingType, nothingType}
	tree13Tile                    = FieldTile{tree13Type, nothingType, nothingType, nothingType}
	tree23Tile                    = FieldTile{tree23Type, nothingType, nothingType, nothingType}
	tree33Tile                    = FieldTile{tree33Type, nothingType, nothingType, nothingType}
	tree43Tile                    = FieldTile{tree43Type, nothingType, nothingType, nothingType}
	tree14Tile                    = FieldTile{tree14Type, nothingType, nothingType, nothingType}
	tree24Tile                    = FieldTile{tree24Type, nothingType, nothingType, nothingType}
	tree34Tile                    = FieldTile{tree34Type, nothingType, nothingType, nothingType}
	tree44Tile                    = FieldTile{tree44Type, nothingType, nothingType, nothingType}
)
