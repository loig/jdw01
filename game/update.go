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

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/loig/jdw01/util"
)

// Update implements one of the required methods
// for the ebiten.Game interface
func (g *Game) Update(screen *ebiten.Image) error {

	switch g.state {
	case initGame:
		// Check for gamepad
		gpIDs := ebiten.GamepadIDs()
		if len(gpIDs) <= 0 {
			return errNoGamePad
		}

		g.gamepadID = gpIDs[0]
		g.state = playingBlue
		return nil

	case playingBlue, playingPink, playingWhite:
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
					g.state = blueSpecialMove
					currentCharacter.state = specialMove
					currentCharacter.specialMoveCurrentFrame = 0
				} else {
					currentCharacter.state = idle
				}
			case pinkDownFieldMove:
				if g.state == playingPink {
					g.state = pinkSpecialMoveDown
					currentCharacter.state = specialMove
					currentCharacter.specialMoveCurrentFrame = 0
				}
			case pinkUpFieldMove:
				if g.state == playingPink {
					g.state = pinkSpecialMoveUp
					currentCharacter.state = specialMove
					currentCharacter.specialMoveCurrentFrame = 0
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
					g.state = blueSpecialMove
					currentCharacter.state = specialMove
					currentCharacter.specialMoveCurrentFrame = 0
				} else {
					currentCharacter.state = idle
				}
			case pinkDownFieldMove:
				if g.state == playingPink {
					g.state = pinkSpecialMoveDown
					currentCharacter.state = specialMove
					currentCharacter.specialMoveCurrentFrame = 0
				}
			case pinkUpFieldMove:
				if g.state == playingPink {
					g.state = pinkSpecialMoveUp
					currentCharacter.state = specialMove
					currentCharacter.specialMoveCurrentFrame = 0
				}
			case noFieldMove:
				currentCharacter.state = idle
			}
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(11)):
			// up
			if g.state == playingWhite {
				if g.fieldOkForWhiteSpecialMove(currentCharacter.x, currentCharacter.y, -1) {
					g.state = whiteSpecialMove
					currentCharacter.state = specialMove
					currentCharacter.x = math.Round(currentCharacter.x)
				}
			}
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(13)):
			//down
			if g.state == playingWhite {
				if g.fieldOkForWhiteSpecialMove(currentCharacter.x, currentCharacter.y, 1) {
					g.state = whiteSpecialMove
					currentCharacter.state = specialMove
					currentCharacter.x = math.Round(currentCharacter.x)
					currentCharacter.y += 0.5
				}
			}
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(5)):
			// switch right
			if currentCharacter.state == idle {
				switch g.state {
				case playingBlue:
					g.state = playingPink
				case playingPink:
					g.state = playingWhite
					if g.whiteCharacter.state == specialMove {
						g.state = whiteSpecialMoveIdle
					}
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
					if g.whiteCharacter.state == specialMove {
						g.state = whiteSpecialMoveIdle
					}
				case playingPink:
					g.state = playingBlue
				case playingWhite:
					g.state = playingPink
				}
			}
		default:
			currentCharacter.state = idle
		}

	case blueSpecialMove:
		g.blueCharacter.specialMoveCurrentFrame++
		if g.blueCharacter.specialMoveCurrentFrame == g.blueCharacter.specialMoveNumFrames {
			if g.blueCharacter.facing == right {
				g.blueCharacter.x += 1.75
			} else {
				g.blueCharacter.x -= 1.75
			}
		}
		if g.blueCharacter.specialMoveCurrentFrame > g.blueCharacter.specialMoveNumFrames {
			if g.blueCharacter.animationStep == 0 {
				g.state = playingBlue
				g.blueCharacter.state = idle
			}
		}

	case pinkSpecialMoveDown:
		g.pinkCharacter.specialMoveCurrentFrame++
		// horizontal move
		t := float64(g.pinkCharacter.specialMoveCurrentFrame) / float64(g.pinkCharacter.specialMoveNumFrames)
		tprev := float64(g.pinkCharacter.specialMoveCurrentFrame-1) / float64(g.pinkCharacter.specialMoveNumFrames)
		frameHorizontalMove := util.SmoothStop6(t) - util.SmoothStop6(tprev)
		if g.pinkCharacter.facing == right {
			g.pinkCharacter.x += 0.75 * frameHorizontalMove
		} else {
			g.pinkCharacter.x -= 0.75 * frameHorizontalMove
		}
		// vertical move
		firstSplit := g.pinkCharacter.specialMoveNumFrames / 3
		secondSplit := g.pinkCharacter.specialMoveNumFrames - firstSplit
		if g.pinkCharacter.specialMoveCurrentFrame <= firstSplit {
			t := float64(g.pinkCharacter.specialMoveCurrentFrame) / float64(firstSplit)
			tprev := float64(g.pinkCharacter.specialMoveCurrentFrame-1) / float64(firstSplit)
			frameVerticalMove := util.SmoothStop6(t) - util.SmoothStop6(tprev)
			g.pinkCharacter.y -= 0.25 * frameVerticalMove
		} else {
			t := float64(g.pinkCharacter.specialMoveCurrentFrame-firstSplit) / float64(secondSplit)
			tprev := float64(g.pinkCharacter.specialMoveCurrentFrame-1-firstSplit) / float64(secondSplit)
			frameVerticalMove := util.SmoothStart3(t) - util.SmoothStart3(tprev)
			g.pinkCharacter.y += 1.25 * frameVerticalMove
		}
		// end of move
		if g.pinkCharacter.specialMoveCurrentFrame >= g.pinkCharacter.specialMoveNumFrames {
			g.state = playingPink
			g.pinkCharacter.state = idle
			g.pinkCharacter.y = math.Round(g.pinkCharacter.y)
		}

	case pinkSpecialMoveUp:
		g.pinkCharacter.specialMoveCurrentFrame++
		if g.pinkCharacter.facing == right {
			g.pinkCharacter.x += 0.75 / float64(g.pinkCharacter.specialMoveNumFrames)
		} else {
			g.pinkCharacter.x -= 0.75 / float64(g.pinkCharacter.specialMoveNumFrames)
		}
		// vertical move
		firstSplit := g.pinkCharacter.specialMoveNumFrames / 3
		secondSplit := g.pinkCharacter.specialMoveNumFrames - firstSplit
		if g.pinkCharacter.specialMoveCurrentFrame <= firstSplit {
			t := float64(g.pinkCharacter.specialMoveCurrentFrame) / float64(firstSplit)
			tprev := float64(g.pinkCharacter.specialMoveCurrentFrame-1) / float64(firstSplit)
			frameVerticalMove := util.SmoothStop3(t) - util.SmoothStop3(tprev)
			g.pinkCharacter.y -= 1.25 * frameVerticalMove
		} else {
			t := float64(g.pinkCharacter.specialMoveCurrentFrame-firstSplit) / float64(secondSplit)
			tprev := float64(g.pinkCharacter.specialMoveCurrentFrame-1-firstSplit) / float64(secondSplit)
			frameVerticalMove := util.SmoothStart3(t) - util.SmoothStart3(tprev)
			g.pinkCharacter.y += 0.25 * frameVerticalMove
		}
		if g.pinkCharacter.specialMoveCurrentFrame >= g.pinkCharacter.specialMoveNumFrames {
			g.state = playingPink
			g.pinkCharacter.state = idle
			g.pinkCharacter.y = math.Round(g.pinkCharacter.y)
		}

	case whiteSpecialMove, whiteSpecialMoveIdle:
		switch {
		case ebiten.IsGamepadButtonPressed(g.gamepadID, ebiten.GamepadButton(11)):
			// up
			switch g.getLaderFieldMove(g.whiteCharacter.x, g.whiteCharacter.y, -g.whiteCharacter.speed/2) {
			case normalFieldMove:
				g.whiteCharacter.y -= g.whiteCharacter.speed / 2
				g.state = whiteSpecialMove
			case endOfLaderFieldMove:
				g.whiteCharacter.y -= g.whiteCharacter.speed / 2
				g.whiteCharacter.y = math.Round(g.whiteCharacter.y)
				g.state = playingWhite
			case noFieldMove:
				g.state = whiteSpecialMoveIdle
			}
		case ebiten.IsGamepadButtonPressed(g.gamepadID, ebiten.GamepadButton(13)):
			// down
			switch g.getLaderFieldMove(g.whiteCharacter.x, g.whiteCharacter.y, +g.whiteCharacter.speed/2) {
			case normalFieldMove:
				g.whiteCharacter.y += g.whiteCharacter.speed / 2
				g.state = whiteSpecialMove
			case endOfLaderFieldMove:
				g.whiteCharacter.y += g.whiteCharacter.speed / 2
				g.whiteCharacter.y = math.Round(g.whiteCharacter.y)
				g.state = playingWhite
			case noFieldMove:
				g.state = whiteSpecialMoveIdle
			}
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(5)):
			// switch right
			if g.state == whiteSpecialMoveIdle {
				g.state = playingBlue
			}
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(4)):
			// switch left
			if g.state == whiteSpecialMoveIdle {
				g.state = playingPink
			}
		default:
			g.state = whiteSpecialMoveIdle
		}
	}

	// update animations
	g.updateAnimation()

	// update camera
	g.setCameraPosition()

	return nil
}
