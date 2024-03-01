package inputs

import (
	"log"
	"time"

	"github.com/andrewvota/at-at/internal/tui/messages"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Submit):
            if m.Spinning {
                return m, nil
            }
			cmd := messages.SendInput(messages.Input(m.TextInput.Value()))
			m.Spinning = true
			time.AfterFunc(3*time.Second, func() {
                m.Spinning = false
                m.TextInput.Reset()
                log.Println("Done")
			})
			return m, cmd
		}
	}

	var cmds []tea.Cmd
	var cmd tea.Cmd
	if m.Spinning {
		m.Spinner, cmd = m.Spinner.Update(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	} else {
		m.TextInput, cmd = m.TextInput.Update(msg)
		cmds = append(cmds, cmd)
		cmds = append(cmds, m.Spinner.Tick)
		return m, tea.Batch(cmds...)
	}
}
