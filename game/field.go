// Package game implements ebiten.Game
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
package game

import (
	"math"
)

type fieldTile struct {
	tile       fieldType
	decoration fieldType
	destructed fieldType
}

type fieldType struct {
	kind  fieldKind
	lookX int
	lookY int
}

type fieldKind int

const (
	floor fieldKind = iota
	traversableWall
	wall
	backgroundWall
	lader
	nothing
	floorLader
)

func isBackgroundField(field fieldTile) bool {
	switch field.tile.kind {
	case nothing, backgroundWall, lader:
		return true
	}
	return false
}

func isLaderField(field fieldTile) bool {
	return field.tile.kind == lader || field.decoration.kind == lader
}

func isFloorField(field fieldTile) bool {
	switch field.tile.kind {
	case floor, floorLader, wall, traversableWall:
		return true
	}
	return false
}

func isTraversableField(field fieldTile) bool {
	switch field.tile.kind {
	case traversableWall:
		return true
	}
	return false
}

func isBreakableField(field fieldTile) bool {
	switch field.tile.kind {
	case wall:
		return true
	}
	return false
}

type fieldMove int

const (
	noFieldMove fieldMove = iota
	normalFieldMove
	blueFieldMove
	pinkUpFieldMove
	pinkDownFieldMove
	endOfLaderFieldMove
)

func (g *Game) getFieldMove(xinit, yinit, offset float64) fieldMove {
	xreach := xinit + offset
	if offset >= 0 {
		xreach += 0.25
	} else {
		xreach -= 0.25
	}
	intx := int(math.Round(xreach))
	inty := int(math.Round(yinit))
	if len(g.field) < inty+2 || inty < -1 {
		return noFieldMove
	}
	if (offset >= 0 && len(g.field[inty+1]) < intx+1) || (offset < 0 && intx < 0) {
		return noFieldMove
	}
	if !(isFloorField(g.field[inty+1][intx])) {
		if isBackgroundField(g.field[inty+1][intx]) &&
			isBackgroundField(g.field[inty][intx]) &&
			(len(g.field) >= inty+3) &&
			isFloorField(g.field[inty+2][intx]) {
			return pinkDownFieldMove
		}
		return noFieldMove
	}
	if !isBackgroundField(g.field[inty][intx]) {
		if (offset >= 0 &&
			len(g.field[inty+1]) >= intx+2 &&
			isFloorField(g.field[inty+1][intx+1]) &&
			isBackgroundField(g.field[inty][intx+1])) ||
			(offset < 0 &&
				intx >= 1 &&
				isFloorField(g.field[inty+1][intx-1]) &&
				isBackgroundField(g.field[inty][intx-1])) {
			return blueFieldMove
		}
		if inty >= 1 && intx >= 1 &&
			isBackgroundField(g.field[inty-1][intx-1]) &&
			isBackgroundField(g.field[inty-1][intx]) &&
			isFloorField(g.field[inty][intx]) {
			return pinkUpFieldMove
		}
		return noFieldMove
	}
	return normalFieldMove
}

func (g *Game) fieldOkForWhiteSpecialMove(x, y float64, direction int) bool {
	intx := int(math.Round(x))
	inty := int(math.Round(y))
	switch {
	case direction < 0:
		//go up
		return isLaderField(g.field[inty][intx])
	case direction > 0:
		//go down
		return len(g.field) > inty+1 && isLaderField(g.field[inty+1][intx])
	}
	return false
}

func (g *Game) getLaderFieldMove(xinit, yinit, offset float64) fieldMove {
	yreach := yinit + offset
	if offset >= 0 {
		yreach += 0.5
	} else {
		yreach -= 0.5
	}
	intx := int(math.Round(xinit))
	inty := int(math.Round(yreach))
	if offset >= 0 {
		if inty >= len(g.field) || intx < 0 || intx >= len(g.field[inty]) {
			return noFieldMove
		}
		if !isBackgroundField(g.field[inty][intx]) &&
			!isLaderField(g.field[inty][intx]) &&
			!isFloorField(g.field[inty][intx]) {
			return noFieldMove
		}
		if isFloorField(g.field[inty][intx]) &&
			!isLaderField(g.field[inty][intx]) {
			return endOfLaderFieldMove
		}
	} else {
		if inty+1 >= 0 && inty+2 < len(g.field) && intx < len(g.field[inty+1]) && intx < len(g.field[inty+2]) && isFloorField(g.field[inty+2][intx]) &&
			!isLaderField(g.field[inty+1][intx]) {
			return endOfLaderFieldMove
		}
		if inty < 0 || intx < 0 || intx >= len(g.field[inty]) {
			return noFieldMove
		}
		if !isBackgroundField(g.field[inty][intx]) &&
			!isLaderField(g.field[inty][intx]) {
			return noFieldMove
		}
	}

	return normalFieldMove
}

func (g *Game) strikeEffectOnField(xinit, yinit float64, direction int) {
	xreach := xinit
	if direction > 0 {
		xreach += 0.75
	} else {
		xreach -= 0.75
	}
	intx := int(math.Round(xreach))
	inty := int(math.Round(yinit))
	if inty >= 0 && inty < len(g.field) && intx >= 0 && intx < len(g.field[inty]) {
		if isBreakableField(g.field[inty][intx]) {
			g.field[inty][intx].tile = g.field[inty][intx].destructed
		}
	}
}

var (
	nothingType         = fieldType{nothing, 10, 1}
	wallType            = fieldType{wall, 4, 2}
	traversableWallType = fieldType{traversableWall, 7, 3}
	backgroundWallType  = fieldType{backgroundWall, 4, 0}
	floorType           = fieldType{floor, 3, 0}
	laderType           = fieldType{lader, 8, 4}
)

var (
	nothingTile         = fieldTile{nothingType, nothingType, nothingType}
	wallTile            = fieldTile{wallType, nothingType, nothingType}
	underWorldWallTile  = fieldTile{wallType, nothingType, backgroundWallType}
	traversableWallTile = fieldTile{traversableWallType, nothingType, nothingType}
	backgroundWallTile  = fieldTile{backgroundWallType, nothingType, nothingType}
	floorTile           = fieldTile{floorType, nothingType, nothingType}
	laderTile           = fieldTile{laderType, nothingType, nothingType}
	floorladerTile      = fieldTile{floorType, laderType, nothingType}
)

func (g *Game) setInitialField() {
	field := [][]fieldTile{
		[]fieldTile{nothingTile, nothingTile, nothingTile, wallTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{floorTile, floorTile, floorladerTile, floorTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{nothingTile, nothingTile, floorladerTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{nothingTile, floorTile, floorTile, floorTile, floorTile, floorladerTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{nothingTile, traversableWallTile, nothingTile, nothingTile, nothingTile, laderTile, wallTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{backgroundWallTile, backgroundWallTile, backgroundWallTile, backgroundWallTile, backgroundWallTile, backgroundWallTile, backgroundWallTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{backgroundWallTile, backgroundWallTile, floorTile, floorTile, underWorldWallTile, backgroundWallTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile},
		[]fieldTile{backgroundWallTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
	}
	g.field = field
}
