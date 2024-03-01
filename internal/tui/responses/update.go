package responses

import (
	"fmt"

	"github.com/andrewvota/at-at/internal/tui/messages"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    currentValue := m.TextArea.Value()
    switch msg := msg.(type) {
    case messages.SendInputMessage:
        m.TextArea.SetValue(fmt.Sprintf("%s\n%s", msg.Input, currentValue))
    }

    var cmd tea.Cmd
    m.TextArea, cmd = m.TextArea.Update(msg)
    return m, cmd
}
