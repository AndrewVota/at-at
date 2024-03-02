package commands

import "github.com/charmbracelet/bubbles/list"

type Model struct {
	WindowWidth  int
	WindowHeight int

	List     list.Model
	Focused  bool
	Spinning bool
}

type Command struct {
	Command string
	Details    string
}

func (p *Command) Title() string       { return p.Command }
func (p *Command) Description() string { return p.Details }
func (p *Command) FilterValue() string { return p.Command }

func New() *Model {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Commands"

	return &Model{
		List:    l,
		Focused: false,
        Spinning: false,
	}
}
