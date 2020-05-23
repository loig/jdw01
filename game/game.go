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
	"errors"

	"github.com/loig/jdw01/world"
)

// Game implements ebiten.Game interface
type Game struct {
	state          gameState
	gamepadID      int
	camera         cameraInfo
	screenWidth    int
	screenHeight   int
	blueCharacter  character
	whiteCharacter character
	pinkCharacter  character
	field          [][]world.FieldTile
}

// Current state of the game
type gameState int

const (
	initGame gameState = iota
	playingBlue
	blueSpecialMove
	blueStrike
	playingWhite
	whiteSpecialMove
	whiteSpecialMoveIdle
	whiteStrike
	playingPink
	pinkSpecialMoveDirectDown
	pinkSpecialMoveDirectUp
	pinkSpecialMoveDown
	pinkSpecialMoveUp
	pinkStrike
)

// Errors for communications with main program
var errEndGame = errors.New("End")
var errNoGamePad = errors.New("GamePad")

// Init initializes the Game structure
func (g *Game) Init() (err error) {

	// Load graphics and sounds
	err = loadAssets()
	if err != nil {
		return err
	}

	// Set screen size
	g.screenWidth = 1024
	g.screenHeight = 512

	// Setup animations
	g.initAnimation()

	// World size
	fieldWidth := 200
	fieldHeight := 100

	// Set field
	var (
		blueX, blueY, pinkX, pinkY, whiteX, whiteY float64
	)
	g.field, blueX, blueY, pinkX, pinkY, whiteX, whiteY = world.GenerateField(fieldWidth, fieldHeight)

	// Initialize minimap
	if err = g.initMiniMap(); err != nil {
		return err
	}

	// Initial positions of characters
	g.blueCharacter.x = blueX
	g.blueCharacter.y = blueY
	g.blueCharacter.state = idle
	g.blueCharacter.speed = 0.11

	g.whiteCharacter.x = whiteX
	g.whiteCharacter.y = whiteY
	g.whiteCharacter.state = idle
	g.whiteCharacter.speed = 0.09
	g.whiteCharacter.facing = left

	g.pinkCharacter.x = pinkX
	g.pinkCharacter.y = pinkY
	g.pinkCharacter.state = idle
	g.pinkCharacter.speed = 0.13

	// Set camera
	g.setCameraPosition()

	// Set initial game state
	g.state = initGame

	return nil

}
