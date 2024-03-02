package messages

import tea "github.com/charmbracelet/bubbletea"


type SendPortNameMessage struct {
	PortName string
}

func SendPortName(name string) tea.Cmd {
	return func() tea.Msg {
		return SendPortNameMessage{PortName: name}
	}
}
