package devices

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
	}

	if m.Focused {
		m.List, cmd = m.List.Update(msg)
		cmds = append(cmds, cmd)
	}

	currentDeviceMake := m.List.SelectedItem().(*Device).Make
	currentDeviceModel := m.List.SelectedItem().(*Device).Model
	if m.CurrentDeviceMake != currentDeviceMake && m.CurrentDeviceModel != currentDeviceModel {
		m.CurrentDeviceMake = currentDeviceMake
		m.CurrentDeviceModel = currentDeviceModel
		cmd = messages.SendDeviceMessage(currentDeviceMake, currentDeviceModel)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
