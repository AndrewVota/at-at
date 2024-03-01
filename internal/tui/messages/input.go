package messages

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Input string

type SendInputMessage struct {
	Input Input
}

func SendInput(input Input) tea.Cmd {
	return func() tea.Msg {
		return SendInputMessage{Input: input}
	}
}
