package repl

import (
	"github.com/andrewvota/at-at/tui/messages"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Submit key.Binding
}

var DefaultKeyMap = KeyMap{
	Submit: key.NewBinding(key.WithKeys("enter")),
}

type Model struct {
	// General settings
	KeyMap KeyMap
	focus  bool
	width  int
	height int

	// State

	// Components
}

func New() Model {
	return Model{
		KeyMap: DefaultKeyMap,
		focus:  true,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.focus {
		return m, nil
	}

	var (
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Submit):
			return m, messages.ChangeStateTo(messages.StateMenu)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return "repl..."
}

// ---

func (m *Model) Focus() {
	m.focus = true
}

func (m *Model) Blur() {
	m.focus = false
}
