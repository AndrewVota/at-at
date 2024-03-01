package inputs

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	Keys         KeyMap
	WindowWidth  int
	WindowHeight int

	TextInput textinput.Model
	Focused   bool
	Spinner   spinner.Model
    Spinning  bool
}

func New() *Model {
	ti := textinput.New()
	ti.Placeholder = "Enter custom commands here..."
	ti.CharLimit = 20

    s := spinner.New()
    s.Spinner = spinner.Dot


	return &Model{
		Keys:      DefaultKeyMap(),
		TextInput: ti,
		Focused:   false,
        Spinner:   s,
        Spinning:  false,
	}
}
