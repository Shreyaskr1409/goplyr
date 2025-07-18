package main

import (
	"log"
	"os"

	"github.com/Shreyaskr1409/goplyr/cmd"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println("Application starts")

	pw := cmd.InitPlayerWindow()
	p := tea.NewProgram(pw, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalln(err)
	}
}
