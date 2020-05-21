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
	kind      fieldKind
	tileLookX int
	tileLookY int
}

type fieldKind int

const (
	floor fieldKind = iota
	traversableWall
	wall
	backgroundWall
	scale
	nothing
)

func isBackgroundField(field fieldKind) bool {
	return field == nothing || field == backgroundWall
}

type fieldMove int

const (
	noFieldMove fieldMove = iota
	normalFieldMove
	blueFieldMove
	pinkUpFieldMove
	pinkDownFieldMove
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
	if !(g.field[inty+1][intx].kind == floor) {
		if isBackgroundField(g.field[inty+1][intx].kind) &&
			isBackgroundField(g.field[inty][intx].kind) &&
			(len(g.field) >= inty+3) &&
			g.field[inty+2][intx].kind == floor {
			return pinkDownFieldMove
		}
		return noFieldMove
	}
	if !isBackgroundField(g.field[inty][intx].kind) {
		if (offset >= 0 && len(g.field[inty+1]) < intx+2) || (offset < 0 && intx < 1) {
			return noFieldMove
		}
		if (offset >= 0 && !(g.field[inty+1][intx+1].kind == floor)) ||
			(offset < 0 && !(g.field[inty+1][intx-1].kind == floor)) {
			return noFieldMove
		}
		if (offset >= 0 && isBackgroundField(g.field[inty][intx+1].kind)) ||
			(offset < 0 && isBackgroundField(g.field[inty][intx-1].kind)) {
			return blueFieldMove
		}
		return noFieldMove
	}
	return normalFieldMove
}

func (g *Game) setInitialField() {
	field := [][]fieldTile{
		[]fieldTile{fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}},
		[]fieldTile{fieldTile{nothing, 10, 1}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{nothing, 10, 1}},
		[]fieldTile{fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}},
		[]fieldTile{fieldTile{nothing, 10, 1}, fieldTile{traversableWall, 7, 3}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{traversableWall, 7, 3}, fieldTile{wall, 4, 2}},
		[]fieldTile{fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}},
		[]fieldTile{fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}},
		[]fieldTile{fieldTile{backgroundWall, 4, 0}, fieldTile{backgroundWall, 4, 0}, fieldTile{backgroundWall, 4, 0}, fieldTile{backgroundWall, 4, 0}, fieldTile{backgroundWall, 4, 0}, fieldTile{backgroundWall, 4, 0}, fieldTile{backgroundWall, 4, 0}},
		[]fieldTile{fieldTile{backgroundWall, 4, 0}, fieldTile{backgroundWall, 4, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{backgroundWall, 4, 0}, fieldTile{floor, 3, 0}},
		[]fieldTile{fieldTile{backgroundWall, 4, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}},
		[]fieldTile{fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}},
	}
	g.field = field
}
