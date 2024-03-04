package devices

import (
	"github.com/charmbracelet/bubbles/list"
)

type Model struct {
	WindowWidth  int
	WindowHeight int

	List    list.Model
	Focused bool

	CurrentDeviceMake  string
	CurrentDeviceModel string
}

type Device struct {
	Make  string
	Model string
	Type  string
}

func (p *Device) Title() string       { return p.Make + " " + p.Model }
func (p *Device) Description() string { return p.Type }
func (p *Device) FilterValue() string { return p.Make + " " + p.Model }

func New() *Model {
	items := []list.Item{
		&Device{Make: "Apple", Model: "iPhone 12", Type: "Smartphone"},
		&Device{Make: "Apple", Model: "iPhone 12 Pro", Type: "Smartphone"},
		&Device{Make: "Apple", Model: "iPhone 12 Pro Max", Type: "Smartphone"},
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)

	return &Model{
		WindowWidth:  0,
		WindowHeight: 0,

		List:    l,
		Focused: false,

		CurrentDeviceMake:  "",
		CurrentDeviceModel: "",
	}
}
