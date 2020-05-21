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

import "math"

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
	scale
	nothing
)

func (g *Game) setInitialField() {
	field := [][]fieldTile{
		[]fieldTile{fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}},
		[]fieldTile{fieldTile{nothing, 10, 1}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{nothing, 10, 1}},
		[]fieldTile{fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}},
		[]fieldTile{fieldTile{wall, 7, 3}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}},
		[]fieldTile{fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}},
		[]fieldTile{fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}},
		[]fieldTile{fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}, fieldTile{nothing, 10, 1}},
		[]fieldTile{fieldTile{floor, 3, 0}, fieldTile{nothing, 10, 1}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}, fieldTile{floor, 3, 0}},
	}
	g.field = field
}

func (g *Game) isValidField(xinit, yinit, xoffset float64) bool {
	xreach := xinit + xoffset
	if xoffset >= 0 {
		xreach += 0.25
	} else {
		xreach -= 0.25
	}
	intx := int(math.Round(xreach))
	inty := int(math.Round(yinit))
	if len(g.field) < inty+2 || inty < -1 {
		return false
	}
	if len(g.field[inty+1]) < intx+1 || intx < 0 {
		return false
	}
	if !(g.field[inty+1][intx].kind == floor) {
		return false
	}
	return g.field[inty][intx].kind == nothing
}
