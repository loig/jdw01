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
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type gamepadbuttons struct {
	left  int
	right int
	up    int
	down  int
	x     int
	lb    int
	rb    int
}

func (g *Game) getGamepadPlugged() bool {

	gpIDs := ebiten.GamepadIDs()
	if len(gpIDs) <= 0 {
		return false
	}
	g.gamepadID = gpIDs[0]
	g.buttons = gamepadbuttons{-1, -1, -1, -1, -1, -1, -1}
	return true
}

func (g *Game) getButtonPressed() int {
	maxButton := ebiten.GamepadButton(ebiten.GamepadButtonNum(g.gamepadID))
	for b := ebiten.GamepadButton(g.gamepadID); b < maxButton; b++ {
		if inpututil.IsGamepadButtonJustPressed(g.gamepadID, b) {
			return int(b)
		}
	}
	return -1
}

func (g *Game) setButton(button *int, wanted int) bool {
	if g.buttons.left == wanted ||
		g.buttons.right == wanted ||
		g.buttons.down == wanted ||
		g.buttons.up == wanted ||
		g.buttons.x == wanted ||
		g.buttons.lb == wanted ||
		g.buttons.rb == wanted {
		return false
	}
	*button = wanted
	return true
}
