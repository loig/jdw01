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
	"github.com/loig/jdw01/world"
)

// Update implements one of the required methods
// for the ebiten.Game interface
func (g *Game) Update(screen *ebiten.Image) error {

	// check for game completion
	if g.blueCharacter.x >= g.goalX &&
		g.pinkCharacter.x >= g.goalX &&
		g.whiteCharacter.x >= g.goalX {
		g.state = theEnd
	}

	switch g.state {

	case initGame:
		// Check for gamepad
		gpIDs := ebiten.GamepadIDs()
		if len(gpIDs) <= 0 {
			return errNoGamePad
		}
		g.gamepadID = gpIDs[0]
		g.field = world.Tuto1Field
		g.state = tuto1
		g.tutoStep = 0
		g.tutoFrame = 0
		return nil

	case tuto1:
		g.setCameraPosition()
		g.updateAnimation()
		g.updateTuto1()
		return nil

	case tuto2:
		g.state = tuto3

	case tuto3:
		g.state = tuto4

	case tuto4:
		g.state = playingBlue

	case playingBlue, playingPink, playingWhite:
		// get current player
		var currentCharacter *character
		switch g.state {
		case playingBlue:
			currentCharacter = &g.blueCharacter
		case playingWhite:
			currentCharacter = &g.whiteCharacter
		case playingPink:
			currentCharacter = &g.pinkCharacter
		}
		// perform its move
		switch {
		case ebiten.IsGamepadButtonPressed(g.gamepadID, ebiten.GamepadButton(12)):
			g.tryRightMove(currentCharacter)
		case ebiten.IsGamepadButtonPressed(g.gamepadID, ebiten.GamepadButton(14)):
			g.tryLeftMove(currentCharacter)
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(11)):
			g.tryUpMove(currentCharacter)
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(13)):
			g.tryDownMove(currentCharacter)
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(5)):
			g.tryRightSwitch(currentCharacter)
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(4)):
			g.tryLeftSwitch(currentCharacter)
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(2)):
			currentCharacter.state = strike
			currentCharacter.strikeCurrentFrame = 0
			switch g.state {
			case playingBlue:
				g.state = blueStrike
			case playingPink:
				g.state = pinkStrike
			case playingWhite:
				g.state = whiteStrike
			}
		default:
			currentCharacter.state = idle
		}

	case blueSpecialMove:
		g.performBlueSpecialMove()

	case pinkSpecialMoveDirectDown:
		g.performPinkSpecialMoveDirectDown()

	case pinkSpecialMoveDirectUp:
		g.performPinkSpecialMoveDirectUp()

	case pinkSpecialMoveDown:
		g.performPinkSpecialMoveDown()

	case pinkSpecialMoveUp:
		g.performPinkSpecialMoveUp()

	case whiteSpecialMove, whiteSpecialMoveIdle:
		switch {
		case ebiten.IsGamepadButtonPressed(g.gamepadID, ebiten.GamepadButton(11)):
			g.tryWhiteSpecialMoveUp()
		case ebiten.IsGamepadButtonPressed(g.gamepadID, ebiten.GamepadButton(13)):
			g.tryWhiteSpecialMoveDown()
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(5)):
			g.tryWhiteSpecialSwitchRight()
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(4)):
			g.tryWhiteSpecialSwitchLeft()
		default:
			g.state = whiteSpecialMoveIdle
		}

	case blueStrike, pinkStrike, whiteStrike:
		var currentCharacter *character
		var nextState gameState
		switch g.state {
		case blueStrike:
			currentCharacter = &g.blueCharacter
			nextState = playingBlue
		case pinkStrike:
			currentCharacter = &g.pinkCharacter
			nextState = playingPink
		case whiteStrike:
			currentCharacter = &g.whiteCharacter
			nextState = playingWhite
		}
		currentCharacter.strikeCurrentFrame++
		if currentCharacter.strikeCurrentFrame >= currentCharacter.strikeNumFrames {
			if currentCharacter.facing == right {
				g.strikeEffectOnField(currentCharacter.x, currentCharacter.y, 1)
			} else {
				g.strikeEffectOnField(currentCharacter.x, currentCharacter.y, -1)
			}
			currentCharacter.state = idle
			g.state = nextState
		}
	}

	// update animations
	g.updateAnimation()

	// update camera
	g.setCameraPosition()

	// update minimap
	g.updateMinimap()

	return nil
}
