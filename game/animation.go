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

// general animation management
func (g *Game) initAnimation() {
	g.blueCharacterInitAnimation()
	g.whiteCharacterInitAnimation()
	g.pinkCharacterInitAnimation()
}

func (g *Game) updateAnimation() {
	g.blueCharacter.updateCharacterAnimation()
	g.pinkCharacter.updateCharacterAnimation()
	if g.whiteCharacter.state != specialMove || g.state == whiteSpecialMove || g.state == tuto2 {
		g.whiteCharacter.updateCharacterAnimation()
	}
}

// blue character animation management
func (g *Game) blueCharacterInitAnimation() {
	g.blueCharacter.animationFrame = 0
	g.blueCharacter.animationStep = 0
	g.blueCharacter.idleFrames = []int{12, 9, 9, 12}
	g.blueCharacter.moveFrames = []int{5, 5, 5, 5, 5, 5}
	g.blueCharacter.specialMoveFrames = []int{7, 4, 4, 4, 4, 4, 1, 5, 5, 5, 5, 5, 5}
	for i := 0; i < 7; i++ {
		g.blueCharacter.specialMoveNumFrames += g.blueCharacter.specialMoveFrames[i]
	}
	g.blueCharacter.strikeFrames = []int{2, 2, 6, 4, 10, 2}
	for _, numFrames := range g.blueCharacter.strikeFrames {
		g.blueCharacter.strikeNumFrames += numFrames
	}
	g.blueCharacter.previousState = idle
}

// white character animation management
func (g *Game) whiteCharacterInitAnimation() {
	g.whiteCharacter.animationFrame = 0
	g.whiteCharacter.animationStep = 0
	g.whiteCharacter.idleFrames = []int{13, 15, 15, 13}
	g.whiteCharacter.moveFrames = []int{7, 7, 7, 7, 7, 7}
	g.whiteCharacter.specialMoveFrames = []int{5, 5, 5, 5}
	g.whiteCharacter.strikeFrames = []int{2, 5, 7, 12, 2}
	for _, numFrames := range g.whiteCharacter.strikeFrames {
		g.whiteCharacter.strikeNumFrames += numFrames
	}
	g.whiteCharacter.previousState = idle
}

// pink character animation management
func (g *Game) pinkCharacterInitAnimation() {
	g.pinkCharacter.animationFrame = 0
	g.pinkCharacter.animationStep = 0
	g.pinkCharacter.idleFrames = []int{9, 9, 9, 9}
	g.pinkCharacter.moveFrames = []int{4, 4, 4, 4, 4, 4}
	g.pinkCharacter.specialMoveFrames = []int{3, 3, 3, 3, 3, 3, 3, 3}
	for _, numFrames := range g.pinkCharacter.specialMoveFrames {
		g.pinkCharacter.specialMoveNumFrames += numFrames
	}
	g.pinkCharacter.strikeFrames = []int{2, 2, 5, 3, 5, 3, 5, 2}
	for _, numFrames := range g.pinkCharacter.strikeFrames {
		g.pinkCharacter.strikeNumFrames += numFrames
	}
	g.pinkCharacter.previousState = idle
}

// general character animation management
func (c *character) updateCharacterAnimation() {

	if c.state != c.previousState {
		c.animationFrame = 0
		c.animationStep = 0
		c.previousState = c.state
		return
	}

	c.animationFrame++

	var frames []int
	switch c.state {
	case idle:
		frames = c.idleFrames
	case move:
		frames = c.moveFrames
	case specialMove:
		frames = c.specialMoveFrames
	case strike:
		frames = c.strikeFrames
	}

	if c.animationFrame >= frames[c.animationStep] {
		c.animationFrame = 0
		c.animationStep = (c.animationStep + 1) % len(frames)
	}

}
