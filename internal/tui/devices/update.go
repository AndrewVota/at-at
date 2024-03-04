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

	case messages.StateMessage:
		if msg.State == messages.SelectingDevice {
			m.Focused = true
		} else {
			m.Focused = false
		}
	}

	if m.Focused {
		m.List, cmd = m.List.Update(msg)
		cmds = append(cmds, cmd)
	}

	newDeviceMake := m.List.SelectedItem().(*Device).Make
	newDeviceModel := m.List.SelectedItem().(*Device).Model
	newDeviceBaudRate := m.List.SelectedItem().(*Device).BaudeRate
	newDeviceDataBits := m.List.SelectedItem().(*Device).DataBits
	newDeviceStopBits := m.List.SelectedItem().(*Device).StopBits
	newDeviceParity := m.List.SelectedItem().(*Device).Parity
	if m.CurrentDeviceMake != newDeviceMake && m.CurrentDeviceModel != newDeviceModel {
		m.CurrentDeviceMake = newDeviceMake
		m.CurrentDeviceModel = newDeviceModel
		cmd = messages.SendDeviceMessage(newDeviceMake, newDeviceModel, newDeviceBaudRate, newDeviceDataBits, newDeviceStopBits, newDeviceParity)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
