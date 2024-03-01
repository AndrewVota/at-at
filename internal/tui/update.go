package tui

import (
	"github.com/andrewvota/at-at/internal/tui/commands"
	"github.com/andrewvota/at-at/internal/tui/devices"
	"github.com/andrewvota/at-at/internal/tui/inputs"
	"github.com/andrewvota/at-at/internal/tui/messages"
	"github.com/andrewvota/at-at/internal/tui/ports"
	"github.com/andrewvota/at-at/internal/tui/responses"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.WindowWidth = msg.Width
		m.WindowHeight = msg.Height
		m.PortsComponent.WindowWidth = m.WindowWidth
		m.PortsComponent.WindowHeight = m.WindowHeight
		m.DevicesComponent.WindowWidth = m.WindowWidth
		m.DevicesComponent.WindowHeight = m.WindowHeight
		m.CommandsComponent.WindowWidth = m.WindowWidth
		m.CommandsComponent.WindowHeight = m.WindowHeight
		m.InputsComponent.WindowWidth = m.WindowWidth
		m.InputsComponent.WindowHeight = m.WindowHeight
		m.ResponsesComponent.WindowWidth = m.WindowWidth
		m.ResponsesComponent.WindowHeight = m.WindowHeight

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Quit):
			if m.State != messages.TypingInput {
				return m, tea.Quit
			}

		case key.Matches(msg, m.Keys.ToggleActiveRight):
			switch m.State {
			case messages.SelectingPort:
				m.State = messages.SelectingDevice
				m.PortsComponent.Focused = false
				m.DevicesComponent.Focused = true

			case messages.SelectingDevice:
				m.State = messages.SelectingCommand
				m.DevicesComponent.Focused = false
				m.CommandsComponent.Focused = true

			case messages.SelectingCommand:
				m.State = messages.TypingInput
				m.CommandsComponent.Focused = false
				m.InputsComponent.Focused = true
				m.InputsComponent.TextInput.Focus()

			case messages.TypingInput:
				m.State = messages.SelectingPort
				m.InputsComponent.Focused = false
				m.InputsComponent.TextInput.Blur()
				m.PortsComponent.Focused = true
			}

		case key.Matches(msg, m.Keys.ToggleActiveLeft):
			switch m.State {
			case messages.SelectingPort:
				m.State = messages.TypingInput
				m.PortsComponent.Focused = false
				m.InputsComponent.Focused = true
				m.InputsComponent.TextInput.Focus()

			case messages.SelectingDevice:
				m.State = messages.SelectingPort
				m.DevicesComponent.Focused = false
				m.PortsComponent.Focused = true

			case messages.SelectingCommand:
				m.State = messages.SelectingDevice
				m.CommandsComponent.Focused = false
				m.DevicesComponent.Focused = true

			case messages.TypingInput:
				m.State = messages.SelectingCommand
				m.InputsComponent.Focused = false
				m.InputsComponent.TextInput.Blur()
				m.CommandsComponent.Focused = true
			}
		}
	}

	switch m.State {
	case messages.SelectingPort:
		var cmd tea.Cmd
		component, cmd := m.PortsComponent.Update(msg)
		m.PortsComponent = component.(*ports.Model)
		return m, cmd

	case messages.SelectingDevice:
		var cmd tea.Cmd
		component, cmd := m.DevicesComponent.Update(msg)
		m.DevicesComponent = component.(*devices.Model)
		return m, cmd

	case messages.SelectingCommand:
		var cmd tea.Cmd
		component, cmd := m.CommandsComponent.Update(msg)
		m.CommandsComponent = component.(*commands.Model)
		return m, cmd

	case messages.TypingInput:
		var cmds []tea.Cmd
		var cmd tea.Cmd

		component, cmd := m.InputsComponent.Update(msg)
		m.InputsComponent = component.(*inputs.Model)
		cmds = append(cmds, cmd)

		component, cmd = m.ResponsesComponent.Update(msg)
		m.ResponsesComponent = component.(*responses.Model)
		cmds = append(cmds, cmd)

		return m, tea.Batch(cmds...)
	}

	return m, nil
}
