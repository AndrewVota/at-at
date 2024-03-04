package ports

import (
	"github.com/andrewvota/at-at/internal/tui/messages"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.WindowWidth = msg.Width
		m.WindowHeight = msg.Height

	case messages.StateMessage:
		if msg.State == messages.SelectingPort {
			m.Focused = true
		} else {
			m.Focused = false
		}
	}

	if m.Focused {
		m.List, cmd = m.List.Update(msg)
		cmds = append(cmds, cmd)
	}

	currentPort := m.List.SelectedItem().(*Port).Name
	if m.CurrentPort != currentPort {
		m.CurrentPort = currentPort
		cmd = messages.SendPortMessage(currentPort)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
