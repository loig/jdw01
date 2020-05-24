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
		gpID := getGamepadPlugged()
		switch g.tutoStep {
		case 0:
			if gpID != -1 {
				g.gamepadID = gpID
				g.tutoStep++
			}
		case 1:
			buttonPressed := g.getButtonPressed()
			if buttonPressed >= 0 {
				g.buttons.left = buttonPressed
				g.tutoStep++
			}
		case 2:
			buttonPressed := g.getButtonPressed()
			if buttonPressed >= 0 {
				g.buttons.right = buttonPressed
				g.tutoStep++
			}
		case 3:
			buttonPressed := g.getButtonPressed()
			if buttonPressed >= 0 {
				g.buttons.up = buttonPressed
				g.tutoStep++
			}
		case 4:
			buttonPressed := g.getButtonPressed()
			if buttonPressed >= 0 {
				g.buttons.down = buttonPressed
				g.tutoStep++
			}
		case 5:
			buttonPressed := g.getButtonPressed()
			if buttonPressed >= 0 {
				g.buttons.x = buttonPressed
				g.tutoStep++
			}
		case 6:
			buttonPressed := g.getButtonPressed()
			if buttonPressed >= 0 {
				g.buttons.lb = buttonPressed
				g.tutoStep++
			}
		case 7:
			buttonPressed := g.getButtonPressed()
			if buttonPressed >= 0 {
				g.buttons.rb = buttonPressed
				g.tutoStep++
			}
		default:
			g.field = world.Tuto1Field
			g.state = tuto1
			g.tutoStep = 0
			g.tutoFrame = 0
		}
		return nil

	case theEnd:
		if g.tutoStep == 0 {
			g.tutoFrame++
			if g.tutoFrame > 200 {
				g.tutoStep++
				g.tutoFrame = 0
			}
		} else {
			if inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(g.buttons.x)) {
				return errEndGame
			}
		}
		return nil

	case tuto1:
		if inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(g.buttons.x)) {
			g.tutoStep = 0
			g.tutoFrame = 0
			g.field = world.Tuto2Field
			g.blueCharacter.x = world.BlueXTuto2
			g.blueCharacter.y = world.BlueYTuto2
			g.blueCharacter.state = idle
			g.blueCharacter.animationFrame = 0
			g.blueCharacter.animationStep = 0
			g.whiteCharacter.x = world.WhiteXTuto2
			g.whiteCharacter.y = world.WhiteYTuto2
			g.whiteCharacter.state = idle
			g.whiteCharacter.animationFrame = 0
			g.whiteCharacter.animationStep = 0
			g.pinkCharacter.x = world.PinkXTuto2
			g.pinkCharacter.y = world.PinkYTuto2
			g.pinkCharacter.state = idle
			g.pinkCharacter.animationFrame = 0
			g.pinkCharacter.animationStep = 0
			g.state = tuto2
		} else {
			g.setCameraPosition()
			g.updateAnimation()
			g.updateTuto1()
		}
		return nil

	case tuto2:
		if inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(g.buttons.x)) {
			g.tutoStep = 0
			g.tutoFrame = 0
			g.field = world.GetTuto3Field()
			g.blueCharacter.x = world.BlueXTuto3
			g.blueCharacter.y = world.BlueYTuto3
			g.blueCharacter.state = idle
			g.blueCharacter.animationFrame = 0
			g.blueCharacter.animationStep = 0
			g.whiteCharacter.x = world.WhiteXTuto3
			g.whiteCharacter.y = world.WhiteYTuto3
			g.whiteCharacter.state = idle
			g.whiteCharacter.animationFrame = 0
			g.whiteCharacter.animationStep = 0
			g.pinkCharacter.x = world.PinkXTuto3
			g.pinkCharacter.y = world.PinkYTuto3
			g.pinkCharacter.state = idle
			g.pinkCharacter.animationFrame = 0
			g.pinkCharacter.animationStep = 0
			g.state = tuto3
		} else {
			g.setCameraPosition()
			g.updateAnimation()
			g.updateTuto2()
		}
		return nil

	case tuto3:
		if inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(g.buttons.x)) {
			g.tutoStep = 0
			g.tutoFrame = 0
			g.field = world.GetTuto4Field()
			g.blueCharacter.x = world.BlueXTuto4
			g.blueCharacter.y = world.BlueYTuto4
			g.blueCharacter.state = idle
			g.blueCharacter.animationFrame = 0
			g.blueCharacter.animationStep = 0
			g.whiteCharacter.x = world.WhiteXTuto4
			g.whiteCharacter.y = world.WhiteYTuto4
			g.whiteCharacter.state = idle
			g.whiteCharacter.animationFrame = 0
			g.whiteCharacter.animationStep = 0
			g.pinkCharacter.x = world.PinkXTuto4
			g.pinkCharacter.y = world.PinkYTuto4
			g.pinkCharacter.state = idle
			g.pinkCharacter.animationFrame = 0
			g.pinkCharacter.animationStep = 0
			g.state = tuto4
		} else {
			g.setCameraPosition()
			g.updateAnimation()
			g.updateTuto3()
		}
		return nil

	case tuto4:
		if inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(g.buttons.x)) {
			g.field = g.world
			g.blueCharacter.x = g.blueStartX
			g.blueCharacter.y = g.blueStartY
			g.blueCharacter.speed = 0.11
			g.blueCharacter.state = idle
			g.blueCharacter.animationFrame = 0
			g.blueCharacter.animationStep = 0
			g.whiteCharacter.x = g.whiteStartX
			g.whiteCharacter.y = g.whiteStartY
			g.whiteCharacter.speed = 0.09
			g.whiteCharacter.state = idle
			g.whiteCharacter.animationFrame = 0
			g.whiteCharacter.animationStep = 0
			g.pinkCharacter.x = g.pinkStartX
			g.pinkCharacter.y = g.pinkStartY
			g.pinkCharacter.speed = 0.13
			g.pinkCharacter.state = idle
			g.pinkCharacter.animationFrame = 0
			g.pinkCharacter.animationStep = 0
			g.state = playingBlue
			g.tutoStep = 0
			g.tutoFrame = 0
		} else {
			g.setCameraPosition()
			g.updateAnimation()
			g.updateTuto4()
		}
		return nil

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
		case ebiten.IsGamepadButtonPressed(g.gamepadID, ebiten.GamepadButton(g.buttons.right)):
			g.tryRightMove(currentCharacter)
		case ebiten.IsGamepadButtonPressed(g.gamepadID, ebiten.GamepadButton(g.buttons.left)):
			g.tryLeftMove(currentCharacter)
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(g.buttons.up)):
			g.tryUpMove(currentCharacter)
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(g.buttons.down)):
			g.tryDownMove(currentCharacter)
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(g.buttons.rb)):
			g.tryRightSwitch(currentCharacter)
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(g.buttons.lb)):
			g.tryLeftSwitch(currentCharacter)
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(g.buttons.x)):
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
		case ebiten.IsGamepadButtonPressed(g.gamepadID, ebiten.GamepadButton(g.buttons.up)):
			g.tryWhiteSpecialMoveUp()
		case ebiten.IsGamepadButtonPressed(g.gamepadID, ebiten.GamepadButton(g.buttons.down)):
			g.tryWhiteSpecialMoveDown()
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(g.buttons.rb)):
			g.tryWhiteSpecialSwitchRight()
		case inpututil.IsGamepadButtonJustPressed(g.gamepadID, ebiten.GamepadButton(g.buttons.lb)):
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

	// update sound
	g.updateSound()

	// update minimap
	g.updateMinimap()

	return nil
}
