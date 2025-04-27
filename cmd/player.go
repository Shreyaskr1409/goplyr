package cmd

import (
	"fmt"
	"strings"

	// "log"

	"github.com/Shreyaskr1409/goplyr/cmd/core/player"
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
	playbackStatus := PLAY

	pw.song = &Song{
		song:        "Instant Crush (feat. Julian Casablancas)",
		album:       "Random Access Memories",
		artist:      "Daft Punk",
		duration:    "5:38",
		year:        "2013",
		albumArtURI: "",
	}
	pw.playerState = &PlayerState{
		playbackMode:   &playbackMode,
		playbackStatus: &playbackStatus,
		timestamp:      "00:00",
		volume:         100,
	}

	return pw
}

func (p *PlayerWindow) Init() tea.Cmd {
	p.width = 40
	p.height = 40
	keyStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("0"))

	util.MsgAppendln(&p.messages, fmt.Sprint(keyStyle.Render("SONG:    "), " ", p.song.song))
	util.MsgAppendln(&p.messages, fmt.Sprint(keyStyle.Render("ALBUM:   "), " ", p.song.album))
	util.MsgAppendln(&p.messages, fmt.Sprint(keyStyle.Render("ARTIST:  "), " ", p.song.artist))
	util.MsgAppendln(&p.messages, fmt.Sprint(keyStyle.Render("DURATION:"), " ", p.song.duration))

	return nil
}

func (p *PlayerWindow) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return p, tea.Quit
		case "ctrl+p":
			if *p.playerState.playbackStatus == PLAY {
				// CODE TO PLAY THE SONG
			} else {
				// CODE TO PLAY THE SONG
			}
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
	boxWidth := int(float32(p.width) / 3.5)
	if boxWidth < 30 {
		boxWidth = 30
	}
	boxHeight := p.height
	paddingWidth := 2
	paddingHeight := 1

	styleBox := lipgloss.NewStyle().
		Width(boxWidth-2).
		Height(boxHeight-2).
		Align(lipgloss.Left, lipgloss.Top).
		Padding(paddingHeight, paddingWidth).
		Border(lipgloss.NormalBorder())

	innerWidth := boxWidth - 2 - paddingWidth*2    // 2 for borders and 4 for padding
	artWidth := innerWidth - 2                     // 2 for borders and 4 for padding
	artHeight := int(float32(innerWidth)*9/20) - 2 // 20 : 9 is the ratio which makes a square in my fonts

	playerSummary := ""
	styleContent := lipgloss.NewStyle().AlignHorizontal(lipgloss.Left)
	styleASCII := lipgloss.NewStyle().
		AlignHorizontal(lipgloss.Center).
		Width(artWidth).
		Border(lipgloss.DoubleBorder())
	s := ""

	ascii, err := util.ImageToASCII("./test_art3.png", uint(artWidth), uint(artHeight))
	if err != nil {
		fmt.Println("Error:", err)
		ascii = util.GenerateFallbackASCII(uint(p.width/4-2), uint(p.width/4-2))
	}

	horizontalRule := strings.Repeat("\u2500", innerWidth)
	s = fmt.Sprint(s, horizontalRule, "\n")
	for i := range p.messages {
		s = fmt.Sprint(s, p.messages[i])
	}
	s = fmt.Sprint(s, horizontalRule)

	controls := []string{
		lipgloss.NewStyle().Foreground(lipgloss.Color("14")).Render("Play/Pause → Ctrl+P     "),
		lipgloss.NewStyle().Foreground(lipgloss.Color("11")).Render("Exit       → Ctrl+C or Q"),
	}

	styleControls := lipgloss.NewStyle().AlignHorizontal(lipgloss.Center).
		Width(innerWidth - 2).
		Border(lipgloss.DoubleBorder()).
		Background(lipgloss.Color("0")).
		BorderBackground(lipgloss.Color("0"))
	formattedControls := make([]string, len(controls))
	for i, control := range controls {
		formattedControls[i] = styleControls.Render(control)
	}

	playerSummary = fmt.Sprint(playerSummary,
		styleBox.Render(fmt.Sprint(styleASCII.Render(ascii),
			"\n",
			styleContent.Render(s),
			"\n",
			styleControls.Render(strings.Join(controls, "\n")),
		)))

	return playerSummary
}
