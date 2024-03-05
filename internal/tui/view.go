package tui

import "github.com/charmbracelet/lipgloss"

func (m *Model) View() string {
	return lipgloss.NewStyle().Margin(2, 2).MaxWidth(m.WindowWidth).MaxHeight(m.WindowHeight).Render(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.PortsComponent.View(),
			m.DevicesComponent.View(),
			m.CommandsComponent.View(),
		),
	)
}
