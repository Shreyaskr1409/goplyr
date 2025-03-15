package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hajimehoshi/oto/v2"
)

type PlayerWindow struct {
	song        *Song
	playerState *PlayerState
	player      oto.Player
	pause       chan bool
	isPaused    bool
}

func (p *PlayerWindow) Init() tea.Cmd {
	return nil
}

func (p *PlayerWindow) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}

func (p *PlayerWindow) View() string {
	return "hello"
}

func InitPlayerWindow() *PlayerWindow {
	pw := &PlayerWindow{}
	playbackMode := NONE

	pw.song = &Song{}
	pw.playerState = &PlayerState{
		playbackMode: &playbackMode,
		timestamp:    "00:00",
		volume:       100,
	}

	return pw
}
