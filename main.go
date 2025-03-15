package main

import (
	"log"

	"github.com/Shreyaskr1409/goplyr/cmd"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	pw := cmd.InitPlayerWindow()
	p := tea.NewProgram(pw, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalln(err)
	}
}
