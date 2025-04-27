package player

import "log"

func (p *Player) Pause() {
	if p.isPlaying && !p.isPaused {
		p.pauseChan <- true
		p.isPaused = true
	}
}

func (p *Player) Resume() {
	if p.isPlaying && p.isPaused {
		p.pauseChan <- false
		p.isPaused = false
	}
}

func (p *Player) Stop() {
	if p.player != nil {
		err := p.player.Close()
		if err != nil {
			log.Println("Error in closing the player: ", err)
		}
		p.player = nil
	}
	p.isPaused = false
	p.isPlaying = false
}

func (p *Player) IsPaused() bool {
	return p.isPaused
}

func (p *Player) IsPlaying() bool {
	return p.isPlaying
}
