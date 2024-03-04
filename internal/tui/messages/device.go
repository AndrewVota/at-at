package messages

import tea "github.com/charmbracelet/bubbletea"

type DeviceMessage struct {
	MakeName  string
	ModelName string
}

func SendDeviceMessage(make string, model string) tea.Cmd {
	return func() tea.Msg {
		return DeviceMessage{MakeName: make, ModelName: model}
	}
}
