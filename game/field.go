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

	"github.com/loig/jdw01/world"
)

type fieldMove int

const (
	noFieldMove fieldMove = iota
	normalFieldMove
	blueFieldMove
	pinkUpFieldMove
	pinkDownFieldMove
	pinkDownFieldMoveOrNormalMove
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
	if world.IsBackgroundField(g.field[inty+1][intx]) &&
		world.IsBackgroundField(g.field[inty][intx]) &&
		(len(g.field) >= inty+3) &&
		world.IsFloorField(g.field[inty+2][intx]) {
		if world.IsFloorField(g.field[inty+1][intx]) {
			return pinkDownFieldMoveOrNormalMove
		} else {
			return pinkDownFieldMove
		}
	}
	if !(world.IsFloorField(g.field[inty+1][intx])) {
		return noFieldMove
	}
	if !world.IsBackgroundField(g.field[inty][intx]) {
		if (offset >= 0 &&
			len(g.field[inty+1]) >= intx+2 &&
			world.IsFloorField(g.field[inty+1][intx+1]) &&
			world.IsBackgroundField(g.field[inty][intx+1])) ||
			(offset < 0 &&
				intx >= 1 &&
				world.IsFloorField(g.field[inty+1][intx-1]) &&
				world.IsBackgroundField(g.field[inty][intx-1])) {
			return blueFieldMove
		}
		if inty >= 1 && intx >= 1 &&
			world.IsBackgroundField(g.field[inty-1][intx-1]) &&
			world.IsBackgroundField(g.field[inty-1][intx]) &&
			world.IsFloorField(g.field[inty][intx]) {
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
		return world.IsLaderField(g.field[inty][intx])
	case direction > 0:
		//go down
		return len(g.field) > inty+1 && world.IsLaderField(g.field[inty+1][intx])
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
		if !world.IsBackgroundField(g.field[inty][intx]) &&
			!world.IsLaderField(g.field[inty][intx]) &&
			!world.IsFloorField(g.field[inty][intx]) {
			return noFieldMove
		}
		if world.IsFloorField(g.field[inty][intx]) &&
			!world.IsLaderField(g.field[inty][intx]) {
			return endOfLaderFieldMove
		}
	} else {
		if inty+1 >= 0 && inty+2 < len(g.field) && intx < len(g.field[inty+1]) && intx < len(g.field[inty+2]) && world.IsFloorField(g.field[inty+2][intx]) &&
			!world.IsLaderField(g.field[inty+1][intx]) {
			return endOfLaderFieldMove
		}
		if inty < 0 || intx < 0 || intx >= len(g.field[inty]) {
			return noFieldMove
		}
		if !world.IsBackgroundField(g.field[inty][intx]) &&
			!world.IsLaderField(g.field[inty][intx]) {
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
		if world.IsBreakableField(g.field[inty][intx]) {
			g.field[inty][intx].Tile = g.field[inty][intx].Destructed
		}
	}
}
