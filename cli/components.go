package cli

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Window interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
}

type StatusBar struct {
	song      string
	album     string
	artist    string
	volume    string
	length    string
	timestamp string
}
