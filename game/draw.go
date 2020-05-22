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
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var op *ebiten.DrawImageOptions
var sub image.Rectangle

// Draw implements one of the required methods
// for the ebiten.Game interface
func (g *Game) Draw(screen *ebiten.Image) {

	switch g.state {
	case initGame:
		return
	default:
		// get bounds
		xmin, ymin, xmax, ymax := g.visibleRectangle()

		// Draw the background
		op = &ebiten.DrawImageOptions{}
		screen.DrawImage(dayBackgroundImage, op)

		// Draw the field
		for y := ymin; y < ymax; y++ {
			for x := xmin; x < xmax; x++ {
				op = &ebiten.DrawImageOptions{}
				op.GeoM.Scale(2, 2)
				op.GeoM.Translate(float64(x)*32, float64(y)*32)
				g.applyCamera(op)
				sub = image.Rect(
					16*g.field[y][x].tile.lookX, 16*g.field[y][x].tile.lookY,
					16+16*g.field[y][x].tile.lookX, 16+16*g.field[y][x].tile.lookY,
				)
				screen.DrawImage(tilesImage.SubImage(sub).(*ebiten.Image), op)
				sub = image.Rect(
					16*g.field[y][x].decoration.lookX, 16*g.field[y][x].decoration.lookY,
					16+16*g.field[y][x].decoration.lookX, 16+16*g.field[y][x].decoration.lookY,
				)
				screen.DrawImage(tilesImage.SubImage(sub).(*ebiten.Image), op)
			}
		}

		// Draw the blue guy
		op = &ebiten.DrawImageOptions{}
		if g.blueCharacter.facing == left {
			var mirorM ebiten.GeoM
			mirorM.SetElement(0, 0, -1)
			mirorM.SetElement(1, 1, 1)
			op.GeoM.Concat(mirorM)
			op.GeoM.Translate(32, 0)
		}
		op.GeoM.Translate((g.blueCharacter.x)*32, (g.blueCharacter.y)*32)
		g.applyCamera(op)
		sub = image.Rect(
			0+32*g.blueCharacter.animationStep, 0+32*int(g.blueCharacter.state),
			32+32*g.blueCharacter.animationStep, 32+32*int(g.blueCharacter.state),
		)
		screen.DrawImage(blueCharacterImage.SubImage(sub).(*ebiten.Image), op)

		// Draw the white guy
		op = &ebiten.DrawImageOptions{}
		if g.whiteCharacter.facing == left {
			var mirorM ebiten.GeoM
			mirorM.SetElement(0, 0, -1)
			mirorM.SetElement(1, 1, 1)
			op.GeoM.Concat(mirorM)
			op.GeoM.Translate(32, 0)
		}
		op.GeoM.Translate((g.whiteCharacter.x)*32, (g.whiteCharacter.y)*32)
		g.applyCamera(op)
		sub = image.Rect(
			0+32*g.whiteCharacter.animationStep, 0+32*int(g.whiteCharacter.state),
			32+32*g.whiteCharacter.animationStep, 32+32*int(g.whiteCharacter.state),
		)
		screen.DrawImage(whiteCharacterImage.SubImage(sub).(*ebiten.Image), op)

		// Draw the pink guy
		op = &ebiten.DrawImageOptions{}
		if g.pinkCharacter.facing == left {
			var mirorM ebiten.GeoM
			mirorM.SetElement(0, 0, -1)
			mirorM.SetElement(1, 1, 1)
			op.GeoM.Concat(mirorM)
			op.GeoM.Translate(32, 0)
		}
		op.GeoM.Translate((g.pinkCharacter.x)*32, (g.pinkCharacter.y)*32)
		g.applyCamera(op)
		sub = image.Rect(
			0+32*g.pinkCharacter.animationStep, 0+32*int(g.pinkCharacter.state),
			32+32*g.pinkCharacter.animationStep, 32+32*int(g.pinkCharacter.state),
		)
		screen.DrawImage(pinkCharacterImage.SubImage(sub).(*ebiten.Image), op)

		// DEBUG
		ebitenutil.DrawLine(screen, float64(g.screenWidth)/2, 0, float64(g.screenWidth)/2, float64(g.screenHeight), color.White)
		ebitenutil.DrawLine(screen, 0, float64(g.screenHeight)/2, float64(g.screenWidth), float64(g.screenHeight)/2, color.White)
		s := fmt.Sprint("FPS: ", ebiten.CurrentFPS(), "\n", "TPS: ", ebiten.CurrentTPS(), "\n", "Camera: ", xmin, ", ", ymin, ", ", xmax, ",", ymax)
		ebitenutil.DebugPrint(screen, s)
	}
}
