package ports

import (
	"github.com/andrewvota/at-at/internal/tui/messages"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
    cmds = append(cmds, cmd)

    selectedPort := m.List.SelectedItem()
    portName := selectedPort.(*Port).Name
    cmd = messages.SendPortName(portName)
    cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
