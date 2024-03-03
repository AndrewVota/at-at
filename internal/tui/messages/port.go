package messages

import tea "github.com/charmbracelet/bubbletea"

type PortMessage struct {
	PortName string
}

func SendPortMessage(portName string) tea.Cmd {
	return func() tea.Msg {
		return PortMessage{PortName: portName}
	}
}
