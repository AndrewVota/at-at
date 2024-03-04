package messages

import tea "github.com/charmbracelet/bubbletea"

type State int

const (
	SelectingPort State = iota
	SelectingDevice
	SelectingCommand
	TypingInput
)

type StateMessage struct {
	State State
}

func SendStateMessage(state State) tea.Cmd {
	return func() tea.Msg {
		return StateMessage{State: state}
	}
}
