package messages

import tea "github.com/charmbracelet/bubbletea"

type DeviceMessage struct {
	Make     string
	Model    string
	BaudRate int
	DataBits int
	StopBits float32
	Parity   string
}

func SendDeviceMessage(make string, model string, baudRate int, dataBits int, stopBits float32, parity string) tea.Cmd {
	return func() tea.Msg {
		return DeviceMessage{Make: make, Model: model, BaudRate: baudRate, DataBits: dataBits, StopBits: stopBits, Parity: parity}
	}
}
