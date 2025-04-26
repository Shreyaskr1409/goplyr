package cmd

import (
	"fmt"
	// "log"

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
	p.width = 40
	p.height = 40

	util.MsgAppendln(&p.messages, fmt.Sprint("SONG:     ", p.song.song))
	util.MsgAppendln(&p.messages, fmt.Sprint("ALBUM:    ", p.song.album))
	util.MsgAppendln(&p.messages, fmt.Sprint("ARTIST:   ", p.song.artist))
	util.MsgAppend(&p.messages, fmt.Sprint("DURATION: ", p.song.duration))

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
	stylePage := lipgloss.NewStyle().Width(p.width).Height(p.height)
	page := ""
	page = fmt.Sprint(page, stylePage.Render(p.PlayerSummary()))

	return page
}

func (p *PlayerWindow) PlayerSummary() string {
	styleBox := lipgloss.NewStyle().Width(int(float32(p.width)/3.5)-2).Height(p.height-2).Align(lipgloss.Left, lipgloss.Top).Padding(1, 2).Border(lipgloss.NormalBorder())

	// innerH := p.height - 4
	// innerW := p.width - 4

	playerSummary := ""
	styleContent := lipgloss.NewStyle().AlignHorizontal(lipgloss.Left)
	styleASCII := lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)
	s := ""

	ascii, err := util.ImageToASCII("./test_art2.png", 20, 9)
	if err != nil {
		fmt.Println("Error:", err)
		ascii = util.GenerateFallbackASCII(uint(p.width/4-2), uint(p.width/4-2))
	}

	for i := range p.messages {
		s = fmt.Sprint(s, p.messages[i])
	}

	playerSummary = fmt.Sprint(playerSummary, styleBox.Render(
		fmt.Sprint(styleASCII.Render(ascii), "\n", styleContent.Render(s))))

	return playerSummary
}
