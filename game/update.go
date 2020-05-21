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
)

// Update implements one of the required methods
// for the ebiten.Game interface
func (g *Game) Update(screen *ebiten.Image) error {

	if g.state == initGame {
		// Check for gamepad
		gpIDs := ebiten.GamepadIDs()
		if len(gpIDs) <= 0 {
			return errNoGamePad
		}

		g.gamepadID = gpIDs[0]
		g.state = playing
		return nil
	}

	// update animations
	g.updateAnimation()

	if ebiten.IsGamepadButtonPressed(g.gamepadID, ebiten.GamepadButton(12)) {
		// right
		g.blueCharacter.facing = right
		g.blueCharacter.state = move
		return nil
	}

	if ebiten.IsGamepadButtonPressed(g.gamepadID, ebiten.GamepadButton(14)) {
		// right
		g.blueCharacter.facing = left
		g.blueCharacter.state = move
		return nil
	}

	g.blueCharacter.state = idle

	return nil
}
