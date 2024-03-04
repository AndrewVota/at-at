package tui

import "github.com/charmbracelet/lipgloss"

func (m *Model) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.PortsComponent.View(),
		m.DevicesComponent.View(),
	)
}
