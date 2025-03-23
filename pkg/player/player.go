package player

import (
	"github.com/hajimehoshi/oto"
	"log"
)

type Player struct {
	player   oto.Player
	pause    chan bool
	isPaused bool
}

func (p *Player) Pause() {
	if !p.isPaused {
		p.pause <- true
		p.isPaused = false
	}
}

func (p *Player) Play() {
	if p.isPaused {
		p.pause <- false
		p.isPaused = false
	}
}

func (p *Player) Stop() {
	err := p.player.Close()
	if err != nil {
		log.Fatal("Failed to stop the song")
	}
}
