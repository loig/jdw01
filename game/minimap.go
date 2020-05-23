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
	"math"

	"github.com/hajimehoshi/ebiten"
)

var miniMap *ebiten.Image
var miniMapOverlay *ebiten.Image

func (g *Game) initMiniMap() (err error) {

	miniMap, err = ebiten.NewImage(len(g.field)*16+32, len(g.field[0])*16+32, ebiten.FilterDefault)
	if err != nil {
		return err
	}
	miniMapOverlay, err = ebiten.NewImage(len(g.field)*16+32, len(g.field[0])*16+32, ebiten.FilterDefault)
	if err != nil {
		return err
	}

	miniMap.Fill(color.Black)
	// Draw the field
	for y := 0; y < len(g.field); y++ {
		for x := 0; x < len(g.field[0]); x++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x)*16+16, float64(y)*16+16)
			sub = image.Rect(
				16*g.field[y][x].Tile.LookX, 16*g.field[y][x].Tile.LookY,
				16+16*g.field[y][x].Tile.LookX, 16+16*g.field[y][x].Tile.LookY,
			)
			miniMap.DrawImage(tilesImage.SubImage(sub).(*ebiten.Image), op)
			sub = image.Rect(
				16*g.field[y][x].Decoration.LookX, 16*g.field[y][x].Decoration.LookY,
				16+16*g.field[y][x].Decoration.LookX, 16+16*g.field[y][x].Decoration.LookY,
			)
			miniMap.DrawImage(tilesImage.SubImage(sub).(*ebiten.Image), op)
		}
	}

	return err
}

func (g *Game) updateMinimap() {
	if miniMapOverlay.Clear() != nil {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(8, 8)
	op.GeoM.Translate(math.Round(g.whiteCharacter.x)*16, math.Round(g.whiteCharacter.y)*16)
	sub := image.Rect(
		11, 73,
		15, 77,
	)
	miniMapOverlay.DrawImage(whiteCharacterImage.SubImage(sub).(*ebiten.Image), op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(16, 16)
	op.GeoM.Translate(math.Round(g.blueCharacter.x)*16, math.Round(g.blueCharacter.y)*16)
	sub = image.Rect(
		8, 108,
		10, 110,
	)
	miniMapOverlay.DrawImage(blueCharacterImage.SubImage(sub).(*ebiten.Image), op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(8, 8)
	op.GeoM.Translate(math.Round(g.pinkCharacter.x)*16, math.Round(g.pinkCharacter.y)*16)
	sub = image.Rect(
		8, 13,
		12, 17,
	)
	miniMapOverlay.DrawImage(pinkCharacterImage.SubImage(sub).(*ebiten.Image), op)

}
