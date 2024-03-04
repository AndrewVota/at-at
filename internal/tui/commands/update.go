package commands

import (
	"github.com/andrewvota/at-at/internal/device"
	"github.com/andrewvota/at-at/internal/tui/messages"
	"github.com/charmbracelet/bubbles/list"
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
		if msg.State == messages.SelectingCommand {
			m.Focused = true
		} else {
			m.Focused = false
		}

	case messages.PortMessage:
		currentPortName := msg.PortName
		if m.CurrentPortName != currentPortName {
			m.CurrentPortName = currentPortName
		}

	case messages.DeviceMessage:
		currentDeviceMake := msg.Make
		currentDeviceModel := msg.Model
		currentDeviceBaudRate := msg.BaudRate
		currentDeviceDateBits := msg.DataBits
		currentDeviceStopBits := msg.StopBits
		currentDeviceParity := msg.Parity
		if m.CurrentDeviceMake != currentDeviceMake && m.CurrentDeviceModel != currentDeviceModel {
			m.CurrentDeviceMake = currentDeviceMake
			m.CurrentDeviceModel = currentDeviceModel
			m.CurrentDeviceBaudRate = currentDeviceBaudRate
			m.CurrentDeviceDataBits = currentDeviceDateBits
			m.CurrentDeviceStopBits = currentDeviceStopBits
			m.CurrentDeviceParity = currentDeviceParity
		}

		commands, err := device.GetConfigForMakeAndModel(m.CurrentDeviceMake, m.CurrentDeviceModel)
		if err != nil {
			panic(err)
		}

		items := make([]list.Item, len(commands.Commands))
		for i, c := range commands.Commands {
			items[i] = &Command{Command: c.Command, Details: c.Details}
		}

		cmd = m.List.SetItems(items)
		cmds = append(cmds, cmd)
	}

	if m.Focused {
		m.List, cmd = m.List.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
