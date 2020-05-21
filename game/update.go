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
		g.state = playingBlue
		return nil
	}

	if g.state == playingBlue || g.state == playingPink || g.state == playingWhite {
		// current player
		var currentCharacter *character
		switch g.state {
		case playingBlue:
			currentCharacter = &g.blueCharacter
		case playingWhite:
			currentCharacter = &g.whiteCharacter
		case playingPink:
			currentCharacter = &g.pinkCharacter
		}

		switch {
		case ebiten.IsGamepadButtonPressed(g.gamepadID, ebiten.GamepadButton(12)):
			// right
			currentCharacter.facing = right
			switch g.getFieldMove(currentCharacter.x, currentCharacter.y, currentCharacter.speed) {
			case normalFieldMove:
				currentCharacter.x += currentCharacter.speed
				currentCharacter.state = move
			case blueFieldMove:
				if g.state == playingBlue {
					g.state = blueSpecialMovePhase1
					currentCharacter.state = specialMove
				} else {
					currentCharacter.state = idle
				}
			case noFieldMove:
				currentCharacter.state = idle
			}
		case ebiten.IsGamepadButtonPressed(g.gamepadID, ebiten.GamepadButton(14)):
			// left
			currentCharacter.facing = left
			switch g.getFieldMove(currentCharacter.x, currentCharacter.y, -currentCharacter.speed) {
			case normalFieldMove:
				currentCharacter.x -= currentCharacter.speed
				currentCharacter.state = move
			case blueFieldMove:
				if g.state == playingBlue {
					g.state = blueSpecialMovePhase1
					currentCharacter.state = specialMove
				} else {
					currentCharacter.state = idle
				}
			case noFieldMove:
				currentCharacter.state = idle
			}
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(5)):
			// switch right
			if currentCharacter.state == idle {
				switch g.state {
				case playingBlue:
					g.state = playingPink
				case playingPink:
					g.state = playingWhite
				case playingWhite:
					g.state = playingBlue
				}
			}
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(4)):
			// switch left
			if currentCharacter.state == idle {
				switch g.state {
				case playingBlue:
					g.state = playingWhite
				case playingPink:
					g.state = playingBlue
				case playingWhite:
					g.state = playingPink
				}
			}
		default:
			currentCharacter.state = idle
		}

	}

	if g.state == blueSpecialMovePhase1 {
		if g.blueCharacter.animationStep >= 6 {
			if g.blueCharacter.facing == right {
				g.blueCharacter.x += 1.75
			} else {
				g.blueCharacter.x -= 1.75
			}
			g.state = blueSpecialMovePhase2
		}
	}

	if g.state == blueSpecialMovePhase2 {
		if g.blueCharacter.animationStep < 6 {
			g.state = playingBlue
		}
	}

	// update animations
	g.updateAnimation()

	// update camera
	g.setCameraPosition()

	return nil
}
