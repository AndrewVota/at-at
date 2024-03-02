package commands

import (
	"log"
	"strings"

	"github.com/andrewvota/at-at/internal/tui/messages"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.SendCommandsMessage:
		items := make([]list.Item, len(msg.Commands))
		for i, c := range msg.Commands {
            log.Printf("Command: %s, Details: %s", c.Command, c.Details)
			items[i] = &Command{Command: strings.Replace(c.Command, "\r", "", -1), Details: c.Details}
		}
		l := list.New(items, list.NewDefaultDelegate(), 0, 0)
		l.Title = "Commands"
		m.List = l
        return m, nil
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}
