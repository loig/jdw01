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
)

func (g *Game) tryRightMove(currentCharacter *character) {
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
	case blueOrPinkUpFieldMove:
		if g.state == playingBlue {
			g.state = blueSpecialMove
			currentCharacter.state = specialMove
			currentCharacter.specialMoveCurrentFrame = 0
		} else if g.state == playingPink {
			g.state = pinkSpecialMoveUp
			currentCharacter.state = specialMove
			currentCharacter.specialMoveCurrentFrame = 0
		} else {
			currentCharacter.state = idle
		}
	case pinkDownFieldMoveOrNormalMove:
		if g.state == playingPink {
			g.state = pinkSpecialMoveDown
			currentCharacter.state = specialMove
			currentCharacter.specialMoveCurrentFrame = 0
		} else {
			currentCharacter.x += currentCharacter.speed
			currentCharacter.state = move
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
}

func (g *Game) tryLeftMove(currentCharacter *character) {
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
	case blueOrPinkUpFieldMove:
		if g.state == playingBlue {
			g.state = blueSpecialMove
			currentCharacter.state = specialMove
			currentCharacter.specialMoveCurrentFrame = 0
		} else if g.state == playingPink {
			g.state = pinkSpecialMoveUp
			currentCharacter.state = specialMove
			currentCharacter.specialMoveCurrentFrame = 0
		} else {
			currentCharacter.state = idle
		}
	case pinkDownFieldMoveOrNormalMove:
		if g.state == playingPink {
			g.state = pinkSpecialMoveDown
			currentCharacter.state = specialMove
			currentCharacter.specialMoveCurrentFrame = 0
		} else {
			currentCharacter.x -= currentCharacter.speed
			currentCharacter.state = move
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
}

func (g *Game) tryUpMove(currentCharacter *character) {
	if g.state == playingWhite {
		if g.fieldOkForWhiteSpecialMove(currentCharacter.x, currentCharacter.y, -1) {
			g.state = whiteSpecialMove
			currentCharacter.state = specialMove
			currentCharacter.x = math.Round(currentCharacter.x)
		}
		return
	}
	if g.state == playingPink {
		if g.fieldOkForPinkSpecialMove(currentCharacter.x, currentCharacter.y, -1) {
			g.state = pinkSpecialMoveDirectUp
			currentCharacter.specialMoveCurrentFrame = 0
			currentCharacter.state = specialMove
			currentCharacter.x = math.Round(currentCharacter.x)
		}
	}
}

func (g *Game) tryDownMove(currentCharacter *character) {
	if g.state == playingWhite {
		if g.fieldOkForWhiteSpecialMove(currentCharacter.x, currentCharacter.y, 1) {
			g.state = whiteSpecialMove
			currentCharacter.state = specialMove
			currentCharacter.x = math.Round(currentCharacter.x)
			currentCharacter.y += 0.5
		}
		return
	}
	if g.state == playingPink {
		if g.fieldOkForPinkSpecialMove(currentCharacter.x, currentCharacter.y, 1) {
			g.state = pinkSpecialMoveDirectDown
			currentCharacter.specialMoveCurrentFrame = 0
			currentCharacter.state = specialMove
			currentCharacter.x = math.Round(currentCharacter.x)
		}
		return
	}
}
