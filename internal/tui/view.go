package tui

import "github.com/charmbracelet/lipgloss"

func (m *Model) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.PortsComponent.View(),
		m.DevicesComponent.View(),
		m.CommandsComponent.View(),

		lipgloss.JoinVertical(
			lipgloss.Top,
			m.InputsComponent.View(),
			m.ResponsesComponent.View(),
		),
	)
}
