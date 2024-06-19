package tui

import (
	"log"

	"github.com/andrewvota/at-at/tui/menu"
	"github.com/andrewvota/at-at/tui/messages"
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

	// State
	state messages.State

	// Components
	menu menu.Model
}

func New() Model {
	return Model{
		KeyMap: DefaultKeyMap,
		focus:  false,

		menu: menu.New(),
	}
}

func (m Model) Init() tea.Cmd {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	cmd = m.menu.Init()
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("Message: %s", msg)

	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Quit):
			return m, tea.Quit
		}
	}

	m.menu, cmd = m.menu.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.menu.View()
}

// Ensure that model fulfils the tea.Model interface at compile time.
var _ tea.Model = (*Model)(nil)
