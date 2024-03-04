package tui

import (
	"github.com/andrewvota/at-at/internal/tui/commands"
	"github.com/andrewvota/at-at/internal/tui/devices"
	"github.com/andrewvota/at-at/internal/tui/messages"
	"github.com/andrewvota/at-at/internal/tui/ports"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.Keys.ToggleActiveRight):
			switch m.State {
			case messages.SelectingPort:
				m.State = messages.SelectingDevice
				cmd = messages.SendStateMessage(m.State)
				cmds = append(cmds, cmd)

			case messages.SelectingDevice:
				m.State = messages.SelectingCommand
				cmd = messages.SendStateMessage(m.State)
				cmds = append(cmds, cmd)

			case messages.SelectingCommand:
				m.State = messages.SelectingDevice
				cmd = messages.SendStateMessage(m.State)
				cmds = append(cmds, cmd)
			}

		case key.Matches(msg, m.Keys.ToggleActiveLeft):
			switch m.State {
			case messages.SelectingPort:
				m.State = messages.SelectingCommand
				cmd = messages.SendStateMessage(m.State)
				cmds = append(cmds, cmd)

			case messages.SelectingDevice:
				m.State = messages.SelectingPort
				cmd = messages.SendStateMessage(m.State)
				cmds = append(cmds, cmd)

			case messages.SelectingCommand:
				m.State = messages.SelectingDevice
				cmd = messages.SendStateMessage(m.State)
				cmds = append(cmds, cmd)
			}
		}

	case tea.WindowSizeMsg:
		m.WindowWidth = msg.Width
		m.WindowHeight = msg.Height
	}

	// Delegate messages to the components
	c, cmd := m.PortsComponent.Update(msg)
	m.PortsComponent = c.(*ports.Model)
	cmds = append(cmds, cmd)

	c, cmd = m.DevicesComponent.Update(msg)
	m.DevicesComponent = c.(*devices.Model)
	cmds = append(cmds, cmd)

	c, cmd = m.CommandsComponent.Update(msg)
	m.CommandsComponent = c.(*commands.Model)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
