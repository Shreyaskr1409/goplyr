package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Window interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
}

type StatusBar struct {
	song        *Song
	playerState *PlayerState
}

type Song struct {
	song        string
	album       string
	artist      string
	duration    string
	albumArtURI string
	year        string
	filepath    string
}

const (
	NONE PlaybackMode = iota
	REPEAT
	REPEAT_ALL
	SHUFFLE
)

const (
	PAUSE PlaybackStatus = iota
	PLAY
)

type (
	PlaybackMode   int
	PlaybackStatus int
	PlayerState    struct {
		playbackMode   *PlaybackMode
		playbackStatus *PlaybackStatus
		timestamp      string
		volume         int
	}
)
