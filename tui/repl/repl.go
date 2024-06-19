package repl

import (
	"github.com/andrewvota/at-at/tui/messages"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Submit key.Binding
	Return key.Binding
}

var DefaultKeyMap = KeyMap{
	Submit: key.NewBinding(key.WithKeys("enter")),
	Return: key.NewBinding(key.WithKeys("esc")),
}

type Model struct {
	// General settings
	KeyMap KeyMap
	focus  bool
	width  int
	height int

	// State

	// Components
	// textInput textinput.Model
	textArea textarea.Model
}

func New() Model {
	var textArea = textarea.New()
	textArea.Focus()

	return Model{
		KeyMap: DefaultKeyMap,
		focus:  false,
		width:  0,
		height: 0,

		// textInput: textinput.New(),
		textArea: textArea,
	}
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.focus {
		return m, nil
	}

	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Return):
			return m, messages.ChangeStateTo(messages.StateMenu)
		}
	}

	m.textArea, cmd = m.textArea.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.textArea.View()
}

// ---

func (m *Model) Focus() {
	m.focus = true
}

func (m *Model) Blur() {
	m.focus = false
}
