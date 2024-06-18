package selector

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
	NextElement key.Binding
	PrevElement key.Binding
}

var DefaultKeyMap = KeyMap{
	NextElement: key.NewBinding(key.WithKeys("right", "l", "d")),
	PrevElement: key.NewBinding(key.WithKeys("left", "h", "a")),
}

type Model struct {
	// General settings
	KeyMap KeyMap
	focus  bool

	// Text settings
	placeholder   string
	choices       []string
	currentChoice int

	// Styling
	TextStyleFocused lipgloss.Style
	TextStyleBlurred lipgloss.Style
}

func New() Model {
	return Model{
		KeyMap: DefaultKeyMap,
		focus:  false,

		placeholder:   "",
		choices:       []string{""},
		currentChoice: 0,

		TextStyleFocused: lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		TextStyleBlurred: lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.focus {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.PrevElement):
			if m.currentChoice == 0 {
				m.currentChoice = len(m.choices) - 1
			} else {
				m.currentChoice--
			}
		case key.Matches(msg, m.KeyMap.NextElement):
			if m.currentChoice == len(m.choices)-1 {
				m.currentChoice = 0
			} else {
				m.currentChoice++
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	if len(m.choices) == 1 && m.choices[0] == m.placeholder {
		return m.placeholderView()
	}
	return m.choiceView()
}

// ---

func (m *Model) Focus() {
	m.focus = true
}

func (m *Model) Blur() {
	m.focus = false
}

func (m *Model) Reset() {
	m.currentChoice = 0
	m.choices = []string{m.placeholder}
}

func (m *Model) AddChoices(choices ...string) {
	m.choices = append(m.choices, choices...)

	if len(m.choices) > 1 && m.choices[0] == m.placeholder {
		m.choices = m.choices[1:]
	}
}

func (m *Model) SetPlaceholder(s string) {
	var oldPlaceholder = m.placeholder
	m.placeholder = s

	if len(m.choices) == 1 && m.choices[m.currentChoice] == oldPlaceholder {
		m.choices[0] = m.placeholder
	}
}

// ---

func (m *Model) placeholderView() string {
	var t = fmt.Sprintf("< %s >", m.placeholder)

	if m.focus {
		return m.TextStyleFocused.Render(t)
	}
	return m.TextStyleBlurred.Render(t)
}

func (m *Model) choiceView() string {
	var t = fmt.Sprintf("< %s >", m.choices[m.currentChoice])

	if m.focus {
		return m.TextStyleFocused.Render(t)
	}
	return m.TextStyleBlurred.Render(t)
}
