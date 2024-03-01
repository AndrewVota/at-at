package inputs

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	Keys         KeyMap
	WindowWidth  int
	WindowHeight int

	TextInput textinput.Model
	Focused   bool
}

func New() *Model {
	ti := textinput.New()
	ti.Placeholder = "Enter custom commands here..."
	ti.CharLimit = 20

	return &Model{
		Keys:      DefaultKeyMap(),
		TextInput: ti,
		Focused:   false,
	}
}
