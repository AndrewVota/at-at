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

	currentDeviceMake := m.List.SelectedItem().(*Device).Make
	currentDeviceModel := m.List.SelectedItem().(*Device).Model
	currentDeviceBaudeRate := m.List.SelectedItem().(*Device).BaudeRate
	currentDeviceDataBits := m.List.SelectedItem().(*Device).DataBits
	currentDeviceStopBits := m.List.SelectedItem().(*Device).StopBits
	currentDeviceParity := m.List.SelectedItem().(*Device).Parity
	if m.CurrentDeviceMake != currentDeviceMake && m.CurrentDeviceModel != currentDeviceModel {
		m.CurrentDeviceMake = currentDeviceMake
		m.CurrentDeviceModel = currentDeviceModel
		cmd = messages.SendDeviceMessage(currentDeviceMake, currentDeviceModel, currentDeviceBaudeRate, currentDeviceDataBits, currentDeviceStopBits, currentDeviceParity)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
