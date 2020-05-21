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
	field          [][]fieldTile
}

// Current state of the game
type gameState int

const (
	initGame gameState = iota
	playingBlue
	blueSpecialMovePhase1
	blueSpecialMovePhase2
	playingWhite
	whiteSpecialMove
	playingPink
	pinkSpecialMove
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
	g.screenWidth = 600
	g.screenHeight = 400

	// Setup animations
	g.initAnimation()

	// Initial positions of characters
	g.blueCharacter.x = 3
	g.blueCharacter.y = 3
	g.blueCharacter.state = idle
	g.blueCharacter.speed = 0.1

	g.whiteCharacter.x = g.blueCharacter.x
	g.whiteCharacter.y = g.blueCharacter.y - 3
	g.whiteCharacter.state = idle
	g.whiteCharacter.speed = 0.08

	g.pinkCharacter.x = g.blueCharacter.x
	g.pinkCharacter.y = g.blueCharacter.y + 3
	g.pinkCharacter.state = idle
	g.pinkCharacter.speed = 0.12

	// Set field
	g.setInitialField()

	// Set camera
	g.setCameraPosition()

	// Set initial game state
	g.state = initGame

	return nil

}
