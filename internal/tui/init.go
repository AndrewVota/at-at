package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Init() tea.Cmd {
    // Wonky work-around for poor early decisions
    // Trigger right-left key sequence to start program
    // and populate the initial state of devices and commands
    cmds := []tea.Cmd{
        tea.Tick(time.Millisecond*10, func(t time.Time) tea.Msg {
            return tea.KeyMsg{Type: tea.KeyRight}
        }),
        tea.Tick(time.Millisecond*20, func(t time.Time) tea.Msg {
            return tea.KeyMsg{Type: tea.KeyLeft}
        }),
    }

    return tea.Batch(cmds...)
}
