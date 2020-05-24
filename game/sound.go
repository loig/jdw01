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
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type audioManagerInfo struct {
	audioContext *audio.Context
	soundPlayer  *audio.Player
	currentSound soundType
}

type soundType int

const (
	moveSound soundType = iota
	pinkSound
	blueSound
	noSound
)

var moveSoundBytes []byte

func (g *Game) initSound() (err error) {
	g.audioManager.audioContext, err = audio.NewContext(44100)

	soundFile, err := ebitenutil.OpenFile("assets/move.mp3")
	sound, err := mp3.Decode(g.audioManager.audioContext, soundFile)
	if err != nil {
		return err
	}
	moveSoundBytes, err = ioutil.ReadAll(sound)
	if err != nil {
		return err
	}

	return err
}

func (g *Game) updateSound() {
	soundToPlay := noSound
	if g.state == playingBlue && g.blueCharacter.state == move {
		soundToPlay = moveSound
	}
	fmt.Println(soundToPlay, g.audioManager.currentSound)

	if soundToPlay != g.audioManager.currentSound {
		if g.audioManager.soundPlayer != nil {
			g.audioManager.soundPlayer.Pause()
		}
		switch soundToPlay {
		case noSound:
			g.audioManager.soundPlayer = nil
		case moveSound:
			g.audioManager.soundPlayer, _ = audio.NewPlayerFromBytes(
				g.audioManager.audioContext, moveSoundBytes,
			)
			g.audioManager.soundPlayer.Play()
		}
		g.audioManager.currentSound = soundToPlay
		return
	}

	if g.audioManager.soundPlayer != nil && !g.audioManager.soundPlayer.IsPlaying() {
		g.audioManager.soundPlayer.Rewind()
		g.audioManager.soundPlayer.Play()
	}
}
