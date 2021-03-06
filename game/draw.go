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
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/loig/jdw01/world"
)

var op *ebiten.DrawImageOptions
var sub image.Rectangle
var endImage *ebiten.Image

// Draw implements one of the required methods
// for the ebiten.Game interface
func (g *Game) Draw(screen *ebiten.Image) {

	switch {
	case g.state == initGame:
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Scale(2, 2)
		op.GeoM.Translate(float64(g.screenWidth)/2-100, float64(g.screenHeight)/2-100)
		switch g.tutoStep {
		case 0:
			ebitenutil.DebugPrint(screen, "No gamepad")
		case 1:
			ebitenutil.DebugPrint(screen, "Press left")
			screen.DrawImage(buttonLeftImage, op)
		case 2:
			ebitenutil.DebugPrint(screen, "Press right")
			screen.DrawImage(buttonRightImage, op)
		case 3:
			ebitenutil.DebugPrint(screen, "Press up")
			screen.DrawImage(buttonUpImage, op)
		case 4:
			ebitenutil.DebugPrint(screen, "Press down")
			screen.DrawImage(buttonDownImage, op)
		case 5:
			ebitenutil.DebugPrint(screen, "Press X")
			screen.DrawImage(buttonXImage, op)
		case 6:
			ebitenutil.DebugPrint(screen, "Press LB")
			screen.DrawImage(buttonLImage, op)
		case 7:
			ebitenutil.DebugPrint(screen, "Press RB")
			screen.DrawImage(buttonRImage, op)
		}
		return

	case g.state == theEnd && (g.tutoFrame > 2 || g.tutoStep > 0):
		op = &ebiten.DrawImageOptions{}
		saturationScale := 0.1
		valueScale := 0.2
		if g.tutoStep == 0 {
			saturationScale = float64(200-g.tutoFrame)/200 + 0.1
			if saturationScale > 1 {
				saturationScale = 1
			}
			valueScale = float64(200-g.tutoFrame)/200 + 0.2
			if valueScale > 1 {
				valueScale = 1
			}
		}
		op.ColorM.ChangeHSV(0, saturationScale, valueScale)
		screen.DrawImage(endImage, op)
		if g.tutoStep != 0 {
			text.Draw(screen, "Home!", endFont, 400, 350, color.White)
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Scale(0.5, 0.5)
			op.GeoM.Translate(float64(g.screenWidth)-10-50, float64(g.screenHeight)-10-50)
			screen.DrawImage(buttonXImage, op)
		}

	default:
		// get bounds
		xmin, ymin, xmax, ymax := g.visibleRectangle()

		// Draw the background
		op = &ebiten.DrawImageOptions{}
		screen.DrawImage(dayBackgroundImage, op)

		// Draw the field
		if ymin < 0 {
			ymin = 0
		}
		for y := ymin; y < ymax; y++ {
			for x := xmin; x < xmax; x++ {
				op = &ebiten.DrawImageOptions{}
				op.GeoM.Scale(2, 2)
				op.GeoM.Translate(float64(x)*32, float64(y)*32)
				g.applyCamera(op)
				if x < 0 || x >= len(g.field[0]) || y >= len(g.field) {
					if g.state != tuto1 && g.state != tuto2 && g.state != tuto3 && g.state != tuto4 {
						sub = image.Rect(
							16*world.DummyX, 16*world.DummyY,
							16+16*world.DummyX, 16+16*world.DummyY,
						)
						screen.DrawImage(tilesImage.SubImage(sub).(*ebiten.Image), op)
					}
				} else {
					sub = image.Rect(
						16*g.field[y][x].Tile.LookX, 16*g.field[y][x].Tile.LookY,
						16+16*g.field[y][x].Tile.LookX, 16+16*g.field[y][x].Tile.LookY,
					)
					screen.DrawImage(tilesImage.SubImage(sub).(*ebiten.Image), op)
					sub = image.Rect(
						16*g.field[y][x].Decoration.LookX, 16*g.field[y][x].Decoration.LookY,
						16+16*g.field[y][x].Decoration.LookX, 16+16*g.field[y][x].Decoration.LookY,
					)
					screen.DrawImage(tilesImage.SubImage(sub).(*ebiten.Image), op)
				}
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

		// Draw the front layout of the field
		if ymax > len(g.field) {
			ymax = len(g.field)
		}
		if xmin < 0 {
			xmin = 0
		}
		if xmax > len(g.field[0]) {
			xmax = len(g.field[0])
		}
		for y := ymin; y < ymax; y++ {
			for x := xmin; x < xmax; x++ {
				op = &ebiten.DrawImageOptions{}
				op.GeoM.Scale(2, 2)
				op.GeoM.Translate(float64(x)*32, float64(y)*32)
				g.applyCamera(op)
				sub = image.Rect(
					16*g.field[y][x].Info.LookX, 16*g.field[y][x].Info.LookY,
					16+16*g.field[y][x].Info.LookX, 16+16*g.field[y][x].Info.LookY,
				)
				screen.DrawImage(tilesImage.SubImage(sub).(*ebiten.Image), op)
			}
		}

		if g.state != tuto1 && g.state != tuto2 && g.state != tuto3 && g.state != tuto4 && g.state != theEnd {
			// Draw the minimap
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Scale(0.125, 0.125)
			op.GeoM.Translate(5, float64(g.screenHeight)-5-(float64(len(g.field))*16+32)*0.125)
			screen.DrawImage(miniMap, op)
			screen.DrawImage(miniMapOverlay, op)
		} else {
			// Draw buttons
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Scale(1.25, 1.25)
			switch g.state {
			case tuto1:
				screen.DrawImage(buttonRightImage, op)
			case tuto2:
				if g.tutoStep == 0 {
					screen.DrawImage(buttonUpImage, op)
				} else {
					screen.DrawImage(buttonDownImage, op)
				}
			case tuto3:
				screen.DrawImage(buttonXImage, op)
			case tuto4:
				screen.DrawImage(buttonRImage, op)
			}
			if g.state != theEnd {
				op = &ebiten.DrawImageOptions{}
				op.GeoM.Scale(0.5, 0.5)
				op.GeoM.Translate(float64(g.screenWidth)-10-50, float64(g.screenHeight)-10-50)
				screen.DrawImage(buttonXImage, op)
			}
		}

		if g.state == theEnd && g.tutoFrame == 2 {
			endImage, _ = ebiten.NewImage(g.screenWidth, g.screenHeight, ebiten.FilterDefault)
			op = &ebiten.DrawImageOptions{}
			endImage.DrawImage(screen, op)
		}

		// DEBUG
		/*
			ebitenutil.DrawLine(screen, float64(g.screenWidth)/2, 0, float64(g.screenWidth)/2, float64(g.screenHeight), color.White)
			ebitenutil.DrawLine(screen, 0, float64(g.screenHeight)/2, float64(g.screenWidth), float64(g.screenHeight)/2, color.White)
			s := fmt.Sprint("FPS: ", ebiten.CurrentFPS(), "\n", "TPS: ", ebiten.CurrentTPS(), "\n", "Camera: ", xmin, ", ", ymin, ", ", xmax, ",", ymax, "\n", "Blue: ", g.blueCharacter.x, ", ", g.blueCharacter.y)
			ebitenutil.DebugPrint(screen, s)
		*/
	}

}
