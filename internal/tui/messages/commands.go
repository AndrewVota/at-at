package messages

import tea "github.com/charmbracelet/bubbletea"

type Command struct {
	Command string
	Details string
}

type SendCommandsMessage struct {
	Commands []Command
}

func SendCommands(commands []Command) tea.Cmd {
	return func() tea.Msg {
		return SendCommandsMessage{Commands: commands}
	}
}
