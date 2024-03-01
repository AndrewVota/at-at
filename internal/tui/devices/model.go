package devices

import "github.com/charmbracelet/bubbles/list"

type Model struct {
	WindowWidth  int
	WindowHeight int

	List    list.Model
	Focused bool
}

type Device struct {
	Make  string
	Model string
	Type  string
}

var (
	device1 = Device{Make: "Make 1", Model: "Model 1", Type: "Type 1"}
	device2 = Device{Make: "Make 2", Model: "Model 2", Type: "Type 2"}
	device3 = Device{Make: "Make 3", Model: "Model 3", Type: "Type 3"}
	device4 = Device{Make: "Make 4", Model: "Model 4", Type: "Type 4"}
	device5 = Device{Make: "Make 5", Model: "Model 5", Type: "Type 5"}
)

func (p *Device) Title() string       { return p.Make + " " + p.Model }
func (p *Device) Description() string { return p.Type }
func (p *Device) FilterValue() string { return p.Make + " " + p.Model }

func New() *Model {
	items := []list.Item{
		&device1,
		&device2,
		&device3,
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Devices"

	return &Model{
		List:    l,
		Focused: false,
	}
}
