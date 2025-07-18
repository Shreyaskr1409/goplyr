package player

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

type Player struct {
	audioContext *oto.Context
	player       oto.Player
	pauseChan    chan bool
	isPaused     bool
	isPlaying    bool
	filePath     string
}

func InitPlayer() (*Player, error) {
	op := &oto.NewContextOptions{
		SampleRate:   44100,
		ChannelCount: 2,
		Format:       oto.FormatSignedInt16LE,
	}
	audioContext, readyChan, err := oto.NewContext(op.SampleRate, op.ChannelCount, op.Format)
	if err != nil {
		return nil, err
	}
	<-readyChan // pauses till the channel is ready

	return &Player{
		audioContext: audioContext,
		pauseChan:    make(chan bool),
		isPaused:     false,
		isPlaying:    false,
	}, nil
}

func (p *Player) PlayFile(filePath string) error {
	if p.isPlaying {
		p.Stop()
	}

	p.filePath = filePath

	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error encountered: ", err)
		return errors.New(fmt.Sprint("Error opening file: ", err))
	}

	// Decode MP3 file to raw PCM
	decoder, err := mp3.NewDecoder(file)
	if err != nil {
		log.Println("Error decoding MP3: ", err)
		return errors.New(fmt.Sprint("Error decoding MP3 file: ", err))
	}

	p.player = p.audioContext.NewPlayer(decoder)

	go p.playRoutine()

	p.isPlaying = true
	p.isPaused = false

	return nil
}

func (p *Player) playRoutine() {
	if p.player == nil {
		return
	}

	p.player.Play()

	for {
		select {
		case pause := <-p.pauseChan:
			if pause {
				p.player.Pause()
			} else {
				p.player.Play()
			}
		case <-time.After(100 * time.Millisecond):
			if !p.player.IsPlaying() {
				p.isPlaying = false
				p.isPaused = false
				return
			}
		}
	}
}
