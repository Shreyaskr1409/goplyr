package cmd

import (
	"fmt"

	"github.com/Shreyaskr1409/goplyr/cmd/util"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hajimehoshi/oto/v2"
)

type PlayerWindow struct {
	song        *Song
	playerState *PlayerState
	player      oto.Player
	pause       chan bool
	isPaused    bool
	width       int
	height      int
	messages    []string
}

func InitPlayerWindow() *PlayerWindow {
	pw := &PlayerWindow{}
	playbackMode := NONE

	pw.song = &Song{
		song:        "Norway",
		album:       "Teen Dream",
		artist:      "Beach House",
		duration:    "3:54",
		year:        "2010",
		albumArtURI: "",
	}
	pw.playerState = &PlayerState{
		playbackMode: &playbackMode,
		timestamp:    "00:00",
		volume:       100,
	}

	return pw
}

func (p *PlayerWindow) Init() tea.Cmd {
	util.MsgAppendln(&p.messages, fmt.Sprint("SONG:     ", p.song.song))
	util.MsgAppendln(&p.messages, fmt.Sprint("ALBUM:    ", p.song.album))
	util.MsgAppendln(&p.messages, fmt.Sprint("ARTIST:   ", p.song.artist))
	util.MsgAppend(&p.messages, fmt.Sprint("DURATION: ", p.song.duration))
	// in above paragraph, in last line, i removed the new line part as it was creating
	// unnecessary new line when displayed in the bottom

	return nil
}

func (p *PlayerWindow) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return p, tea.Quit
		}

	case tea.WindowSizeMsg:
		p.width = msg.Width
		p.height = msg.Height
	}

	return p, nil
}

func (p *PlayerWindow) View() string {
	stylePage := lipgloss.NewStyle().Width(p.width).Height(p.height).Align(lipgloss.Left, lipgloss.Bottom)
	styleContent := lipgloss.NewStyle().AlignHorizontal(lipgloss.Left)
	page := ""
	s := ""

	for i := range p.messages {
		s = fmt.Sprint(s, p.messages[i])
	}
	page = fmt.Sprint(page, stylePage.Render(styleContent.Render(s)))

	return page
}
