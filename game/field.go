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
	kind            fieldKind
	tileLookX       int
	tileLookY       int
	hasDecoration   bool
	decorationLookX int
	decorationLookY int
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

func isBackgroundField(field fieldKind) bool {
	return field == nothing || field == backgroundWall || field == lader
}

func isLaderField(field fieldKind) bool {
	return field == lader || field == floorLader
}

func isFloorField(field fieldKind) bool {
	return field == floor || field == floorLader
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
	if !(isFloorField(g.field[inty+1][intx].kind)) {
		if isBackgroundField(g.field[inty+1][intx].kind) &&
			isBackgroundField(g.field[inty][intx].kind) &&
			(len(g.field) >= inty+3) &&
			isFloorField(g.field[inty+2][intx].kind) {
			return pinkDownFieldMove
		}
		return noFieldMove
	}
	if !isBackgroundField(g.field[inty][intx].kind) {
		if (offset >= 0 &&
			len(g.field[inty+1]) >= intx+2 &&
			g.field[inty+1][intx+1].kind == floor &&
			isBackgroundField(g.field[inty][intx+1].kind)) ||
			(offset < 0 &&
				intx >= 1 &&
				g.field[inty+1][intx-1].kind == floor &&
				isBackgroundField(g.field[inty][intx-1].kind)) {
			return blueFieldMove
		}
		if inty >= 1 && intx >= 1 &&
			isBackgroundField(g.field[inty-1][intx-1].kind) &&
			isBackgroundField(g.field[inty-1][intx].kind) &&
			g.field[inty][intx].kind == floor {
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
		return isLaderField(g.field[inty][intx].kind)
	case direction > 0:
		//go down
		return len(g.field) > inty+1 && isLaderField(g.field[inty+1][intx].kind)
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
	if inty < 0 || (offset < 0 &&
		!isBackgroundField(g.field[inty][intx].kind) &&
		!isLaderField(g.field[inty][intx].kind)) ||
		inty >= len(g.field) ||
		(offset >= 0 &&
			!isBackgroundField(g.field[inty][intx].kind) &&
			!isLaderField(g.field[inty][intx].kind) &&
			!isFloorField(g.field[inty][intx].kind)) {
		return noFieldMove
	}
	if (offset >= 0 &&
		isFloorField(g.field[inty][intx].kind) &&
		!isLaderField(g.field[inty][intx].kind)) ||
		(offset < 0 &&
			isFloorField(g.field[inty+2][intx].kind) &&
			!isLaderField(g.field[inty][intx].kind)) {
		return endOfLaderFieldMove
	}
	return normalFieldMove
}

var (
	nothingTile         = fieldTile{nothing, 10, 1, false, 0, 0}
	wallTile            = fieldTile{wall, 4, 2, false, 0, 0}
	traversableWallTile = fieldTile{traversableWall, 7, 3, false, 0, 0}
	backgroundWallTile  = fieldTile{backgroundWall, 4, 0, false, 0, 0}
	floorTile           = fieldTile{floor, 3, 0, false, 0, 0}
	laderTile           = fieldTile{lader, 8, 4, false, 0, 0}
	floorladerTile      = fieldTile{floorLader, 3, 0, true, 8, 4}
)

func (g *Game) setInitialField() {
	field := [][]fieldTile{
		[]fieldTile{nothingTile, nothingTile, wallTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{floorTile, floorTile, floorladerTile, floorTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{nothingTile, nothingTile, laderTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{nothingTile, floorTile, floorTile, floorTile, floorTile, floorladerTile, nothingTile},
		[]fieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, laderTile, nothingTile},
		[]fieldTile{nothingTile, traversableWallTile, nothingTile, nothingTile, nothingTile, laderTile, wallTile},
		[]fieldTile{floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile},
		[]fieldTile{nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile, nothingTile},
		[]fieldTile{backgroundWallTile, backgroundWallTile, backgroundWallTile, backgroundWallTile, backgroundWallTile, backgroundWallTile, backgroundWallTile},
		[]fieldTile{backgroundWallTile, backgroundWallTile, floorTile, floorTile, floorTile, backgroundWallTile, floorTile},
		[]fieldTile{backgroundWallTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile},
		[]fieldTile{floorTile, floorTile, floorTile, floorTile, floorTile, floorTile, floorTile},
	}
	g.field = field
}
