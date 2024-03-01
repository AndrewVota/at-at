package messages

import tea "github.com/charmbracelet/bubbletea"

type State int

const (
	SelectingPort State = iota
	SelectingDevice
    SelectingCommand
    TypingInput
)

type ChangeStateMessage struct {
	State State
}

func ChangeStateTo(state State) tea.Cmd {
	return func() tea.Msg {
		return ChangeStateMessage{State: state}
	}
}
