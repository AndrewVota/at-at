package selector

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const placeholder = "NO SERIAL DEVICE DETECTED"

type Model struct {
	keys          keyMap
	choices       []string
	currentChoice int
	focused       bool
}

func NewSelector() Model {
	return Model{
		keys:          defaultKeyMap(),
		choices:       []string{placeholder},
		currentChoice: 0,
		focused:       false,
	}
}

type keyMap struct {
	Left  key.Binding
	Right key.Binding
}

func defaultKeyMap() keyMap {
	return keyMap{
		Left: key.NewBinding(
			key.WithKeys("left", "h"),
		),
		Right: key.NewBinding(
			key.WithKeys("right", "l"),
		),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.focused {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Left):
			if m.currentChoice == 0 {
				m.currentChoice = len(m.choices) - 1
			} else {
				m.currentChoice--
			}
		case key.Matches(msg, m.keys.Right):
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
	var unfocusedStyle = lipgloss.NewStyle().
		Bold(false)

	var focusedStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4"))

	var style lipgloss.Style

	if m.focused {
		style = focusedStyle
	} else {
		style = unfocusedStyle
	}

	if len(m.choices) == 1 {
		return fmt.Sprintf(style.Render("%s"), m.choices[m.currentChoice])

	}
	return fmt.Sprintf(style.Render("< %s >"), m.choices[m.currentChoice])
}

func (m *Model) AddChoices(choices ...string) {
	for _, choice := range choices {
		m.choices = append(m.choices, choice)
	}

	if len(m.choices) > 1 && m.choices[0] == placeholder {
		m.choices = m.choices[1:]
	}
}
