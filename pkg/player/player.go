package player

import "github.com/hajimehoshi/oto"

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
	p.player.Close()
}
