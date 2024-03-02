package commands

import (
	"log"
	"strings"

	"github.com/andrewvota/at-at/internal/serial"
	"github.com/andrewvota/at-at/internal/tui/messages"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.SendCommandsMessage:
		items := make([]list.Item, len(msg.Commands))
		for i, c := range msg.Commands {
			items[i] = &Command{Command: strings.Replace(c.Command, "\r", "", -1), Details: c.Details}
		}
		l := list.New(items, list.NewDefaultDelegate(), 0, 0)
		l.Title = "Commands"
		m.List = l
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Submit):
			currentCommand := m.List.SelectedItem().(*Command).Command
			_, err := serial.GetInstance().SendCommand(currentCommand)
            if err != nil {
                log.Println(err)
            }
            return m, nil
		}
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}
