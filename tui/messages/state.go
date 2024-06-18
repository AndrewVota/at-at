package messages

import tea "github.com/charmbracelet/bubbletea"

type State int

const (
	StateMenu State = iota
	StateRepl
)

type ChangeStateMessage struct {
	State State
}

func ChangeStateTo(state State) tea.Cmd {
	return func() tea.Msg {
		return ChangeStateMessage{State: state}
	}
}
