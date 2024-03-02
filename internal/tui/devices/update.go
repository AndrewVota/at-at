package devices

import (
	"github.com/andrewvota/at-at/internal/tui/messages"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
    cmds = append(cmds, cmd)

	selected := m.List.SelectedItem()
    messagesCommands := make([]messages.Command, len(selected.(*Device).Commands))
    for i, c := range selected.(*Device).Commands {
        messagesCommands[i] = messages.Command{Command: c.Command, Details: c.Details}
    }

    cmd = messages.SendCommands(messagesCommands)
    cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
