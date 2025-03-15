package cli

import tea "github.com/charmbracelet/bubbletea"

type CliTool struct {
	windows    map[string]Window
	windowName string
	statusBar  StatusBar
}

func (m *CliTool) Init() tea.Cmd {
	return m.windows[m.windowName].Init()
}

func (m *CliTool) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.windows[m.windowName].Update(msg)
}

func (m *CliTool) View() string {
	return m.windows[m.windowName].View()
}

func InitGoplyr() *CliTool {
	c := &CliTool{}

	return c
}
