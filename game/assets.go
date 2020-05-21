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
	_ "image/png" //for ebitenutil.NewImageFromFile

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var blueCharacterImage *ebiten.Image
var whiteCharacterImage *ebiten.Image
var pinkCharacterImage *ebiten.Image

func loadAssets() (err error) {

	blueCharacterImage, _, err = ebitenutil.NewImageFromFile("assets/dude.png", ebiten.FilterDefault)
	if err != nil {
		return err
	}

	whiteCharacterImage, _, err = ebitenutil.NewImageFromFile("assets/owlet.png", ebiten.FilterDefault)
	if err != nil {
		return err
	}

	pinkCharacterImage, _, err = ebitenutil.NewImageFromFile("assets/pink.png", ebiten.FilterDefault)
	if err != nil {
		return err
	}

	return nil
}
