package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Quit key.Binding
}

var DefaultKeyMap = KeyMap{
	Quit: key.NewBinding(key.WithKeys("ctrl+c")),
}

type Model struct {
	// General settings
	KeyMap KeyMap
	focus  bool
}

func New() Model {
	return Model{
		KeyMap: DefaultKeyMap,
		focus:  false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Quit):
			return m, tea.Quit
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return ""
}

// Ensure that model fulfils the tea.Model interface at compile time.
var _ tea.Model = (*Model)(nil)
