package commands

import "github.com/charmbracelet/bubbles/list"

type Model struct {
	WindowWidth  int
	WindowHeight int

	List     list.Model
	Focused  bool
	Spinning bool
}

type Device struct {
	Command string
	Desc    string
}

var (
	device1 = Device{Command: "Command 1", Desc: "Desc 1"}
	device2 = Device{Command: "Command 2", Desc: "Desc 2"}
	device3 = Device{Command: "Command 3", Desc: "Desc 3"}
	device4 = Device{Command: "Command 4", Desc: "Desc 4"}
	device5 = Device{Command: "Command 5", Desc: "Desc 5"}
)

func (p *Device) Title() string       { return p.Command }
func (p *Device) Description() string { return p.Desc }
func (p *Device) FilterValue() string { return p.Command }

func New() *Model {
	items := []list.Item{
		&device1,
		&device2,
		&device3,
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Commands"

	return &Model{
		List:    l,
		Focused: false,
        Spinning: false,
	}
}
