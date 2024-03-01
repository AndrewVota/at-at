package responses

import (
	"github.com/charmbracelet/bubbles/textarea"
)

type Model struct {
	WindowWidth  int
	WindowHeight int

	TextArea textarea.Model
}

func New() *Model {
	t := textarea.New()
    
	return &Model{
		TextArea: t,
	}
}
