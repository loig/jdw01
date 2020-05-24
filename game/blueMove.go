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

func (g *Game) performBlueSpecialMove() {
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
			if g.state != tuto1 {
				g.state = playingBlue
			}
			g.blueCharacter.state = idle
		}
	}
}
